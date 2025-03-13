package dispatch

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	messaging "github.com/Yux77Yux/platform_backend/microservices/aggregator/messaging"
)

func InitialAddViewChain() *AddViewChain {
	_chain := &AddViewChain{
		Head:       &AddViewListener{prev: nil},
		Tail:       &AddViewListener{next: nil},
		Count:      0,
		exeChannel: make(chan *[]*common.UserAction, EXE_CHANNEL_COUNT),
		listenerPool: sync.Pool{
			New: func() any {
				return &AddViewListener{
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
type AddViewChain struct {
	Head         *AddViewListener // 责任链的头部
	Tail         *AddViewListener
	Count        int32 // 监听者数量
	nodeMux      sync.Mutex
	exeChannel   chan *[]*common.UserAction
	listenerPool sync.Pool
}

func (chain *AddViewChain) ExecuteBatch() {
	for AddViewsPtr := range chain.exeChannel {
		go func(AddViewsPtr *[]*common.UserAction) {
			views := *AddViewsPtr
			anyViews := &common.AnyUserAction{
				Actions: views,
			}
			// 插入数据库
			err := messaging.SendMessage(context.Background(), AddView, AddView, anyViews)
			if err != nil {
				log.Printf("error: SendMessage AddView error")
			}

			// 放回对象池
			*AddViewsPtr = views[:0]
			insertPool.Put(AddViewsPtr)
		}(AddViewsPtr)
	}
}

// 处理评论请求的函数
func (chain *AddViewChain) HandleRequest(data protoreflect.ProtoMessage) {
	listener := chain.FindListener(data)
	if listener == nil {
		// 如果没有找到合适的监听者，创建一个新的监听者
		listener = chain.CreateListener(data)
	}
	listener.Dispatch(data)
}

// 查找责任链中的合适监听者
func (chain *AddViewChain) FindListener(data protoreflect.ProtoMessage) ListenerInterface {
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
func (chain *AddViewChain) CreateListener(data protoreflect.ProtoMessage) ListenerInterface {
	newListener := chain.listenerPool.Get().(*AddViewListener)
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
func (chain *AddViewChain) DestroyListener(listener ListenerInterface) {
	// 找到前一个节点（假设 chain.Head 是链表的头部）
	current, ok := listener.(*AddViewListener)
	if !ok {
		log.Printf("invalid type: expected *AddViewListener")
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
