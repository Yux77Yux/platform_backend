package receiver

import (
	"fmt"
	"log"

	event "github.com/Yux77Yux/platform_backend/generated/common/event"
	messaging "github.com/Yux77Yux/platform_backend/microservices/aggregator/messaging"
)

type MessagequeueInterface = messaging.MessagequeueInterface

const (
	UPDATE_CREATION_ACTION_COUNT = "InteractionCount"

	ADD_VIEW = "ADD_VIEW"
)

var (
	ExchangesConfig = map[string]string{
		event.Exchange_EXCHANGE_ADD_VIEW.String(): "direct",
	}
)

func Init() {
	rabbitMQ := messaging.GetRabbitMQ()
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
		case event.Exchange_EXCHANGE_ADD_VIEW.String():
			go messaging.ListenToQueue(exchange, ADD_VIEW, event.RoutingKey_KEY_ADD_VIEW.String(), addViewProcessor)
		}
	}
}
