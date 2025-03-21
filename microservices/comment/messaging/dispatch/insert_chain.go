package dispatch

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	tools "github.com/Yux77Yux/platform_backend/microservices/comment/tools"
)

func InitialInsertChain() *InsertChain {
	_chain := &InsertChain{
		Head:       &InsertListener{prev: nil},
		Tail:       &InsertListener{next: nil},
		Count:      0,
		exeChannel: make(chan *[]*generated.Comment, 8),
		pool: sync.Pool{
			New: func() any {
				slice := make([]*generated.Comment, 0, MAX_BATCH_SIZE)
				return &slice
			},
		},
	}
	_chain.listenerPool = sync.Pool{
		New: func() any {
			return &InsertListener{
				chain:           _chain,
				timeoutDuration: 12 * time.Second,
				updateInterval:  3 * time.Second,
			}
		},
	}
	_chain.Head.next = _chain.Tail
	_chain.Tail.prev = _chain.Head
	go _chain.ExecuteBatch()
	return _chain
}

// 责任链
type InsertChain struct {
	Head         *InsertListener // 责任链的头部
	Tail         *InsertListener
	Count        int32 // 监听者数量
	nodeMux      sync.Mutex
	exeChannel   chan *[]*generated.Comment
	listenerPool sync.Pool
	pool         sync.Pool
}

func (chain *InsertChain) Close(signal chan any) {
	cond := sync.NewCond(&chain.nodeMux)

	chain.nodeMux.Lock()
	for atomic.LoadInt32(&chain.Count) > 0 {
		cond.Wait() // 等待 Count 变成 0
	}
	chain.nodeMux.Unlock()

	// 关闭信号通道
	close(signal)
}

func (chain *InsertChain) GetPoolObj() any {
	return chain.pool.Get()
}

func (chain *InsertChain) ExecuteBatch() {
	for insertCommentsPtr := range chain.exeChannel {
		go func(insertCommentsPtr *[]*generated.Comment) {
			insertComments := *insertCommentsPtr
			// 插入数据库
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			affectedCount, err := db.BatchInsert(ctx, insertComments)
			cancel()
			if err != nil {
				tools.LogError("", "db BatchInsert", err)
				// 入死信
				return
			}

			// 更新Redis
			id := insertComments[0].GetCreationId()
			ctx, cancel = context.WithTimeout(context.Background(), time.Second*30)
			err = cache.UpdateCommentsCount(ctx, id, affectedCount)
			cancel()
			if err != nil {
				tools.LogError("", "cache UpdateCommentsCount", err)
				// 入死信
				return
			}

			*insertCommentsPtr = insertComments[:0]
			chain.pool.Put(insertCommentsPtr)
		}(insertCommentsPtr)
	}
}

// 处理评论请求的函数
func (chain *InsertChain) HandleRequest(data protoreflect.ProtoMessage) {
	listener := chain.FindListener(data)
	if listener == nil {
		// 如果没有找到合适的监听者，创建一个新的监听者
		listener = chain.CreateListener(data)
	}
	listener.Dispatch(data)
}

// 查找责任链中的合适监听者
func (chain *InsertChain) FindListener(data protoreflect.ProtoMessage) ListenerInterface {
	comment, ok := data.(*generated.Comment)
	if !ok {
		log.Printf(": expected *generated.Comment")
	}

	creationId := comment.GetCreationId()

	chain.nodeMux.Lock()
	next := chain.Head.next
	prev := chain.Tail.prev
	for {
		if prev == chain.Head {
			break
		}
		if next.creationId == creationId {
			chain.nodeMux.Unlock()
			return next
		}
		if prev.creationId == creationId {
			chain.nodeMux.Unlock()
			return prev
		}
		if prev == next || prev.prev == next {
			// 找不到
			break
		}
		prev = prev.prev
		next = next.next
	}
	chain.nodeMux.Unlock()
	return nil // 没有找到合适的监听者
}

// 创建一个新的监听者
func (chain *InsertChain) CreateListener(data protoreflect.ProtoMessage) ListenerInterface {
	comment, ok := data.(*generated.Comment)
	if !ok {
		log.Printf("invalid type: expected *generated.Comment")
	}

	newListener := chain.listenerPool.Get().(*InsertListener)
	newListener.creationId = comment.GetCreationId()
	newListener.exeChannel = chain.exeChannel

	// 头插法，将新的监听者挂到链中
	chain.nodeMux.Lock()
	next := chain.Head.next

	newListener.next = next
	newListener.prev = chain.Head

	chain.Head.next = newListener
	next.prev = newListener
	chain.nodeMux.Unlock()

	atomic.AddInt32(&chain.Count, 1)

	newListener.StartListening() // 启动监听
	return newListener
}

// 销毁监听者
func (chain *InsertChain) DestroyListener(listener ListenerInterface) {
	// 找到前一个节点（假设 chain.Head 是链表的头部）
	current, ok := listener.(*InsertListener)
	if !ok {
		log.Printf("invalid type: expected *InsertListener")
	}

	chain.nodeMux.Lock()
	prev := current.prev
	next := current.next
	prev.next = next
	next.prev = prev
	chain.nodeMux.Unlock()
	atomic.AddInt32(&chain.Count, -1)

	chain.listenerPool.Put(listener)
}
