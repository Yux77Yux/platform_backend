package messaging

import (
	event "github.com/Yux77Yux/platform_backend/generated/common/event"
)

const (
	UPDATE_CREATION_ACTION_COUNT = "InteractionCount"

	ADD_VIEW = "ADD_VIEW"
)

var (
	connStr         string
	ExchangesConfig = map[string]string{
		event.Exchange_EXCHANGE_ADD_VIEW.String(): "direct",
	}
)

func InitStr(_str string) {
	connStr = _str
}
