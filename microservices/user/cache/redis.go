package cache

import (
	// "context"
	// "fmt"
	"context"
	"fmt"
	"log"
	"time"

	// "github.com/go-redis/redis/v8"

	config "github.com/Yux77Yux/platform_backend/microservices/user/config"
	pkgRedis "github.com/Yux77Yux/platform_backend/pkg/redis_cache"
)

var (
	RedisClient *pkgRedis.RedisClient
)

func init() {
	var err error
	redisAddr := config.REDIS_STR
	redisPassword := config.REDIS_PASSWORD

	if RedisClient, err = pkgRedis.OpenRedis(redisAddr, redisPassword); err != nil {
		log.Printf("error: %v", err)
	}
}

func ExistsUsername(username string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	// 用一个 goroutine 来执行查找
	resultCh := make(chan struct {
		pos int64
		err error
	})

	go func() {
		pos, err := RedisClient.FindElementList(ctx, "List", "Username", username)
		resultCh <- struct {
			pos int64
			err error
		}{pos, err}
	}()

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return -1, fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result.err != nil {
			return -1, result.err
		}
		return result.pos, nil
	}
}

func StoreUsername(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	resultCh := make(chan error)

	go func() {
		err := RedisClient.LPushList(ctx, "List", "Username", username)
		resultCh <- err
	}()

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result != nil {
			return result
		}
		return nil
	}
}
