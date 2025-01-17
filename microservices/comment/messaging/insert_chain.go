package messaging

import (
	"log"
	"time"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
)

func InitialInsertChain() *InsertChain {
	return &InsertChain{
		Head:  new(InsertListener),
		Count: 0,
	}
}

// 责任链
type InsertChain struct {
	Head  *InsertListener // 责任链的头部
	Count int32           // 监听者数量
}

// 查找责任链中的合适监听者
func (chain *InsertChain) FindListenerForUnique(data []byte) ListenerInterface {
	comment := new(generated.Comment)
	err := proto.Unmarshal(data, comment)
	if err != nil {
		log.Printf("error: FindListenerForUnique Unmarshal :%v", err)
	}

	creationID := comment.GetCreationId()

	current := chain.Head.Next().(*InsertListener)
	for current != nil {
		if current.creationID == creationID {
			return current
		}
		current = current.next
	}
	return nil // 没有找到合适的监听者
}

// 销毁监听者
func (chain *InsertChain) DestroyListener(listener ListenerInterface) {
	// 找到前一个节点（假设 chain.Head 是链表的头部）
	current := chain.Head
	for current != nil && current.Next() != listener {
		current = current.next // 遍历链表找到 listener 的前一个节点
	}

	if current != nil {
		// 找到前一个节点后，跳过 listener
		current.next = current.next.next // 删除 listener，调整链表
		chain.Count = chain.Count - 1
	}

	listener.cleanup()
	chain.Count = chain.Count - 1
}

// 创建一个新的监听者
func (chain *InsertChain) CreateListenerForUnique(data []byte) ListenerInterface {
	comment := new(generated.Comment)
	err := proto.Unmarshal(data, comment)
	if err != nil {
		log.Printf("error: CreateListenerForUnique Unmarshal :%v", err)
	}

	newListener := &InsertListener{
		creationID:      comment.GetCreationId(),
		commentChannel:  make(chan *generated.Comment, 50),
		timeoutDuration: 8 * time.Second,
		updateInterval:  5 * time.Second,
	}

	// 头插法，将新的监听者挂到链中
	newListener.next = chain.Head.next
	chain.Head.next = newListener
	chain.Count = chain.Count + 1

	newListener.StartListening() // 启动监听
	return newListener
}

// 处理评论请求的函数
func (chain *InsertChain) HandleRequest(data []byte) {
	listener := chain.FindListenerForUnique(data)
	if listener == nil {
		// 如果没有找到合适的监听者，创建一个新的监听者
		listener = chain.CreateListenerForUnique(data)
	}
	listener.Dispatch(data)
}
