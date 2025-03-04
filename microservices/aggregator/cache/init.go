package cache

import (
	"log"

	pkgCache "github.com/Yux77Yux/platform_backend/pkg/cache"
)

var (
	addr        string
	password    string
	CacheClient CacheInterface
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

func CloseClient() {
	if err := CacheClient.Close(); err != nil {
		log.Println("error: cache client close error.")
		return
	}
}

func Init() {
	CacheClient = GetCacheClient()
}
