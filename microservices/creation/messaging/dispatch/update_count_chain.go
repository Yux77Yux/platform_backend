package dispatch

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	db "github.com/Yux77Yux/platform_backend/microservices/creation/repository"
)

type ExeBody struct {
	id           int64
	newLikeCount int32
	newSaveCount int32
	newViewCount int32
}

func InitialUpdateCountChain() *UpdateCountChain {
	_chain := &UpdateCountChain{
		Count:      0,
		exeChannel: make(chan *ExeBody, EXE_CHANNEL_COUNT),
		listenerPool: sync.Pool{
			New: func() any {
				return &UpdateCountListener{
					timeoutDuration: 3 * time.Minute,
					updateInterval:  10 * time.Second,
				}
			},
		},
	}

	go _chain.ExecuteBatch()
	return _chain
}

// 责任链
type UpdateCountChain struct {
	listenerMap  sync.Map
	Count        int32 // 监听者数量
	exeChannel   chan *ExeBody
	listenerPool sync.Pool
}

func (chain *UpdateCountChain) ExecuteBatch() {
	for UpdateCountsPtr := range chain.exeChannel {
		go func(UpdateCountsPtr *ExeBody) {
			counts := UpdateCountsPtr

			// 插入数据库
			err := db.UpdateCreationCount(
				context.Background(),
				counts.id,
				counts.newSaveCount,
				counts.newLikeCount,
				counts.newViewCount,
			)
			if err != nil {
				log.Printf("error: UpdateCreationCount %v", err)
				// 死信，没做
			}
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
	action, ok := data.(*common.UserAction)
	if !ok {
		log.Printf(": expected *common.UserAction")
	}

	creationId := action.GetId().GetId()

	// 尝试从 listenerMap 中获取监听者
	if listener, exist := chain.listenerMap.Load(creationId); exist {
		return listener.(ListenerInterface)
	}

	return nil
}

// 创建一个新的监听者
func (chain *UpdateCountChain) CreateListener(data protoreflect.ProtoMessage) ListenerInterface {
	action, ok := data.(*common.UserAction)
	if !ok {
		log.Printf(": expected *common.UserAction")
	}
	creationId := action.GetId().GetId()

	// 如果不存在，从对象池获取新的监听者
	newListener, ok := chain.listenerPool.Get().(*UpdateCountListener)
	if !ok {
		log.Printf("FindListener: failed to get listener from pool")
		return nil
	}
	atomic.StoreInt64(&newListener.id, creationId)

	// 存入 map 中
	actual, _ := chain.listenerMap.LoadOrStore(creationId, newListener)

	// 增加计数
	atomic.AddInt32(&chain.Count, 1)

	// `actual` 是 map 中实际存储的监听者，如果别的协程抢先存入，`newListener` 不会被使用
	return actual.(ListenerInterface)
}

// 销毁监听者
func (chain *UpdateCountChain) DestroyListener(listener ListenerInterface) {
	// 找到前一个节点（假设 chain.Head 是链表的头部）
	current, ok := listener.(*UpdateCountListener)
	if !ok {
		log.Printf("invalid type: expected *UpdateCountListener")
	}

	id := current.GetId()
	chain.listenerMap.Delete(id)

	atomic.AddInt32(&chain.Count, -1)
	chain.listenerPool.Put(listener)
}
