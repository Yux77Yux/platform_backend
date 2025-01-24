package dispatch

import (
	"log"
	"sync"
	"sync/atomic"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	cache "github.com/Yux77Yux/platform_backend/microservices/user/cache"
)

/*
	这里的链表不太符合高并发特点的设计，问题在于持有锁时间会很长
	改进的办法是使用堆建立，或者使用HASH对节点进行映射
	先留着，以后再建堆
*/

func InitialRegisterCacheChain() *RegisterCacheChain {
	_chain := &RegisterCacheChain{
		Head:       &RegisterCacheListener{prev: nil},
		Tail:       &RegisterCacheListener{next: nil},
		Count:      0,
		exeChannel: make(chan *[]*generated.UserCredentials, EXE_CHANNEL_COUNT),
	}
	go _chain.ExecuteBatch()
	return _chain
}

// 责任链
type RegisterCacheChain struct {
	Head       *RegisterCacheListener // 责任链的头部
	Tail       *RegisterCacheListener
	nodeMux    sync.Mutex
	Count      int32 // 监听者数量
	exeChannel chan *[]*generated.UserCredentials
}

func (chain *RegisterCacheChain) ComeFromRoot(userCredentials *[]*generated.UserCredentials) {
	chain.exeChannel <- userCredentials
}

func (chain *RegisterCacheChain) ExecuteBatch() {
	for userCredentialsPtr := range chain.exeChannel {
		go func(userCredentialsPtr *[]*generated.UserCredentials) {
			userCredentials := *userCredentialsPtr

			// 插入Redis
			err := cache.StoreCredentials(userCredentials)
			if err != nil {
				log.Printf("error: StoreCredentials error %v", err)
			}

			// 清理结束
			*userCredentialsPtr = userCredentials[:0] // 清空切片内容
			insertUserCredentialsPool.Put(userCredentialsPtr)
		}(userCredentialsPtr)
	}
}

// 处理评论请求的函数
func (chain *RegisterCacheChain) HandleRequest(data protoreflect.ProtoMessage) {
	listener := chain.FindListener(data)
	if listener == nil {
		// 如果没有找到合适的监听者，创建一个新的监听者
		listener = chain.CreateListener(data)
	}
	listener.Dispatch(data)
}

// 查找责任链中的合适监听者
func (chain *RegisterCacheChain) FindListener(data protoreflect.ProtoMessage) ListenerInterface {
	chain.nodeMux.Lock()
	next := chain.Head.next
	prev := chain.Tail.prev
	for {
		if atomic.LoadUint32(&next.count) == 50 {
			chain.nodeMux.Unlock()
			return next
		}
		if atomic.LoadUint32(&prev.count) == 50 {
			chain.nodeMux.Unlock()
			return prev
		}
		if prev == next {
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
func (chain *RegisterCacheChain) CreateListener(data protoreflect.ProtoMessage) ListenerInterface {
	newListener := registerCacheListenerPool.Get().(*RegisterCacheListener)
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
func (chain *RegisterCacheChain) DestroyListener(listener ListenerInterface) {
	// 找到前一个节点（假设 chain.Head 是链表的头部）
	current, ok := listener.(*RegisterCacheListener)
	if !ok {
		log.Printf("invalid type: expected *RegisterCacheListener")
	}

	chain.nodeMux.Lock()
	prev := current.prev
	next := current.next
	prev.next = next
	next.prev = prev
	chain.nodeMux.Unlock()
	atomic.AddInt32(&chain.Count, -1)

	listener.Cleanup()
}
