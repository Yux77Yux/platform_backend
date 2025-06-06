package dispatch

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	tools "github.com/Yux77Yux/platform_backend/microservices/user/tools"
)

func InitialUserBioCacheChain() *UserBioCacheChain {
	_chain := &UserBioCacheChain{
		Head:       &UserBioCacheListener{prev: nil},
		Tail:       &UserBioCacheListener{next: nil},
		Count:      0,
		exeChannel: make(chan *[]*generated.UserUpdateBio, EXE_CHANNEL_COUNT),
		pool: sync.Pool{
			New: func() any {
				slice := make([]*generated.UserUpdateBio, 0, MAX_BATCH_SIZE)
				return &slice
			},
		},
	}
	_chain.listenerPool = sync.Pool{
		New: func() any {
			return &UserBioCacheListener{
				chain:           _chain,
				timeoutDuration: 10 * time.Second,
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
type UserBioCacheChain struct {
	Head         *UserBioCacheListener // 责任链的头部
	Tail         *UserBioCacheListener
	nodeMux      sync.Mutex
	Count        int32 // 监听者数量
	exeChannel   chan *[]*generated.UserUpdateBio
	listenerPool sync.Pool
	pool         sync.Pool
}

func (chain *UserBioCacheChain) Close(signal chan any) {
	cond := sync.NewCond(&chain.nodeMux)

	chain.nodeMux.Lock()
	for atomic.LoadInt32(&chain.Count) > 0 {
		cond.Wait() // 等待 Count 变成 0
	}
	chain.nodeMux.Unlock()

	close(signal)
}

func (chain *UserBioCacheChain) GetPoolObj() any {
	return chain.pool.Get()
}

func (chain *UserBioCacheChain) ExecuteBatch() {
	for usersPtr := range chain.exeChannel {
		go func(usersPtr *[]*generated.UserUpdateBio) {
			users := *usersPtr
			// 更新头像
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			err := cache.UpdateUserBio(ctx, users)
			cancel()
			if err != nil {
				tools.LogError("", "cache UpdateUserBio", err)
				return
			}

			*usersPtr = users[:0]
			chain.pool.Put(usersPtr)
		}(usersPtr)
	}
}

// 处理评论请求的函数
func (chain *UserBioCacheChain) HandleRequest(data protoreflect.ProtoMessage) {
	listener := chain.FindListener(data)
	if listener == nil {
		// 如果没有找到合适的监听者，创建一个新的监听者
		listener = chain.CreateListener(data)
	}
	listener.Dispatch(data)
}

// 查找责任链中的合适监听者
func (chain *UserBioCacheChain) FindListener(data protoreflect.ProtoMessage) ListenerInterface {
	chain.nodeMux.Lock()
	next := chain.Head.next
	prev := chain.Tail.prev
	for {
		if prev == chain.Head {
			break
		}
		if atomic.LoadUint32(&next.count) < LISTENER_CHANNEL_COUNT {
			chain.nodeMux.Unlock()
			return next
		}
		if atomic.LoadUint32(&prev.count) < LISTENER_CHANNEL_COUNT {
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
func (chain *UserBioCacheChain) CreateListener(data protoreflect.ProtoMessage) ListenerInterface {
	newListener := chain.listenerPool.Get().(*UserBioCacheListener)
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
func (chain *UserBioCacheChain) DestroyListener(listener ListenerInterface) {
	// 找到前一个节点（假设 chain.Head 是链表的头部）
	current, ok := listener.(*UserBioCacheListener)
	if !ok {
		log.Printf("invalid type: expected *UserBioCacheListener")
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
