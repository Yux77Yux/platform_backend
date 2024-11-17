package cache

import (
	// "context"
	// "fmt"
	"context"
	"fmt"
	"log"
	"time"
)

type checkTask struct {
	ctx      context.Context
	username string

	resultCh chan struct {
		exist bool
		err   error
	}
}

var (
	checkUsernameChan chan checkTask
)

// 初始化依赖，供cmd集中调用管理
func InitWorker() func() {
	headChan := &ChanChain{}

	checkUsernameChan = make(chan checkTask, 20)

	next := &ChanChain{
		closeChanClosure: func() error {
			close(checkUsernameChan)
			return nil // 返回 nil 或者错误
		},
	}

	headChan.register(next)

	go checkWorker()

	return func() {
		headChan.closeChan()
	}
}

func checkWorker() {
	CacheClient := GetCacheClient()
	if CacheClient == nil {
		close(checkUsernameChan) // 显式关闭通道以避免后续阻塞
		return
	}

	for request := range checkUsernameChan {
		pos, err := CacheClient.FindElementList(request.ctx, "List", "Username", request.username)
		exist := pos > -1

		select {
		case request.resultCh <- struct {
			exist bool
			err   error
		}{exist, err}:
			// 正常写入
		case <-request.ctx.Done():
			log.Printf("warning: context canceled for username: %s", request.username)
		}
	}
}

func ExistsUsername(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	resultCh := make(chan struct {
		exist bool
		err   error
	}, 1)
	checkUsernameChan <- checkTask{
		ctx:      ctx,
		username: username,
		resultCh: resultCh,
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	resultCh := make(chan error)

	go func() {
		CacheClient := GetCacheClient()
		if CacheClient == nil {
			log.Printf("error: failed to connect the cache client")
			return
		}

		err := CacheClient.LPushList(ctx, "List", "Username", username)
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
