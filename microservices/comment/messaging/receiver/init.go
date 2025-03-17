package receiver

var (
	ExchangesConfig = map[string]string{
		"PublishComment": "direct",
		"DeleteComment":  "direct",
	}
)

var (
	// dispatch
	DispatchInsert string
	DispatchDelete string
)

var (
	db         SqlMethod
	dispatcher DispatchInterface
	messaging  MessageQueueMethod
)

func InitDb(_db SqlMethod) {
	db = _db
}

func Run(mq MessageQueueMethod, _dispatch DispatchInterface) {
	messaging = mq
	dispatcher = _dispatch

	for exchange := range ExchangesConfig {
		switch exchange {
		// 不同的exchange使用不同函数
		case EXCHANGE_DELETE_COMMENT:
			go messaging.ListenToQueue(exchange, QUEUE_DELETE_COMMENT, KEY_DELETE_COMMENT, JoinCommentProcessor)
		case EXCHANGE_PUBLISH_COMMENT:
			go messaging.ListenToQueue(exchange, QUEUE_DELETE_COMMENT, KEY_DELETE_COMMENT, DeleteCommentProcessor)
		}
	}
}
