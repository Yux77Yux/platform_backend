package cache

import (
	"log"

	receiver "github.com/Yux77Yux/platform_backend/microservices/aggregator/messaging/receiver"
	tools "github.com/Yux77Yux/platform_backend/microservices/aggregator/tools"
	pkgCache "github.com/Yux77Yux/platform_backend/pkg/cache"
)

var (
	addr     string
	password string
)

func InitStr(Addr, Password string) {
	addr, password = Addr, Password
}

func GetCacheClient() CacheInterface {
	CacheClient := &pkgCache.RedisClient{}
	err := CacheClient.Open(addr, password)
	if err != nil {
		log.Printf("error: failed to connect the cache client: %v", err)
		return nil
	}

	return CacheClient
}

func Run() func() {
	cache := &CacheMethodStruct{
		CacheClient: GetCacheClient(),
	}
	receiver.InitCache(cache)
	return func() {
		err := cache.CacheClient.Close()
		if err != nil {
			tools.LogError("", "Cache Close", err)
		}
	}
}
