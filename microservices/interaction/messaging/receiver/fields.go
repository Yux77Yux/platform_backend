package receiver

import (
	event "github.com/Yux77Yux/platform_backend/generated/event/interaction"
)

var (
	ExchangesConfig = map[string]string{
		EXCHANGE_ADD_COLLECTION:   "direct",
		EXCHANGE_ADD_LIKE:         "direct",
		EXCHANGE_ADD_VIEW:         "direct",
		EXCHANGE_BATCH_UPDATE_DB:  "direct",
		EXCHANGE_CANCEL_LIKE:      "direct",
		EXCHANGE_COMPUTE_CREATION: "direct",
		EXCHANGE_COMPUTE_USER:     "direct",
		EXCHANGE_UPDATE_DB:        "direct",
		// Add more exchanges here
	}
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

var (
	QUEUE_ADD_COLLECTION   = event.Queue_QUEUE_ADD_COLLECTION.String()
	QUEUE_ADD_LIKE         = event.Queue_QUEUE_ADD_LIKE.String()
	QUEUE_ADD_VIEW         = event.Queue_QUEUE_ADD_VIEW.String()
	QUEUE_BATCH_UPDATE_DB  = event.Queue_QUEUE_BATCH_UPDATE_DB.String()
	QUEUE_CANCEL_LIKE      = event.Queue_QUEUE_CANCEL_LIKE.String()
	QUEUE_COMPUTE_CREATION = event.Queue_QUEUE_COMPUTE_CREATION.String()
	QUEUE_COMPUTE_USER     = event.Queue_QUEUE_COMPUTE_USER.String()
	QUEUE_UPDATE_DB        = event.Queue_QUEUE_UPDATE_DB.String()
)
