package dispatch

import (
	event "github.com/Yux77Yux/platform_backend/generated/event/comment"
)

const (
	Insert = "Insert"
	Delete = "Delete"

	LISTENER_CHANNEL_COUNT = 120
	MAX_BATCH_SIZE         = 50
	EXE_CHANNEL_COUNT      = 5
)

var (
	EXCHANGE_PUBLISH_COMMENT = event.Exchange_EXCHANGE_PUBLISH_COMMENT.String()
	EXCHANGE_DELETE_COMMENT  = event.Exchange_EXCHANGE_DELETE_COMMENT.String()
)

var (
	KEY_PUBLISH_COMMENT = event.RoutingKey_KEY_PUBLISH_COMMENT.String()
	KEY_DELETE_COMMENT  = event.RoutingKey_KEY_DELETE_COMMENT.String()
)
