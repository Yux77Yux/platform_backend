package receiver

var (
	dispatcher DispatchInterface
	db         SqlMethod
	messaging  MessageQueueMethod
	cache      CacheMethod
)

func InitDb(_db SqlMethod) {
	db = _db
}

func InitCache(_cache CacheMethod) {
	cache = _cache
}

// 非RPC类型的消息队列的交换机声明
func Run(_messaging MessageQueueMethod, _dispatch DispatchInterface) {
	dispatcher = _dispatch
	messaging = _messaging
	for exchange := range ExchangesConfig {
		switch exchange {
		// 不同的exchange使用不同函数
		case EXCHANGE_REGISTER:
			go messaging.ListenToQueue(exchange, QUEUE_REGISTER, KEY_REGISTER, registerProcessor)
		case EXCHANGE_STORE_USER:
			go messaging.ListenToQueue(exchange, QUEUE_STORE_USER, KEY_STORE_USER, storeUserProcessor)
		case EXCHANGE_STORE_CREDENTIAL:
			go messaging.ListenToQueue(exchange, QUEUE_STORE_CREDENTIAL, KEY_STORE_CREDENTIAL, storeCredentialsProcessor)
		case EXCHANGE_UPDATE_USER_SPACE:
			go messaging.ListenToQueue(exchange, QUEUE_UPDATE_USER_SPACE, KEY_UPDATE_USER_SPACE, updateUserSpaceProcessor)
		case EXCHANGE_UPDATE_USER_BIO:
			go messaging.ListenToQueue(exchange, QUEUE_UPDATE_USER_BIO, KEY_UPDATE_USER_BIO, updateUserBioProcessor)
		case EXCHANGE_UPDATE_USER_AVATAR:
			go messaging.ListenToQueue(exchange, QUEUE_UPDATE_USER_AVATAR, KEY_UPDATE_USER_AVATAR, updateUserAvatarProcessor)
		case EXCHANGE_UPDATE_USER_STATUS:
			go messaging.ListenToQueue(exchange, QUEUE_UPDATE_USER_STATUS, KEY_UPDATE_USER_STATUS, updateUserStatusProcessor)
		case EXCHANGE_FOLLOW:
			go messaging.ListenToQueue(exchange, QUEUE_FOLLOW, KEY_FOLLOW, followProcessor)
		// case EXCHANGE_CANCEL_FOLLOW:
		// 	go messaging.ListenToQueue(exchange, QUEUE_CANCEL_FOLLOW, KEY_CANCEL_FOLLOW, followProcessor)
		case EXCHANGE_DEL_REVIEWER:
			go messaging.ListenToQueue(exchange, QUEUE_DEL_REVIEWER, KEY_DEL_REVIEWER, delReviewerProcessor)
		}
	}
}
