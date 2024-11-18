package cache

import (
	"log"

	pkgCache "github.com/Yux77Yux/platform_backend/pkg/cache"
)

var (
	addr                string
	password            string
	cacheRequestChannel chan RequestHandlerFunc
)

func InitStr(Addr, Password string) {
	addr, password = Addr, Password
}

func GetCacheClient() CacheInterface {
	var cache CacheInterface = &pkgCache.RedisClient{}
	err := cache.Open(addr, password)
	if err != nil {
		log.Printf("error: failed to connect the cache client: %v", err)
		return nil
	}

	return cache
}

func InitWorker(master *RequestProcessor) {
	cacheRequestChannel = master.GetChannel()
}
