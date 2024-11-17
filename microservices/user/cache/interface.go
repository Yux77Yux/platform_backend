package cache

import (
	"context"
	"log"

	pkgCache "github.com/Yux77Yux/platform_backend/pkg/cache"
)

type CacheInterface interface {
	Open(connStr string, password string) error
	Close()

	LPushList(ctx context.Context, direction string, kind string, unique string, value ...interface{}) error
	FindElementList(ctx context.Context, kind string, unique string, value string) (int64, error)
}

var (
	addr     string
	password string
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
