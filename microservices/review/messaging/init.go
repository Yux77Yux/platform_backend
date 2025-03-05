package messaging

import (
	"fmt"
	"log"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue"
)

const (
	New_review      = "New_review"
	Comment_review  = "Comment_review"
	User_review     = "User_review"
	Creation_review = "Creation_review"
	PendingCreation = "PendingCreation"

	Update      = "Update"
	BatchUpdate = "BatchUpdate"

	// USER
	USER_APPROVE  = "UpdateUserStatus"
	USER_REJECTED = "UpdateUserStatus"

	// CREATION
	CREATION_APPROVE  = "UpdateCreationStatus"
	CREATION_REJECTED = "UpdateCreationStatus"
	CREATION_DELETED  = "DeleteCreation"

	// COMMENT
	COMMENT_REJECTED = "DeleteComment"
	COMMENT_DELETED  = "DeleteComment"
)

var (
	connStr         string
	ExchangesConfig = map[string]string{
		New_review:      "direct",
		PendingCreation: "direct",
		Update:          "direct",
		BatchUpdate:     "direct",
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
		case New_review:
			go ListenToQueue(exchange, New_review, New_review, NewReviewProcessor)
		case PendingCreation:
			go ListenToQueue(exchange, PendingCreation, PendingCreation, PendingCreationProcessor)
		case BatchUpdate:
			go ListenToQueue(exchange, BatchUpdate, BatchUpdate, BatchUpdateProcessor)
		case Update:
			go ListenToQueue(exchange, Update, Update, UpdateProcessor)
		}
	}
}
