package dispatch

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	"github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/interaction/repository"
)

func InitialDbChain() *DbInteractionChain {
	_chain := &DbInteractionChain{
		Head:       &DbInteractionsListener{prev: nil},
		Tail:       &DbInteractionsListener{next: nil},
		Count:      0,
		exeChannel: make(chan *[]*generated.OperateInteraction, EXE_CHANNEL_COUNT),
		listenerPool: sync.Pool{
			New: func() any {
				return &DbInteractionsListener{
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
type DbInteractionChain struct {
	Head         *DbInteractionsListener // 责任链的头部
	Tail         *DbInteractionsListener
	Count        int32 // 监听者数量
	nodeMux      sync.Mutex
	exeChannel   chan *[]*generated.OperateInteraction
	listenerPool sync.Pool
}

func (chain *DbInteractionChain) ExecuteBatch() {
	for interactionsPtr := range chain.exeChannel {
		go func(interactionsPtr *[]*generated.OperateInteraction) {
			interactions := *interactionsPtr
			// 插入数据库
			err := db.UpdateInteractions(interactions)
			if err != nil {
				log.Printf("error: UpdateInteractions error %v", err)
				// 死信，但没做
			}

			// 发到消息队列，异步更新数据库中的likes，saves
			actions := make([]*common.UserAction, 0, len(interactions))
			for _, interaction := range interactions {
				action := interaction.GetAction()
				if action == common.Operate_VIEW || action == common.Operate_DEL_VIEW {
					continue
				} else {
					actions = append(actions, &common.UserAction{
						Id: &common.CreationId{
							Id: interaction.GetBase().GetCreationId(),
						},
						Operate: action,
					})
				}
			}
			go func() {
				if len(actions) <= 0 {
					return
				}
				messagingErr := messaging.SendMessage(context.Background(), messaging.UPDATE_CREATION_ACTION_COUNT, messaging.UPDATE_CREATION_ACTION_COUNT, &common.AnyUserAction{
					Actions: actions,
				})
				if messagingErr != nil {
					log.Printf("error: messaging.UPDATE_CREATION_ACTION_COUNT %v", messagingErr)
				}
			}()

			*interactionsPtr = interactions[:0]
			interactionsPool.Put(interactionsPtr)
		}(interactionsPtr)
	}
}

// 处理评论请求的函数
func (chain *DbInteractionChain) HandleRequest(data protoreflect.ProtoMessage) {
	listener := chain.FindListener(data)
	if listener == nil {
		// 如果没有找到合适的监听者，创建一个新的监听者
		listener = chain.CreateListener(data)
	}
	listener.Dispatch(data)
}

// 查找责任链中的合适监听者
func (chain *DbInteractionChain) FindListener(data protoreflect.ProtoMessage) ListenerInterface {
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
func (chain *DbInteractionChain) CreateListener(data protoreflect.ProtoMessage) ListenerInterface {
	newListener := chain.listenerPool.Get().(*DbInteractionsListener)
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
func (chain *DbInteractionChain) DestroyListener(listener ListenerInterface) {
	// 找到前一个节点（假设 chain.Head 是链表的头部）
	current, ok := listener.(*DbInteractionsListener)
	if !ok {
		log.Printf("invalid type: expected *DbInteractionsListener")
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
