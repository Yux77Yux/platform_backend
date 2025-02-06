package dispatch

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
)

func InitialCancelCollectionCacheChain() *CancelCollectionCacheChain {
	_chain := &CancelCollectionCacheChain{
		Head:       &CancelCollectionListener{prev: nil},
		Tail:       &CancelCollectionListener{next: nil},
		Count:      0,
		exeChannel: make(chan *[]*generated.BaseInteraction, EXE_CHANNEL_COUNT),
		listenerPool: sync.Pool{
			New: func() any {
				return &CancelCollectionListener{
					datasChannel:    make(chan *generated.BaseInteraction, LISTENER_CHANNEL_COUNT),
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
type CancelCollectionCacheChain struct {
	Head         *CancelCollectionListener // 责任链的头部
	Tail         *CancelCollectionListener
	Count        int32 // 监听者数量
	nodeMux      sync.Mutex
	exeChannel   chan *[]*generated.BaseInteraction
	listenerPool sync.Pool
}

func (chain *CancelCollectionCacheChain) ExecuteBatch() {
	log.Printf("我他妈来啦!!！ ")
	for interactionsPtr := range chain.exeChannel {
		go func(interactionsPtr *[]*generated.BaseInteraction) {
			interactions := *interactionsPtr
			log.Printf("我他妈来啦！ %v", interactions)
			// 插入数据库
			err := cache.DelCollections(interactions)
			if err != nil {
				log.Printf("error: DelCollections error")
			}

			// 放回对象池
			*interactionsPtr = interactions[:0]
			baseInteractionsPool.Put(interactionsPtr)
		}(interactionsPtr)
	}
}

// 处理评论请求的函数
func (chain *CancelCollectionCacheChain) HandleRequest(data protoreflect.ProtoMessage) {
	listener := chain.FindListener(data)
	if listener == nil {
		// 如果没有找到合适的监听者，创建一个新的监听者
		listener = chain.CreateListener(data)
	}
	listener.Dispatch(data)
}

// 查找责任链中的合适监听者
func (chain *CancelCollectionCacheChain) FindListener(data protoreflect.ProtoMessage) ListenerInterface {
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
func (chain *CancelCollectionCacheChain) CreateListener(data protoreflect.ProtoMessage) ListenerInterface {
	newListener := chain.listenerPool.Get().(*CancelCollectionListener)
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
func (chain *CancelCollectionCacheChain) DestroyListener(listener ListenerInterface) {
	// 找到前一个节点（假设 chain.Head 是链表的头部）
	current, ok := listener.(*CancelCollectionListener)
	if !ok {
		log.Printf("invalid type: expected *CancelCollectionListener")
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
