package main

import (
	"context"
	"os"
	"sync"

	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	httpServer "github.com/Yux77Yux/platform_backend/microservices/interaction/cmd/http_server"
	messaging "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/interaction/repository"
	service "github.com/Yux77Yux/platform_backend/microservices/interaction/service"
	tools "github.com/Yux77Yux/platform_backend/microservices/interaction/tools"
)

func Run(signal chan os.Signal) {
	var (
		closeHttpServer     func(chan any)
		closeServer         func(chan any)
		closeMessagingQueue func()
		closeCache          func()
		closeDataBase       func()
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
		closeDataBase = db.Run()
	}(&wg)
	// 先启动缓存和客户端
	wg.Wait()

	traceID := tools.GetUuid()
	_ctx := context.WithValue(context.Background(), "main", traceID)
	closeMessagingQueue = messaging.Run(_ctx)

	// 最后启动服务器
	closeHttpServer = httpServer.Run()
	closeServer = service.ServerRun()

	<-signal

	s_closed_sig := make(chan any)
	h_closed_sig := make(chan any)
	// 先关闭服务器，防止新的请求，依次关闭消息队列
	closeHttpServer(h_closed_sig)
	<-h_closed_sig

	closeServer(s_closed_sig)
	<-s_closed_sig

	// 关闭消息队列，等待请求停止
	closeMessagingQueue()

	closeDataBase()
	closeCache()

	tools.LogInfo(traceID.String(), "main exit")
	os.Exit(0)
}
