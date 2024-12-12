package messaging

import (
	"fmt"
	"log"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue"
)

var (
	connStr         string
	messageQueue    MessagequeueInterface
	ExchangesConfig = map[string]string{
		"register": "direct",
		// Add more exchanges here
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
