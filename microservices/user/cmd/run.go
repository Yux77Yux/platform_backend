package main

import (
	"context"
	"os"
	"sync"

	cache "github.com/Yux77Yux/platform_backend/microservices/user/cache"
	receiver "github.com/Yux77Yux/platform_backend/microservices/user/messaging/receiver"
	oss "github.com/Yux77Yux/platform_backend/microservices/user/oss"
	db "github.com/Yux77Yux/platform_backend/microservices/user/repository"
	service "github.com/Yux77Yux/platform_backend/microservices/user/service"
	tools "github.com/Yux77Yux/platform_backend/microservices/user/tools"
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
		oss.Run(ctx)
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
