package service

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
	tools "github.com/Yux77Yux/platform_backend/microservices/review/tools"
	middlewares "github.com/Yux77Yux/platform_backend/pkg/middlewares"
)

func ServerRun() func(chan any) {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middlewares.LogInterceptor()))
	reflection.Register(grpcServer)

	go InitServer(grpcServer)

	return func(done chan any) {
		go func() {
			grpcServer.GracefulStop()
			log.Printf("info: server shutting down")
			close(done)
		}()

		// 等待关闭完成或超时
		select {
		case <-done:
			tools.LogInfo("GrpcServer", "Server stopped gracefully")
		case <-time.After(2 * time.Minute):
			grpcServer.Stop()
			tools.LogWarning("GrpcServer", "Server Stop", "ServerRun timeout reached. Forcing shutdown")
		}
	}
}

type Server struct {
	generated.UnimplementedReviewServiceServer
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

	generated.RegisterReviewServiceServer(grpcServer, &Server{}) // 注册 User 服务
	log.Println("info: review server is running on port ", addr)
	if err = grpcServer.Serve(lis); err != nil {
		err = fmt.Errorf("error: failed to serve: %w", err)
		log.Fatalf("%v", err)
	}
}
