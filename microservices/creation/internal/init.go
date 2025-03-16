package internal

import (
	event "github.com/Yux77Yux/platform_backend/generated/event/creation"
)

var (
	EXCHANGE_DELETE_CREATION = event.Exchange_EXCHANGE_DELETE_CREATION.String()

	KEY_DELETE_CREATION = event.RoutingKey_KEY_DELETE_CREATION.String()
)

var (
	db        SqlMethod
	messaging MessageQueueMethod
	cache     CacheMethod
)

func InitDb(_db SqlMethod) {
	db = _db
}

func InitMQ(_messaging MessageQueueMethod) {
	messaging = _messaging
}

func InitCache(_cache CacheMethod) {
	cache = _cache
}
