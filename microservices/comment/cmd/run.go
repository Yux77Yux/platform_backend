package main

import (
	"context"
	"os"
	"sync"

	cache "github.com/Yux77Yux/platform_backend/microservices/comment/cache"
	receiver "github.com/Yux77Yux/platform_backend/microservices/comment/messaging/receiver"
	service "github.com/Yux77Yux/platform_backend/microservices/comment/service"
	tools "github.com/Yux77Yux/platform_backend/microservices/comment/tools"
)

func Run(ctx context.Context) {
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		defer wg.Done()
		cache.Run(ctx)
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		receiver.Run(ctx)
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		service.ServerRun(ctx)
	}()

	wg.Wait()
	tools.LogInfo(tools.GetMainValue(ctx), "main exit")
	os.Exit(0)
}
