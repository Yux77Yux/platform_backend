package messaging

import (
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
)

func InitialListenerChain() *ListenerChain {
	return &ListenerChain{
		Head:  new(CommentListener),
		Count: 0,
	}
}

// 责任链
type ListenerChain struct {
	Head  *CommentListener // 责任链的头部
	Count int32            // 监听者数量
}

// 查找责任链中的合适监听者
func (chain *ListenerChain) FindListenerForCreation(comment *generated.Comment) *CommentListener {
	creationID := comment.GetCreationId()

	current := chain.Head.next
	for current != nil {
		if current.creationID == creationID {
			return current
		}
		current = current.next
	}
	return nil // 没有找到合适的监听者
}

// 销毁监听者
func (chain *ListenerChain) DestroyListener(listener *CommentListener) {
	listener.cleanup()
	listener = listener.next
	chain.Count = chain.Count - 1
}

// 创建一个新的监听者
func (chain *ListenerChain) CreateListenerForCreation(comment *generated.Comment) *CommentListener {
	newListener := &CommentListener{
		creationID:      comment.GetCreationId(),
		commentChannel:  make(chan *generated.Comment, 100),
		timeoutDuration: 10 * time.Second,
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
func (chain *ListenerChain) HandleCommentRequest(comment *generated.Comment) {
	listener := chain.FindListenerForCreation(comment)
	if listener == nil {
		// 如果没有找到合适的监听者，创建一个新的监听者
		listener = chain.CreateListenerForCreation(comment)
	}
	listener.DispatchComment(comment)
}
