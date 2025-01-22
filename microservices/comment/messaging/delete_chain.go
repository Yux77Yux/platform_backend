package messaging

import (
	"log"
	"sync"
	"sync/atomic"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
)

func InitialDeleteChain() *DeleteChain {
	_chain := &DeleteChain{
		Head:       new(DeleteListener),
		Count:      0,
		exeChannel: make(chan []*generated.AfterAuth, 10),
	}
	go _chain.executeBatch()
	return _chain
}

// 责任链
type DeleteChain struct {
	Head       *DeleteListener // 责任链的头部
	nodeMux    sync.RWMutex
	Count      int32 // 监听者数量
	exeChannel chan []*generated.AfterAuth
}

func (chain *DeleteChain) executeBatch() {
	for delComments := range chain.exeChannel {
		go func(delComments []*generated.AfterAuth) {
			count := len(delComments)
			id := delComments[0].GetCreationId()
			err := db.BatchUpdateDeleteStatus(delComments)
			if err != nil {
				log.Printf("error: BatchUpdateDeleteStatus error")
			}

			delCommentsPool.Put(delComments)
		}(delComments)
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
	comment := data.(*generated.AfterAuth)

	creationId := comment.GetCreationId()

	current := chain.Head.next
	for current != nil {
		if current.creationId == creationId {
			return current
		}
		current = current.next
	}
	return nil // 没有找到合适的监听者
}

// 创建一个新的监听者
func (chain *DeleteChain) CreateListener(data protoreflect.ProtoMessage) ListenerInterface {
	comment := data.(*generated.AfterAuth)

	newListener := delListenerPool.Get().(*DeleteListener)
	newListener.creationId = comment.GetCreationId()
	newListener.exeChannel = chain.exeChannel

	// 头插法，将新的监听者挂到链中
	newListener.next = chain.Head.next
	chain.Head.next = newListener
	atomic.AddInt32(&chain.Count, 1)

	newListener.StartListening() // 启动监听
	return newListener
}

// 销毁监听者
func (chain *DeleteChain) DestroyListener(listener ListenerInterface) {
	// 找到前一个节点（假设 chain.Head 是链表的头部）
	current := chain.Head
	chain.nodeMux.RLock()
	for current != nil && current.Next() != listener {
		current = current.next // 遍历链表找到 listener 的前一个节点
	}
	chain.nodeMux.RUnlock()

	if current != nil {
		// 找到前一个节点后，跳过 listener
		chain.nodeMux.Lock()
		current.next = current.next.next // 删除 listener，调整链表
		chain.nodeMux.Unlock()
		atomic.AddInt32(&chain.Count, -1)
	} else {
		log.Printf("error: not found the listener %d", listener.GetId())
	}

	listener.cleanup()
}
