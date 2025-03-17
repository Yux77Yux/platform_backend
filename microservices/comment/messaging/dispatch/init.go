package dispatch

var (
	db    SqlMethod
	cache CacheMethod
)

func InitDb(_db SqlMethod) {
	db = _db
}

func InitCache(_cache CacheMethod) {
	cache = _cache
}
