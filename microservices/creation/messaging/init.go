package messaging

import (
	"fmt"
	"log"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue"
)

const (
	DraftCreation        = "DraftCreation"
	PendingCreation      = "PendingCreation"
	UpdateCreation       = "UpdateCreation"
	UpdateCreationStatus = "UpdateCreationStatus"
)

var (
	connStr         string
	ExchangesConfig = map[string]string{
		DraftCreation:   "direct",
		PendingCreation: "direct",
		UpdateCreation:  "direct",
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
		case PendingCreation:
			go ListenToQueue(exchange, PendingCreation, PendingCreation, pendingCreationProcessor)
		case DraftCreation:
			go ListenToQueue(exchange, DraftCreation, DraftCreation, draftCreationProcessor)
		case UpdateCreation:
			go ListenToQueue(exchange, UpdateCreation, UpdateCreation, updateCreationProcessor)
		case UpdateCreationStatus:
			go ListenToQueue(exchange, UpdateCreationStatus, UpdateCreationStatus, updateCreationStatusProcessor)
		}
	}

	// for _, exchange := range ListenRPCs {
	// 	switch exchange {
	// 	// 不同的exchange使用不同函数
	// 	case "agg_user":
	// 		go ListenRPC(exchange, "getUser", "getUser", getUserProcessor)
	// 	}
}
