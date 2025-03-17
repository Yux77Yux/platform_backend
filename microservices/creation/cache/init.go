package cache

import (
	"log"

	internal "github.com/Yux77Yux/platform_backend/microservices/creation/internal"
	receiver "github.com/Yux77Yux/platform_backend/microservices/creation/messaging/receiver"
	tools "github.com/Yux77Yux/platform_backend/microservices/creation/tools"
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
	client := GetCacheClient()
	cache := &CacheMethodStruct{
		CacheClient: client,
	}

	internal.InitCache(cache)
	receiver.InitCache(cache)

	return func() {
		err := client.Close()
		if err != nil {
			tools.LogError("", "Cache Close", err)
		}
	}
}
