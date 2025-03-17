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
		case EXCHANGE_COMPUTE_CREATION:
			go messaging.ListenToQueue(exchange, QUEUE_COMPUTE_CREATION, KEY_COMPUTE_CREATION, computeSimilarProcessor)
		case EXCHANGE_COMPUTE_USER:
			go messaging.ListenToQueue(exchange, QUEUE_COMPUTE_USER, KEY_COMPUTE_USER, computeUserProcessor)
		case EXCHANGE_UPDATE_DB:
			go messaging.ListenToQueue(exchange, QUEUE_UPDATE_DB, KEY_UPDATE_DB, updateDbInteraction)
		case EXCHANGE_ADD_VIEW:
			go messaging.ListenToQueue(exchange, QUEUE_ADD_VIEW, KEY_ADD_VIEW, addViewProcessor)
		case EXCHANGE_ADD_COLLECTION:
			go messaging.ListenToQueue(exchange, QUEUE_ADD_COLLECTION, KEY_ADD_COLLECTION, addCollectionProcessor)
		case EXCHANGE_ADD_LIKE:
			go messaging.ListenToQueue(exchange, QUEUE_ADD_COLLECTION, KEY_ADD_LIKE, addLikeProcessor)
		case EXCHANGE_CANCEL_LIKE:
			go messaging.ListenToQueue(exchange, QUEUE_CANCEL_LIKE, KEY_CANCEL_LIKE, cancelLikeProcessor)
		case EXCHANGE_BATCH_UPDATE_DB:
			go messaging.ListenToQueue(exchange, QUEUE_BATCH_UPDATE_DB, KEY_BATCH_UPDATE_DB, batchUpdateDbProcessor)
		}
	}
}
