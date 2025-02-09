package dispatch

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	cache "github.com/Yux77Yux/platform_backend/microservices/user/cache"
)

func InitialFollowCacheChain() *FollowCacheChain {
	_chain := &FollowCacheChain{
		Head:       &FollowCacheListener{prev: nil},
		Tail:       &FollowCacheListener{next: nil},
		Count:      0,
		exeChannel: make(chan *[]*generated.Follow, EXE_CHANNEL_COUNT),
		listenerPool: sync.Pool{
			New: func() any {
				return &FollowCacheListener{
					usersChannel:    make(chan *generated.Follow, LISTENER_CHANNEL_COUNT),
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
type FollowCacheChain struct {
	Head         *FollowCacheListener // 责任链的头部
	Tail         *FollowCacheListener
	Count        int32 // 监听者数量
	nodeMux      sync.Mutex
	exeChannel   chan *[]*generated.Follow
	listenerPool sync.Pool
}

func (chain *FollowCacheChain) ExecuteBatch() {
	for FollowUsersPtr := range chain.exeChannel {
		go func(FollowUsersPtr *[]*generated.Follow) {
			FollowUsers := *FollowUsersPtr

			// 插入Redis
			err := cache.Follow(FollowUsers)
			if err != nil {
				log.Printf("error: Follow error %v", err)
			}

			// 放回对象池
			*FollowUsersPtr = FollowUsers[:0]
			followPool.Put(FollowUsersPtr)
		}(FollowUsersPtr)
	}
}

// 处理评论请求的函数
func (chain *FollowCacheChain) HandleRequest(data protoreflect.ProtoMessage) {
	listener := chain.FindListener(data)
	if listener == nil {
		// 如果没有找到合适的监听者，创建一个新的监听者
		listener = chain.CreateListener(data)
	}
	listener.Dispatch(data)
}

// 查找责任链中的合适监听者
func (chain *FollowCacheChain) FindListener(data protoreflect.ProtoMessage) ListenerInterface {
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
func (chain *FollowCacheChain) CreateListener(data protoreflect.ProtoMessage) ListenerInterface {
	newListener := chain.listenerPool.Get().(*FollowCacheListener)
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
func (chain *FollowCacheChain) DestroyListener(listener ListenerInterface) {
	// 找到前一个节点（假设 chain.Head 是链表的头部）
	current, ok := listener.(*FollowCacheListener)
	if !ok {
		log.Printf("invalid type: expected *FollowCacheListener")
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
