package receiver

var (
	dispatcher    DispatchInterface
	db            SqlMethod
	messaging     MessageQueueMethod
	cache         CacheMethod
	search_client SearchServiceInterface
)

func InitSearch(_client SearchServiceInterface) {
	search_client = _client
}

func InitDb(_db SqlMethod) {
	db = _db
}

func InitCache(_cache CacheMethod) {
	cache = _cache
}

// 非RPC类型的消息队列的交换机声明
func Run(_messaging MessageQueueMethod, _dispatch DispatchInterface) {
	messaging = _messaging
	dispatcher = _dispatch
	for exchange := range ExchangesConfig {
		switch exchange {
		// 不同的exchange使用不同函数
		case EXCHANGE_UPDATE_DB_CREATION:
			go messaging.ListenToQueue(exchange, QUEUE_UPDATE_DB_CREATION, KEY_UPDATE_DB_CREATION, updateCreationDbProcessor)
		case EXCHANGE_STORE_CREATION:
			go messaging.ListenToQueue(exchange, QUEUE_STORE_CREATION, KEY_STORE_CREATION, storeCreationProcessor)
		case EXCHANGE_UPDATE_CREATION_STATUS:
			go messaging.ListenToQueue(exchange, QUEUE_UPDATE_CREATION_STATUS, KEY_UPDATE_CREATION_STATUS, updateCreationStatusProcessor)
		case EXCHANGE_UPDATE_CACHE_CREATION:
			go messaging.ListenToQueue(exchange, QUEUE_UPDATE_CACHE_CREATION, KEY_UPDATE_CACHE_CREATION, updateCreationCacheProcessor)
		case EXCHANGE_DELETE_CREATION:
			go messaging.ListenToQueue(exchange, QUEUE_DELETE_CREATION, KEY_DELETE_CREATION, deleteCreationProcessor)
		case EXCHANGE_UPDATE_CREATION_ACTION_COUNT:
			go messaging.ListenToQueue(exchange, QUEUE_UPDATE_CREATION_ACTION_COUNT, KEY_UPDATE_CREATION_ACTION_COUNT, addInteractionCount)
		}
	}
}
