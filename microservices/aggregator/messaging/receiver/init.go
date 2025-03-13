package receiver

import (
	event "github.com/Yux77Yux/platform_backend/generated/common/event"
	messaging "github.com/Yux77Yux/platform_backend/microservices/aggregator/messaging"
)

const (
	UPDATE_CREATION_ACTION_COUNT = "InteractionCount"

	ADD_VIEW = "ADD_VIEW"
)

var (
	ExchangesConfig = messaging.ExchangesConfig
)

func Init(addr string) {
	messaging.InitStr(addr)
	for exchange := range ExchangesConfig {
		switch exchange {
		// 不同的exchange使用不同函数
		case event.Exchange_EXCHANGE_ADD_VIEW.String():
			go messaging.ListenToQueue(exchange, ADD_VIEW, event.RoutingKey_KEY_ADD_VIEW.String(), addViewProcessor)
		}
	}
}
