package main

import (
	"context"
	"os"
	"sync"

	cache "github.com/Yux77Yux/platform_backend/microservices/aggregator/cache"
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
	messaging "github.com/Yux77Yux/platform_backend/microservices/aggregator/messaging"
	service "github.com/Yux77Yux/platform_backend/microservices/aggregator/service"
	tools "github.com/Yux77Yux/platform_backend/microservices/aggregator/tools"
)

func Run(signal chan os.Signal) {
	var (
		closeServer         func(chan any)
		closeMessagingQueue func()
		closeCache          func()
		closeClient         func()
		wg                  sync.WaitGroup
	)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		closeCache = cache.Run()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		closeClient = client.Run()
	}(&wg)
	// 先启动缓存和客户端
	wg.Wait()

	traceID := tools.GetUuid()
	_ctx := context.WithValue(context.Background(), "main", traceID)
	closeMessagingQueue = messaging.Run(_ctx)

	// 最后启动服务器
	closeServer = service.ServerRun()

	<-signal
	s_closed_sig := make(chan any, 1)
	closeServer(s_closed_sig)
	<-s_closed_sig

	closeMessagingQueue()

	closeClient()
	closeCache()

	tools.LogInfo(traceID.String(), "main exit")
	os.Exit(0)
}
