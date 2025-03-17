package internal

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
