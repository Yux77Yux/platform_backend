package dispatch

var (
	messaging MessageQueueMethod
	cache     CacheMethod
	db        SqlMethod
)

func InitMQ(_messaging MessageQueueMethod) {
	messaging = _messaging
}

func InitCache(_cache CacheMethod) {
	cache = _cache
}

func InitDb(_db SqlMethod) {
	db = _db
}
