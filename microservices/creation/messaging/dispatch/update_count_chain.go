package dispatch

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	messaging "github.com/Yux77Yux/platform_backend/microservices/aggregator/messaging"
)

func InitialUpdateCountChain() *UpdateCountChain {
	_chain := &UpdateCountChain{
		Head:       &UpdateCountListener{prev: nil},
		Tail:       &UpdateCountListener{next: nil},
		Count:      0,
		exeChannel: make(chan *[]*common.UserAction, EXE_CHANNEL_COUNT),
		listenerPool: sync.Pool{
			New: func() any {
				return &UpdateCountListener{
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
type UpdateCountChain struct {
	Head         *UpdateCountListener // 责任链的头部
	Tail         *UpdateCountListener
	Count        int32 // 监听者数量
	nodeMux      sync.Mutex
	exeChannel   chan *[]*common.UserAction
	listenerPool sync.Pool
}

func (chain *UpdateCountChain) ExecuteBatch() {
	for UpdateCountsPtr := range chain.exeChannel {
		go func(UpdateCountsPtr *[]*common.UserAction) {
			views := *UpdateCountsPtr
			anyViews := &common.AnyUserAction{
				Actions: views,
			}
			// 插入数据库
			err := messaging.SendMessage(UpdateCount, UpdateCount, anyViews)
			if err != nil {
				log.Printf("error: SendMessage UpdateCount error")
			}

			// 放回对象池
			*UpdateCountsPtr = views[:0]
			insertPool.Put(UpdateCountsPtr)
		}(UpdateCountsPtr)
	}
}

// 处理评论请求的函数
func (chain *UpdateCountChain) HandleRequest(data protoreflect.ProtoMessage) {
	listener := chain.FindListener(data)
	if listener == nil {
		// 如果没有找到合适的监听者，创建一个新的监听者
		listener = chain.CreateListener(data)
	}
	listener.Dispatch(data)
}

// 查找责任链中的合适监听者
func (chain *UpdateCountChain) FindListener(data protoreflect.ProtoMessage) ListenerInterface {
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
func (chain *UpdateCountChain) CreateListener(data protoreflect.ProtoMessage) ListenerInterface {
	newListener := chain.listenerPool.Get().(*UpdateCountListener)
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
func (chain *UpdateCountChain) DestroyListener(listener ListenerInterface) {
	// 找到前一个节点（假设 chain.Head 是链表的头部）
	current, ok := listener.(*UpdateCountListener)
	if !ok {
		log.Printf("invalid type: expected *UpdateCountListener")
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
