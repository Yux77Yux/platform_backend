package messaging

import (
	"fmt"
	"log"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue"
)

var (
	connStr string
)

func InitStr(_str string) {
	connStr = _str
}

const (
	UpdateDbCreation    = "UpdateDbCreation"
	StoreCreationInfo   = "StoreCreationInfo"
	UpdateCacheCreation = "UpdateCacheCreation"

	// Review
	PendingCreation      = "PendingCreation"      // 起点
	UpdateCreationStatus = "UpdateCreationStatus" // 终点
	DeleteCreation       = "DeleteCreation"

	// Interaction Aggregator
	UPDATE_CREATION_ACTION_COUNT = "InteractionCount" // 终点
)

func GetRabbitMQ() MessageQueueInterface {
	var messageQueue MessageQueueInterface = &pkgMQ.RabbitMQClass{}
	if err := messageQueue.Open(connStr); err != nil {
		wiredErr := fmt.Errorf("failed to connect the rabbit client: %w", err)
		log.Printf("error: %v", wiredErr)
		return nil
	}

	return messageQueue
}
