package cache

import (
	"context"
	"fmt"
	"log"
	"time"
)

func ExistsUsername(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resultCh := make(chan struct {
		exist bool
		err   error
	}, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		exist, err := CacheClient.ExistsInSet(ctx, "User", "Username", username)

		select {
		case resultCh <- struct {
			exist bool
			err   error
		}{exist, err}:
			log.Printf("info: completely execute for cache method: ExistsUsername")
		case <-ctx.Done():
			log.Printf("warning: context canceled for cache method: ExistsUsername")
		}
	}

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return false, fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result.err != nil {
			return false, result.err
		}
		return result.exist, nil
	}
}

func StoreUsername(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	resultCh := make(chan error, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		err := CacheClient.AddToSet(ctx, "User", "Username", username)
		resultCh <- err
	}

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
