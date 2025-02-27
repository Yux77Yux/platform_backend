package messaging

import (
	"fmt"
	"log"

	event "github.com/Yux77Yux/platform_backend/generated/common/event"
	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue"
)

const (
	UPDATE_CREATION_ACTION_COUNT = "InteractionCount"

	ADD_VIEW = "ADD_VIEW"
)

var (
	connStr         string
	messageQueue    MessagequeueInterface
	ExchangesConfig = map[string]string{
		event.Exchange_EXCHANGE_ADD_VIEW.String(): "direct",
	}
)

func InitStr(_str string) {
	connStr = _str
}

func GetRabbitMQ() MessagequeueInterface {
	var messageQueue MessagequeueInterface = &pkgMQ.RabbitMQClass{}
	if err := messageQueue.Open(connStr); err != nil {
		wiredErr := fmt.Errorf("failed to connect the rabbit client: %w", err)
		log.Printf("error: %v", wiredErr)
		return nil
	}

	return messageQueue
}

func CloseClient() {
	if err := messageQueue.Close(); err != nil {
		log.Printf("error: failed to close message queue client: %v", err)
	}
}
