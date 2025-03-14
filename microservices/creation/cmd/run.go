package main

import (
	"context"
	"os"
	"sync"

	cache "github.com/Yux77Yux/platform_backend/microservices/creation/cache"
	receiver "github.com/Yux77Yux/platform_backend/microservices/creation/messaging/receiver"
	oss "github.com/Yux77Yux/platform_backend/microservices/creation/oss"
	db "github.com/Yux77Yux/platform_backend/microservices/creation/repository"
	service "github.com/Yux77Yux/platform_backend/microservices/creation/service"
	tools "github.com/Yux77Yux/platform_backend/microservices/creation/tools"
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
