package dispatch

import (
	event "github.com/Yux77Yux/platform_backend/generated/event/aggregator"
)

const (
	LISTENER_CHANNEL_COUNT = 120
	MAX_BATCH_SIZE         = 50
	EXE_CHANNEL_COUNT      = 5
)

const (
	AddView = "AddView"
)

var (
	EXCHANGE_ADD_VIEW                     = event.Exchange_EXCHANGE_ADD_VIEW.String()
	EXCHANGE_UPDATE_CREATION_ACTION_COUNT = event.Exchange_EXCHANGE_UPDATE_CREATION_ACTION_COUNT.String()
)

var (
	KEY_ADD_VIEW                     = event.RoutingKey_KEY_ADD_VIEW.String()
	KEY_UPDATE_CREATION_ACTION_COUNT = event.RoutingKey_KEY_UPDATE_CREATION_ACTION_COUNT.String()
)
