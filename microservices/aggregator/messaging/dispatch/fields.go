package dispatch

import (
	event "github.com/Yux77Yux/platform_backend/generated/event/aggregator"
)

var (
	AddView = "AddView"
)

var (
	EXCHANGE_ADD_VIEW                     = event.Exchange_EXCHANGE_ADD_VIEW.String()
	EXCHANGE_UPDATE_CREATION_ACTION_COUNT = event.Exchange_EXCHANGE_UPDATE_CREATION_ACTION_COUNT.String()
)

var (
	QUEUE_ADD_VIEW = event.Queue_QUEUE_ADD_VIEW.String()
)

var (
	KEY_ADD_VIEW                     = event.RoutingKey_KEY_ADD_VIEW.String()
	KEY_UPDATE_CREATION_ACTION_COUNT = event.RoutingKey_KEY_UPDATE_CREATION_ACTION_COUNT.String()
)
