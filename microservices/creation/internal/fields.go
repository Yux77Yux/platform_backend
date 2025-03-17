package internal

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
	KEY_DELETE_CREATION              = event.RoutingKey_KEY_DELETE_CREATION.String()
	KEY_PEND_CREATION                = event.RoutingKey_KEY_PEND_CREATION.String()
	KEY_STORE_CREATION               = event.RoutingKey_KEY_STORE_CREATION.String()
	KEY_UPDATE_CACHE_CREATION        = event.RoutingKey_KEY_UPDATE_CACHE_CREATION.String()
	KEY_UPDATE_CREATION_ACTION_COUNT = event.RoutingKey_KEY_UPDATE_CREATION_ACTION_COUNT.String()
	KEY_UPDATE_DB_CREATION           = event.RoutingKey_KEY_UPDATE_DB_CREATION.String()
	KEY_UPDATE_CREATION_STATUS       = event.RoutingKey_KEY_UPDATE_CREATION_STATUS.String()
)
