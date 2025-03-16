package internal

import (
	event "github.com/Yux77Yux/platform_backend/generated/event/user"
)

var (
	EXCHANGE_REGISTER           = event.Exchange_EXCHANGE_REGISTER.String()
	EXCHANGE_STORE_USER         = event.Exchange_EXCHANGE_STORE_USER.String()
	EXCHANGE_STORE_CREDENTIAL   = event.Exchange_EXCHANGE_STORE_CREDENTIAL.String()
	EXCHANGE_UPDATE_USER_SPACE  = event.Exchange_EXCHANGE_UPDATE_USER_SPACE.String()
	EXCHANGE_UPDATE_USER_BIO    = event.Exchange_EXCHANGE_UPDATE_USER_BIO.String()
	EXCHANGE_UPDATE_USER_AVATAR = event.Exchange_EXCHANGE_UPDATE_USER_AVATAR.String()
	EXCHANGE_FOLLOW             = event.Exchange_EXCHANGE_FOLLOW.String()
	EXCHANGE_CANCEL_FOLLOW      = event.Exchange_EXCHANGE_CANCEL_FOLLOW.String()
	EXCHANGE_UPDATE_USER_STATUS = event.Exchange_EXCHANGE_UPDATE_USER_STATUS.String()
	EXCHANGE_DEL_REVIEWER       = event.Exchange_EXCHANGE_DEL_REVIEWER.String()
)

var (
	KEY_REGISTER           = event.RoutingKey_KEY_REGISTER.String()
	KEY_STORE_USER         = event.RoutingKey_KEY_STORE_USER.String()
	KEY_STORE_CREDENTIAL   = event.RoutingKey_KEY_STORE_CREDENTIAL.String()
	KEY_UPDATE_USER_SPACE  = event.RoutingKey_KEY_UPDATE_USER_SPACE.String()
	KEY_UPDATE_USER_BIO    = event.RoutingKey_KEY_UPDATE_USER_BIO.String()
	KEY_UPDATE_USER_AVATAR = event.RoutingKey_KEY_UPDATE_USER_AVATAR.String()
	KEY_FOLLOW             = event.RoutingKey_KEY_FOLLOW.String()
	KEY_CANCEL_FOLLOW      = event.RoutingKey_KEY_CANCEL_FOLLOW.String()
	KEY_UPDATE_USER_STATUS = event.RoutingKey_KEY_UPDATE_USER_STATUS.String()
	KEY_DEL_REVIEWER       = event.RoutingKey_KEY_DEL_REVIEWER.String()
)

var (
	QUEUE_REGISTER           = event.Queue_QUEUE_REGISTER.String()
	QUEUE_STORE_USER         = event.Queue_QUEUE_STORE_USER.String()
	QUEUE_STORE_CREDENTIAL   = event.Queue_QUEUE_STORE_CREDENTIAL.String()
	QUEUE_UPDATE_USER_SPACE  = event.Queue_QUEUE_UPDATE_USER_SPACE.String()
	QUEUE_UPDATE_USER_BIO    = event.Queue_QUEUE_UPDATE_USER_BIO.String()
	QUEUE_UPDATE_USER_AVATAR = event.Queue_QUEUE_UPDATE_USER_AVATAR.String()
	QUEUE_FOLLOW             = event.Queue_QUEUE_FOLLOW.String()
	QUEUE_CANCEL_FOLLOW      = event.Queue_QUEUE_CANCEL_FOLLOW.String()
	QUEUE_UPDATE_USER_STATUS = event.Queue_QUEUE_UPDATE_USER_STATUS.String()
	QUEUE_DEL_REVIEWER       = event.Queue_QUEUE_DEL_REVIEWER.String()
)
