package internal

import (
	event "github.com/Yux77Yux/platform_backend/generated/event/review"
)

var (
	EXCHANGE_COMMENT_REVIEW         = event.Exchange_EXCHANGE_COMMENT_REVIEW.String()
	EXCHANGE_USER_REVIEW            = event.Exchange_EXCHANGE_USER_REVIEW.String()
	EXCHANGE_CREATION_REVIEW        = event.Exchange_EXCHANGE_CREATION_REVIEW.String()
	EXCHANGE_NEW_REVIEW             = event.Exchange_EXCHANGE_NEW_REVIEW.String()
	EXCHANGE_UPDATE                 = event.Exchange_EXCHANGE_UPDATE.String()
	EXCHANGE_BATCH_UPDATE           = event.Exchange_EXCHANGE_BATCH_UPDATE.String()
	EXCHANGE_PEND_CREATION          = event.Exchange_EXCHANGE_PEND_CREATION.String()
	EXCHANGE_UPDATE_USER_STATUS     = event.Exchange_EXCHANGE_UPDATE_USER_STATUS.String()
	EXCHANGE_UPDATE_CREATION_STATUS = event.Exchange_EXCHANGE_UPDATE_CREATION_STATUS.String()
	EXCHANGE_DELETE_CREATION        = event.Exchange_EXCHANGE_DELETE_CREATION.String()
	EXCHANGE_DELETE_COMMENT         = event.Exchange_EXCHANGE_DELETE_COMMENT.String()
)

var (
	KEY_COMMENT_REVIEW         = event.RoutingKey_KEY_COMMENT_REVIEW.String()
	KEY_USER_REVIEW            = event.RoutingKey_KEY_USER_REVIEW.String()
	KEY_CREATION_REVIEW        = event.RoutingKey_KEY_CREATION_REVIEW.String()
	KEY_NEW_REVIEW             = event.RoutingKey_KEY_NEW_REVIEW.String()
	KEY_UPDATE                 = event.RoutingKey_KEY_UPDATE.String()
	KEY_BATCH_UPDATE           = event.RoutingKey_KEY_BATCH_UPDATE.String()
	KEY_PEND_CREATION          = event.RoutingKey_KEY_PEND_CREATION.String()
	KEY_UPDATE_USER_STATUS     = event.RoutingKey_KEY_UPDATE_USER_STATUS.String()
	KEY_UPDATE_CREATION_STATUS = event.RoutingKey_KEY_UPDATE_CREATION_STATUS.String()
	KEY_DELETE_CREATION        = event.RoutingKey_KEY_DELETE_CREATION.String()
	KEY_DELETE_COMMENT         = event.RoutingKey_KEY_DELETE_COMMENT.String()
)
