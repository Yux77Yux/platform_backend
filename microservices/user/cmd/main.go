package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
	userMQ "github.com/Yux77Yux/platform_backend/microservices/user/messaging"
	service "github.com/Yux77Yux/platform_backend/microservices/user/service"
)

func main() {
	done := make(chan struct{})
	// 初始化服务器
	go service.ServerRun(done)
	// 初始化dispatcher
	master := userMQ.InitDispatch()
	master.Start()
	internal.EmpowerDispatch(master)
	defer master.Shutdown()

	// 创建一个信道来接收终止信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM) // 捕获中断和终止信号

	// 等待信号
	sig := <-signalChan
	log.Printf("info: received signal: %s. Initiating graceful shutdown...", sig)

	// 创建取消上下文
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	// 取消上下文，通知服务停止
	defer cancel()

	// 等待关闭完成或超时
	select {
	case <-done:
		log.Println("info: graceful shutdown...")
		os.Exit(0)
	case <-ctx.Done():
		log.Println("warning: timeout reached. Forcing shutdown.")
		os.Exit(1)
	}
}
