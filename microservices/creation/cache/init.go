package cache

import (
	"context"
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
	if CacheClient != nil {
		return CacheClient
	}
	CacheClient = &pkgCache.RedisClient{}
	err := CacheClient.Open(addr, password)
	if err != nil {
		log.Printf("error: failed to connect the cache client: %v", err)
		return nil
	}

	return CacheClient
}

func Run(ctx context.Context) error {
	CacheClient = GetCacheClient()
	<-ctx.Done()
	return CacheClient.Close()
}
