package service

import (
	"fmt"
	"log"
	"net"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/auth"
	middlewares "github.com/Yux77Yux/platform_backend/pkg/middlewares"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ServerRun(done chan struct{}) func() {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middlewares.LogInterceptor()))
	reflection.Register(grpcServer) // 启用 gRPC Reflection

	go InitServer(grpcServer)

	return func() {
		go func() {
			grpcServer.GracefulStop()
			log.Printf("info: server shutting down")
			close(done)
		}()

		// 等待关闭完成或超时
		select {
		case <-done:
			log.Println("info: server stopped gracefully.")
		case <-time.After(160 * time.Second):
			log.Println("warning: timeout reached. Forcing shutdown.")
			grpcServer.Stop() // 强制关闭服务器
			close(done)
		}
	}
}

type Server struct {
	generated.UnimplementedAuthServiceServer
}

func InitServer(grpcServer *grpc.Server) {
	var (
		lis net.Listener
		err error
	)

	lis, err = net.Listen("tcp", addr)
	if err != nil {
		err = fmt.Errorf("error: failed to listen: %w", err)
		log.Fatalf("%v", err)
	}

	generated.RegisterAuthServiceServer(grpcServer, &Server{}) // 注册 Auth 服务
	log.Println("info: server is running on port ", addr)
	if err = grpcServer.Serve(lis); err != nil {
		err = fmt.Errorf("error: failed to serve: %w", err)
		log.Fatalf("%v", err)
	}
}
