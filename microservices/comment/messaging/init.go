package messaging

import (
	"fmt"
	"log"
	"sync"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue"
)

var (
	deleteChain     *DeleteChain
	delListenerPool = sync.Pool{
		New: func() any {
			return &DeleteListener{
				commentChannel:  make(chan *generated.AfterAuth, 160),
				timeoutDuration: 12 * time.Second,
				updateInterval:  3 * time.Second,
			}
		},
	}
	delCommentsPool = sync.Pool{
		New: func() any {
			return make([]*generated.AfterAuth, 0, 50)
		},
	}
	insertChain        *InsertChain
	insertListenerPool = sync.Pool{
		New: func() any {
			return &InsertListener{
				commentChannel:  make(chan *generated.Comment, 160),
				timeoutDuration: 12 * time.Second,
				updateInterval:  3 * time.Second,
			}
		},
	}
	commentsPool = sync.Pool{
		New: func() any {
			return make([]*generated.Comment, 0, 50)
		},
	}

	connStr         string
	ExchangesConfig = map[string]string{
		"PublishComment": "direct",
		"DeleteComment":  "direct",
		// Add more exchanges here
	}
)

func InitStr(_str string) {
	connStr = _str
}

func GetRabbitMQ() MessageQueueInterface {
	var messageQueue MessageQueueInterface = &pkgMQ.RabbitMQClass{}
	if err := messageQueue.Open(connStr); err != nil {
		wiredErr := fmt.Errorf("failed to connect the rabbit client: %w", err)
		log.Printf("error: %v", wiredErr)
		return nil
	}

	return messageQueue
}

// 非RPC类型的消息队列的交换机声明
func Init() {
	// 初始化责任链
	insertChain = InitialInsertChain()
	deleteChain = InitialDeleteChain()

	// 初始化 消息队列 交换机
	rabbitMQ := GetRabbitMQ()
	defer rabbitMQ.Close()

	if rabbitMQ == nil {
		log.Printf("error: message queue open failed")
		return
	}
	for exchange, kind := range ExchangesConfig {
		if err := rabbitMQ.ExchangeDeclare(exchange, kind, true, false, false, false, nil); err != nil {
			wiredErr := fmt.Errorf("failed to declare exchange %s : %w", exchange, err)
			log.Printf("error: %v", wiredErr)
		}

		switch exchange {
		// 不同的exchange使用不同函数
		case "PublishComment":
			go ListenToQueue(exchange, "PublishComment", "PublishComment", JoinCommentProcessor)
		case "DeleteComment":
			go ListenToQueue(exchange, "DeleteComment", "DeleteComment", DeleteCommentProcessor)
		}
	}
}
