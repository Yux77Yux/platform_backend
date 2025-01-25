package dispatch

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	cache "github.com/Yux77Yux/platform_backend/microservices/comment/cache"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
)

/*
	这里的链表不太符合高并发特点的设计，问题在于持有锁时间会很长
	改进的办法是使用堆建立，或者使用HASH对节点进行映射
	先留着，以后再建堆
*/

func InitialDeleteChain() *DeleteChain {
	_chain := &DeleteChain{
		Head:       &DeleteListener{prev: nil},
		Tail:       &DeleteListener{next: nil},
		Count:      0,
		exeChannel: make(chan *[]*generated.AfterAuth, 3),
		listenerPool: sync.Pool{
			New: func() any {
				return &DeleteListener{
					commentChannel:  make(chan *generated.AfterAuth, LISTENER_CHANNEL_COUNT),
					timeoutDuration: 10 * time.Second,
					updateInterval:  3 * time.Second,
				}
			},
		},
	}
	_chain.Head.next = _chain.Tail
	_chain.Tail.prev = _chain.Head
	go _chain.ExecuteBatch()
	return _chain
}

// 责任链
type DeleteChain struct {
	Head         *DeleteListener // 责任链的头部
	Tail         *DeleteListener
	nodeMux      sync.Mutex
	Count        int32 // 监听者数量
	exeChannel   chan *[]*generated.AfterAuth
	listenerPool sync.Pool
}

func (chain *DeleteChain) ExecuteBatch() {
	for delCommentsPtr := range chain.exeChannel {
		go func(delCommentsPtr *[]*generated.AfterAuth) {
			delComments := *delCommentsPtr
			// 更新数据库
			affectedCount, err := db.BatchUpdateDeleteStatus(delComments)
			if err != nil {
				log.Printf("error: BatchUpdateDeleteStatus error")
			}
			// 更新Redis
			id := delComments[0].GetCreationId()
			err = cache.UpdateCommentsCount(id, affectedCount)
			if err != nil {
				log.Printf("error: UpdateCommentsCount %v", err)
			}

			*delCommentsPtr = delComments[:0] // 清空切片内容
			delCommentsPool.Put(delCommentsPtr)
		}(delCommentsPtr)
	}
}

// 处理评论请求的函数
func (chain *DeleteChain) HandleRequest(data protoreflect.ProtoMessage) {
	listener := chain.FindListener(data)
	if listener == nil {
		// 如果没有找到合适的监听者，创建一个新的监听者
		listener = chain.CreateListener(data)
	}
	listener.Dispatch(data)
}

// 查找责任链中的合适监听者
func (chain *DeleteChain) FindListener(data protoreflect.ProtoMessage) ListenerInterface {
	comment, ok := data.(*generated.AfterAuth)
	if !ok {
		log.Printf("invalid type: expected *generated.AfterAuth")
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
func (chain *DeleteChain) CreateListener(data protoreflect.ProtoMessage) ListenerInterface {
	comment, ok := data.(*generated.AfterAuth)
	if !ok {
		log.Printf("invalid type: expected *generated.AfterAuth")
	}

	newListener := chain.listenerPool.Get().(*DeleteListener)
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
func (chain *DeleteChain) DestroyListener(listener ListenerInterface) {
	// 找到前一个节点（假设 chain.Head 是链表的头部）
	current, ok := listener.(*DeleteListener)
	if !ok {
		log.Printf("invalid type: expected *DeleteListener")
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
