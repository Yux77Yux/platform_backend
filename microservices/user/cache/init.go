package cache

import (
	"log"

	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/user/messaging/dispatch"
	receiver "github.com/Yux77Yux/platform_backend/microservices/user/messaging/receiver"
	tools "github.com/Yux77Yux/platform_backend/microservices/user/tools"
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
	internal.InitCache(cache)
	receiver.InitCache(cache)
	dispatch.InitCache(cache)
	return func() {
		err := cache.CacheClient.Close()
		if err != nil {
			tools.LogError("", "Cache Close", err)
		}
	}
}
