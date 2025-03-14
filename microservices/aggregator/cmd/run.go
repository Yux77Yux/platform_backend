package main

import (
	"context"
	"os"
	"sync"

	cache "github.com/Yux77Yux/platform_backend/microservices/aggregator/cache"
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
	receiver "github.com/Yux77Yux/platform_backend/microservices/aggregator/messaging/receiver"
	service "github.com/Yux77Yux/platform_backend/microservices/aggregator/service"
	tools "github.com/Yux77Yux/platform_backend/microservices/aggregator/tools"
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
		client.Run(ctx)
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
