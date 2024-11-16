package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	generatedAuth "github.com/Yux77Yux/platform_backend/generated/auth"
	"google.golang.org/grpc"
)

type Server struct {
	generatedAuth.UnimplementedAuthServiceServer
}

func ServerRun() {
	var (
		grpcServer *grpc.Server
		lis        net.Listener
		err        error
	)

	lis, err = net.Listen("tcp", ":8080")
	if err != nil {
		wrappedErr := fmt.Errorf("failed to listen: %w", err)
		log.Fatalf("%v", wrappedErr)
	}

	go func() {
		grpcServer = grpc.NewServer()
		generatedAuth.RegisterAuthServiceServer(grpcServer, &Server{}) // 注册 Auth 服务

		log.Println("info: erver is running on port :50010")
		if err = grpcServer.Serve(lis); err != nil {
			wrappedErr := fmt.Errorf("failed to serve: %w", err)
			log.Fatalf("%v", wrappedErr)
		}
	}()

	// 创建一个信道来接收终止信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM) // 捕获中断和终止信号

	// 等待信号
	sig := <-signalChan
	log.Printf("info: received signal: %s. Shutting down gracefully...", sig)

	// 设置 8 秒的上下文超时，用于优雅地停止 gRPC 服务器
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	// 优雅关闭 gRPC 服务器
	done := make(chan struct{})
	go func() {
		grpcServer.GracefulStop()
		close(done)
	}()

	// 等待关闭完成或超时
	select {
	case <-done:
		log.Println("info: server stopped gracefully.")
	case <-ctx.Done():
		log.Println("warning: timeout reached. Forcing shutdown.")
		grpcServer.Stop() // 强制关闭服务器
	}

	log.Println("info: server exited.")
}
