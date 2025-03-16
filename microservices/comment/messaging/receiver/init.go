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

func Run(mq MessageQueueMethod, _dispatch DispatchInterface) {
	messaging := mq
	dispatch = _dispatch

	for exchange := range ExchangesConfig {
		switch exchange {
		// 不同的exchange使用不同函数
		case PublishComment:
			go messaging.ListenToQueue(exchange, PublishComment, PublishComment, JoinCommentProcessor)
		case DeleteComment:
			go messaging.ListenToQueue(exchange, DeleteComment, DeleteComment, DeleteCommentProcessor)
		}
	}
}
