package internal

var (
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

func InitMQ(_messaging MessageQueueMethod) {
	messaging = _messaging
}

func InitCache(_cache CacheMethod) {
	cache = _cache
}
