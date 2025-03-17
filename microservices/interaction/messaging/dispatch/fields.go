package dispatch

import (
	event "github.com/Yux77Yux/platform_backend/generated/event/interaction"
)

const (
	LISTENER_CHANNEL_COUNT = 80
	MAX_BATCH_SIZE         = 50
	EXE_CHANNEL_COUNT      = 5

	// sql的插入更新删除都用一套
	DbInteraction      = "DbInteraction"
	DbBatchInteraction = "DbBatchInteraction"

	ViewCache       = "ViewCache"
	LikeCache       = "LikeCache"
	CollectionCache = "CollectionCache"
	CancelLikeCache = "CancelLikeCache"
)

var (
	EXCHANGE_ADD_COLLECTION               = event.Exchange_EXCHANGE_ADD_COLLECTION.String()
	EXCHANGE_ADD_LIKE                     = event.Exchange_EXCHANGE_ADD_LIKE.String()
	EXCHANGE_ADD_VIEW                     = event.Exchange_EXCHANGE_ADD_VIEW.String()
	EXCHANGE_BATCH_UPDATE_DB              = event.Exchange_EXCHANGE_BATCH_UPDATE_DB.String()
	EXCHANGE_CANCEL_LIKE                  = event.Exchange_EXCHANGE_CANCEL_LIKE.String()
	EXCHANGE_COMPUTE_CREATION             = event.Exchange_EXCHANGE_COMPUTE_CREATION.String()
	EXCHANGE_COMPUTE_USER                 = event.Exchange_EXCHANGE_COMPUTE_USER.String()
	EXCHANGE_UPDATE_CREATION_ACTION_COUNT = event.Exchange_EXCHANGE_UPDATE_CREATION_ACTION_COUNT.String()
	EXCHANGE_UPDATE_DB                    = event.Exchange_EXCHANGE_UPDATE_DB.String()
)

var (
	KEY_ADD_COLLECTION               = event.RoutingKey_KEY_ADD_COLLECTION.String()
	KEY_ADD_LIKE                     = event.RoutingKey_KEY_ADD_LIKE.String()
	KEY_ADD_VIEW                     = event.RoutingKey_KEY_ADD_VIEW.String()
	KEY_BATCH_UPDATE_DB              = event.RoutingKey_KEY_BATCH_UPDATE_DB.String()
	KEY_CANCEL_LIKE                  = event.RoutingKey_KEY_CANCEL_LIKE.String()
	KEY_COMPUTE_CREATION             = event.RoutingKey_KEY_COMPUTE_CREATION.String()
	KEY_COMPUTE_USER                 = event.RoutingKey_KEY_COMPUTE_USER.String()
	KEY_UPDATE_CREATION_ACTION_COUNT = event.RoutingKey_KEY_UPDATE_CREATION_ACTION_COUNT.String()
	KEY_UPDATE_DB                    = event.RoutingKey_KEY_UPDATE_DB.String()
)
