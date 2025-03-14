package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/Yux77Yux/platform_backend/microservices/aggregator/config" // 保证配置初始化
	tools "github.com/Yux77Yux/platform_backend/microservices/aggregator/tools"
)

func main() {
	traceID := tools.GetUuid()
	_parent := context.WithValue(context.Background(), "main", traceID)
	_ctx, _cancel := context.WithCancel(_parent)

	// 接收终止信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM) // 捕获中断和终止信号

	// 运行
	go Run(_ctx)

	<-signalChan
	_cancel()

	select {
	case <-time.After(3 * time.Minute):
		tools.LogWarning(traceID.String(), "main exit", "timeout reached. Forcing shutdown")
		os.Exit(1)
	}
}
