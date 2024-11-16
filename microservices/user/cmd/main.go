package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Yux77Yux/platform_backend/microservices/user/service"
)

func main() {
	// 创建一个信道来接收终止信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM) // 捕获中断和终止信号

	// 创建取消上下文
	ctx, cancel := context.WithCancel(context.Background())

	// 创建一个信道用于通知优雅关闭完成
	done := make(chan struct{})

	// 启动服务
	go service.ServerRun(ctx, done)

	// 等待信号
	sig := <-signalChan
	log.Printf("info: received signal: %s. Initiating graceful shutdown...", sig)

	// 取消上下文，通知服务停止
	cancel()

	// 等待关闭完成或超时
	select {
	case <-done:
		log.Println("info: program stopped gracefully.")
		os.Exit(0)
	case <-time.After(100 * time.Second):
		log.Println("warning: timeout reached. Forcing shutdown.")
		os.Exit(1)
	}
}
