package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/Yux77Yux/platform_backend/microservices/interaction/config" // 保证配置初始化
	tools "github.com/Yux77Yux/platform_backend/microservices/interaction/tools"
)

func main() {
	// 接收终止信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM) // 捕获中断和终止信号

	// 运行
	go Run(signalChan)

	<-signalChan

	select {
	case <-time.After(3 * time.Minute):
		tools.LogWarning("main", "exit", "timeout reached. Forcing shutdown")
		os.Exit(1)
	}
}
