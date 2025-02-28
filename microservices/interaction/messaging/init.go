package messaging

import (
	"fmt"
	"log"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue"
)

const (
	ComputeSimilarCreation = "ComputeSimilarCreation"
	ComputeUser            = "ComputeUser"

	UpdateDb      = "UpdateDb"
	BatchUpdateDb = "BatchUpdateDb"
	AddCollection = "AddCollection"
	AddLike       = "AddLike"
	AddView       = "AddView"
	CancelLike    = "CancelLike"

	// Creation
	UPDATE_CREATION_ACTION_COUNT = "InteractionCount"
)

var (
	connStr string
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
