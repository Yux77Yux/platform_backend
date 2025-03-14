package main

import (
	"context"
	"os"
	"sync"

	service "github.com/Yux77Yux/platform_backend/microservices/auth/service"
	tools "github.com/Yux77Yux/platform_backend/microservices/auth/tools"
)

func Run(ctx context.Context) {
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		defer wg.Done()
		service.ServerRun(ctx)
	}()

	wg.Wait()
	tools.LogInfo(tools.GetMainValue(ctx), "main exit")
	os.Exit(0)
}
