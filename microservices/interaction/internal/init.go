package internal

var (
	cache     CacheMethod
	db        DataBaseMethod
	messaging MessageQueueMethod
)

func InitCache(_cache CacheMethod) {
	cache = _cache
}

func InitDb(_db DataBaseMethod) {
	db = _db
}

func InitMQ(messaging MessageQueueMethod) {
	messaging = _messaging
}
