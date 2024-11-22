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
		"register_exchange": "direct",
		// Add more exchanges here
	}
)

func InitStr(_str string) {
	connStr = _str
}

func GetRabbitMQ() MessagequeueInterface {
	var messageQueue MessagequeueInterface = &pkgMQ.RabbitMQClass{}
	err := messageQueue.Open(connStr)
	wiredErr := fmt.Errorf("failed to connect the rabbit client: %w", err)
	log.Printf("error: %v", wiredErr)

	return messageQueue
}

func Init() {
	messageQueue = GetRabbitMQ()

	for exchange, kind := range ExchangesConfig {
		err := messageQueue.ExchangeDeclare(exchange, kind, true, false, false, false, nil)
		wiredErr := fmt.Errorf("failed to declare exchange %s : %w", exchange, err)
		log.Printf("error: %v", wiredErr)
	}
}
