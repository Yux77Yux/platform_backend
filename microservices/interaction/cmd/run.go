package main

import (
	"context"
	"os"
	"sync"

	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	receiver "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging/receiver"
	db "github.com/Yux77Yux/platform_backend/microservices/interaction/repository"
	service "github.com/Yux77Yux/platform_backend/microservices/interaction/service"
	tools "github.com/Yux77Yux/platform_backend/microservices/interaction/tools"
)

func Run(ctx context.Context) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := cache.Run(ctx)
		if err != nil {
			tools.LogSuperError(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := db.Run(ctx)
		if err != nil {
			tools.LogSuperError(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		receiver.Run(ctx)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		service.ServerRun(ctx)
	}()

	wg.Wait()
	tools.LogInfo(tools.GetMainValue(ctx), "main exit")
	os.Exit(0)
}
