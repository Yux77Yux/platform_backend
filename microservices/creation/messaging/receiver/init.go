package receiver

import (
	"fmt"
	"log"

	messaging "github.com/Yux77Yux/platform_backend/microservices/creation/messaging"
)

const (
	UpdateDbCreation    = messaging.UpdateDbCreation
	StoreCreationInfo   = messaging.StoreCreationInfo
	UpdateCacheCreation = messaging.UpdateCacheCreation

	// Review
	PendingCreation      = messaging.PendingCreation      // 起点
	UpdateCreationStatus = messaging.UpdateCreationStatus // 终点
	DeleteCreation       = messaging.DeleteCreation

	// Interaction Aggregator
	UPDATE_CREATION_ACTION_COUNT = messaging.UPDATE_CREATION_ACTION_COUNT // 终点
)

var (
	ExchangesConfig = map[string]string{
		UpdateDbCreation:             "direct",
		UpdateCacheCreation:          "direct",
		StoreCreationInfo:            "direct",
		UpdateCreationStatus:         "direct",
		DeleteCreation:               "direct",
		UPDATE_CREATION_ACTION_COUNT: "direct",
		// Add more exchanges here
	}
)

// 非RPC类型的消息队列的交换机声明
func Init() {
	rabbitMQ := messaging.GetRabbitMQ()
	defer rabbitMQ.Close()

	for exchange, kind := range ExchangesConfig {
		if err := rabbitMQ.ExchangeDeclare(exchange, kind, true, false, false, false, nil); err != nil {
			wiredErr := fmt.Errorf("failed to declare exchange %s : %w", exchange, err)
			log.Printf("error: %v", wiredErr)
		}

		switch exchange {
		// 不同的exchange使用不同函数
		case UpdateDbCreation:
			go messaging.ListenToQueue(exchange, UpdateDbCreation, UpdateDbCreation, updateCreationDbProcessor)
		case StoreCreationInfo:
			go messaging.ListenToQueue(exchange, StoreCreationInfo, StoreCreationInfo, storeCreationProcessor)

		case UpdateCreationStatus:
			go messaging.ListenToQueue(exchange, UpdateCreationStatus, UpdateCreationStatus, updateCreationStatusProcessor)
		case UpdateCacheCreation:
			go messaging.ListenToQueue(exchange, UpdateCacheCreation, UpdateCacheCreation, updateCreationCacheProcessor)
		case DeleteCreation:
			go messaging.ListenToQueue(exchange, DeleteCreation, DeleteCreation, deleteCreationProcessor)
		case UPDATE_CREATION_ACTION_COUNT:
			go messaging.ListenToQueue(exchange, UPDATE_CREATION_ACTION_COUNT, UPDATE_CREATION_ACTION_COUNT, addInteractionCount)
		}
	}
}
