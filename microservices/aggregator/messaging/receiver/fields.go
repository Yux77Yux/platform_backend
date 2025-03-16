package receiver

import (
	event "github.com/Yux77Yux/platform_backend/generated/event/aggregator"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/aggregator/messaging/dispatch"
)

var (
	AddView         = dispatch.AddView
	ExchangesConfig = map[string]string{
		EXCHANGE_ADD_VIEW: "direct",
	}
)

var (
	EXCHANGE_ADD_VIEW = event.Exchange_EXCHANGE_ADD_VIEW.String()
)

var (
	QUEUE_ADD_VIEW = event.Queue_QUEUE_ADD_VIEW.String()
)

var (
	KEY_ADD_VIEW = event.RoutingKey_KEY_ADD_VIEW.String()
)
