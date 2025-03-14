package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	tools "github.com/Yux77Yux/platform_backend/microservices/creation/tools"
	middlewares "github.com/Yux77Yux/platform_backend/pkg/middlewares"
)

func ServerRun(ctx context.Context) {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middlewares.LogInterceptor()))
	reflection.Register(grpcServer) // 启用 gRPC Reflection

	go InitServer(grpcServer)

	<-ctx.Done()
	done := make(chan any, 1)
	go func() {
		grpcServer.GracefulStop()
		log.Printf("info: server shutting down")
		close(done)
	}()
	traceId := tools.GetMainValue(ctx)

	// 等待关闭完成或超时
	select {
	case <-done:
		tools.LogInfo(traceId, "ServerRun stopped gracefully")
	case <-time.After(time.Minute):
		grpcServer.Stop()
		tools.LogWarning(traceId, "ServerRun", "timeout reached. Forcing shutdown")
	}
}

type Server struct {
	generated.UnimplementedCreationServiceServer
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

	generated.RegisterCreationServiceServer(grpcServer, &Server{}) // 注册 User 服务
	log.Println("info: creation server is running on port ", addr)
	if err = grpcServer.Serve(lis); err != nil {
		err = fmt.Errorf("error: failed to serve: %w", err)
		log.Fatalf("%v", err)
	}
}
