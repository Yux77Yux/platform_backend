package receiver

import (
	event "github.com/Yux77Yux/platform_backend/generated/event/creation"
)

var (
	EXCHANGE_DELETE_CREATION              = event.Exchange_EXCHANGE_DELETE_CREATION.String()
	EXCHANGE_PEND_CREATION                = event.Exchange_EXCHANGE_PEND_CREATION.String()
	EXCHANGE_STORE_CREATION               = event.Exchange_EXCHANGE_STORE_CREATION.String()
	EXCHANGE_UPDATE_CACHE_CREATION        = event.Exchange_EXCHANGE_UPDATE_CACHE_CREATION.String()
	EXCHANGE_UPDATE_CREATION_ACTION_COUNT = event.Exchange_EXCHANGE_UPDATE_CREATION_ACTION_COUNT.String()
	EXCHANGE_UPDATE_DB_CREATION           = event.Exchange_EXCHANGE_UPDATE_DB_CREATION.String()
	EXCHANGE_UPDATE_CREATION_STATUS       = event.Exchange_EXCHANGE_UPDATE_CREATION_STATUS.String()
)

var (
	QUEUE_DELETE_CREATION              = event.Queue_QUEUE_DELETE_CREATION.String()
	QUEUE_STORE_CREATION               = event.Queue_QUEUE_STORE_CREATION.String()
	QUEUE_UPDATE_CACHE_CREATION        = event.Queue_QUEUE_UPDATE_CACHE_CREATION.String()
	QUEUE_UPDATE_CREATION_ACTION_COUNT = event.Queue_QUEUE_UPDATE_CREATION_ACTION_COUNT.String()
	QUEUE_UPDATE_DB_CREATION           = event.Queue_QUEUE_UPDATE_DB_CREATION.String()
	QUEUE_UPDATE_CREATION_STATUS       = event.Queue_QUEUE_UPDATE_CREATION_STATUS.String()
)

var (
	KEY_DELETE_CREATION              = event.RoutingKey_KEY_DELETE_CREATION.String()
	KEY_PEND_CREATION                = event.RoutingKey_KEY_PEND_CREATION.String()
	KEY_STORE_CREATION               = event.RoutingKey_KEY_STORE_CREATION.String()
	KEY_UPDATE_CACHE_CREATION        = event.RoutingKey_KEY_UPDATE_CACHE_CREATION.String()
	KEY_UPDATE_CREATION_ACTION_COUNT = event.RoutingKey_KEY_UPDATE_CREATION_ACTION_COUNT.String()
	KEY_UPDATE_DB_CREATION           = event.RoutingKey_KEY_UPDATE_DB_CREATION.String()
	KEY_UPDATE_CREATION_STATUS       = event.RoutingKey_KEY_UPDATE_CREATION_STATUS.String()
)

var (
	ExchangesConfig = map[string]string{
		EXCHANGE_DELETE_CREATION:              "direct",
		EXCHANGE_STORE_CREATION:               "direct",
		EXCHANGE_UPDATE_CACHE_CREATION:        "direct",
		EXCHANGE_UPDATE_CREATION_ACTION_COUNT: "direct",
		EXCHANGE_UPDATE_DB_CREATION:           "direct",
		EXCHANGE_UPDATE_CREATION_STATUS:       "direct",
	}
)
