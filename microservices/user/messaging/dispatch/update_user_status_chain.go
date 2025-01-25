package dispatch

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	db "github.com/Yux77Yux/platform_backend/microservices/user/repository"
)

/*
	这里的链表不太符合高并发特点的设计，问题在于持有锁时间会很长
	改进的办法是使用堆建立，或者使用HASH对节点进行映射
	先留着，以后再建堆
*/

func InitialUserStatusChain() *UserStatusChain {
	_chain := &UserStatusChain{
		Head:       &UserStatusListener{prev: nil},
		Tail:       &UserStatusListener{next: nil},
		Count:      0,
		exeChannel: make(chan *[]*generated.UserUpdateStatus, EXE_CHANNEL_COUNT),
		listenerPool: sync.Pool{
			New: func() any {
				return &UserStatusListener{
					userUpdateStatusChannel: make(chan *generated.UserUpdateStatus, LISTENER_CHANNEL_COUNT),
					timeoutDuration:         10 * time.Second,
					updateInterval:          3 * time.Second,
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
type UserStatusChain struct {
	Head         *UserStatusListener // 责任链的头部
	Tail         *UserStatusListener
	nodeMux      sync.Mutex
	Count        int32 // 监听者数量
	exeChannel   chan *[]*generated.UserUpdateStatus
	listenerPool sync.Pool
}

func (chain *UserStatusChain) ExecuteBatch() {
	for usersPtr := range chain.exeChannel {
		go func(usersPtr *[]*generated.UserUpdateStatus) {
			users := *usersPtr
			// 更新头像
			err := db.UserUpdateStatusInTransaction(users)
			if err != nil {
				log.Printf("error: UserUpdateStatusInTransaction error")
			}

			*usersPtr = users[:0]
			userStatusPool.Put(usersPtr)
		}(usersPtr)
	}
}

// 处理评论请求的函数
func (chain *UserStatusChain) HandleRequest(data protoreflect.ProtoMessage) {
	listener := chain.FindListener(data)
	if listener == nil {
		// 如果没有找到合适的监听者，创建一个新的监听者
		listener = chain.CreateListener(data)
	}
	listener.Dispatch(data)
}

// 查找责任链中的合适监听者
func (chain *UserStatusChain) FindListener(data protoreflect.ProtoMessage) ListenerInterface {
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
func (chain *UserStatusChain) CreateListener(data protoreflect.ProtoMessage) ListenerInterface {
	newListener := chain.listenerPool.Get().(*UserStatusListener)
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
func (chain *UserStatusChain) DestroyListener(listener ListenerInterface) {
	// 找到前一个节点（假设 chain.Head 是链表的头部）
	current, ok := listener.(*UserStatusListener)
	if !ok {
		log.Printf("invalid type: expected *UserStatusListener")
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
