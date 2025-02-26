package messaging

import (
	"fmt"
	"log"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue"
)

const (
	UpdateDbCreation    = "UpdateDbCreation"
	StoreCreationInfo   = "StoreCreationInfo"
	UpdateCacheCreation = "UpdateCacheCreation"

	// Review
	PendingCreation      = "PendingCreation"      // 起点
	UpdateCreationStatus = "UpdateCreationStatus" // 终点
	DeleteCreation       = "DeleteCreation"
)

var (
	connStr         string
	ExchangesConfig = map[string]string{
		UpdateDbCreation:     "direct",
		UpdateCacheCreation:  "direct",
		StoreCreationInfo:    "direct",
		UpdateCreationStatus: "direct",
		DeleteCreation:       "direct",
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
		case UpdateDbCreation:
			go ListenToQueue(exchange, UpdateDbCreation, UpdateDbCreation, updateCreationDbProcessor)
		case StoreCreationInfo:
			go ListenToQueue(exchange, StoreCreationInfo, StoreCreationInfo, storeCreationProcessor)

		case UpdateCreationStatus:
			go ListenToQueue(exchange, UpdateCreationStatus, UpdateCreationStatus, updateCreationStatusProcessor)
		case UpdateCacheCreation:
			go ListenToQueue(exchange, UpdateCacheCreation, UpdateCacheCreation, updateCreationCacheProcessor)
		case DeleteCreation:
			go ListenToQueue(exchange, DeleteCreation, DeleteCreation, deleteCreationProcessor)
		}
	}
}
