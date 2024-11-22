package messaging

import (
	"fmt"
	"log"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue"
)

var (
	connStr         string
	rabbitMQ        MessagequeueInterface
	ExchangesConfig = map[string]string{
		"register_exchange": "direct",
		// Add more exchanges here
	}
)

func InitStr(_str string) {
	connStr = _str
}

func GetRabbitMQ() MessagequeueInterface {
	var rabbitMQ MessagequeueInterface = &pkgMQ.RabbitMQClass{}
	err := rabbitMQ.Open(connStr)
	wiredErr := fmt.Errorf("failed to connect the rabbit client: %w", err)
	log.Printf("error: %v", wiredErr)

	return rabbitMQ
}

func Init() {
	rabbitMQ = GetRabbitMQ()

	for exchange, kind := range ExchangesConfig {
		err := rabbitMQ.ExchangeDeclare(exchange, kind, true, false, false, false, nil)
		wiredErr := fmt.Errorf("failed to declare exchange %s : %w", exchange, err)
		log.Printf("error: %v", wiredErr)
	}
}
