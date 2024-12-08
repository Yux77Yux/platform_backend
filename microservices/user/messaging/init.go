package messaging

import (
	"fmt"
	"log"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue"
)

var (
	connStr         string
	ExchangesConfig = map[string]string{
		"register":         "direct",
		"storeUserInCache": "direct",
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

func Init() {
	messageQueue := GetRabbitMQ()
	defer messageQueue.Close()

	if messageQueue == nil {
		log.Printf("error: message queue open failed")
		return
	}
	for exchange, kind := range ExchangesConfig {
		exchangeName := fmt.Sprintf("%s_exchange", exchange)
		queueName := fmt.Sprintf("%s_queue", exchange)
		routeKey := fmt.Sprintf("%s_route", exchange)

		if err := messageQueue.ExchangeDeclare(exchangeName, kind, true, false, false, false, nil); err != nil {
			wiredErr := fmt.Errorf("failed to declare exchange %s : %w", exchange, err)
			log.Printf("error: %v", wiredErr)
		}

		switch exchange {
		// 不同的exchange使用不同函数
		case "register":
			go ListenToQueue(exchangeName, queueName, routeKey, registerProcessor)
		case "storeUserInCache":
			go ListenToQueue(exchangeName, queueName, routeKey, storeUserInCache)
		}
	}
}
