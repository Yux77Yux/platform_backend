package receiver

import (
	event "github.com/Yux77Yux/platform_backend/generated/event/aggregator"
)

var (
	ExchangesConfig = map[string]string{
		EXCHANGE_INCREASE_VIEW: "direct",
	}
)

var (
	EXCHANGE_INCREASE_VIEW = event.Exchange_EXCHANGE_INCREASE_VIEW.String()
)

var (
	QUEUE_INCREASE_VIEW = event.Queue_QUEUE_INCREASE_VIEW.String()
)

var (
	KEY_INCREASE_VIEW = event.RoutingKey_KEY_INCREASE_VIEW.String()
)
