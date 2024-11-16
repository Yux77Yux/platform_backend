package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	generatedUser "github.com/Yux77Yux/platform_backend/generated/user"
	"google.golang.org/grpc"
)

type Server struct {
	generatedUser.UnimplementedUserServiceServer
}

func ServerRun(ctx context.Context, done chan struct{}) {
	var (
		grpcServer *grpc.Server
		lis        net.Listener
		err        error
	)

	lis, err = net.Listen("tcp", ":8080")
	if err != nil {
		err = fmt.Errorf("error: failed to listen: %w", err)
		log.Fatalf("%v", err)
	}

	go func() {
		grpcServer = grpc.NewServer()
		generatedUser.RegisterUserServiceServer(grpcServer, &Server{}) // 注册 User 服务

		log.Println("info: server is running on port :50020")
		if err = grpcServer.Serve(lis); err != nil {
			err = fmt.Errorf("error: failed to serve: %w", err)
			log.Fatalf("%v", err)
		}
	}()

	// 等待信号
	sig := <-ctx.Done()

	go func() {
		log.Printf("info: received signal: %s. Shutting down gracefully...", sig)
		grpcServer.GracefulStop()
		close(done)
	}()

	// 等待关闭完成或超时
	select {
	case <-done:
		log.Println("info: server stopped gracefully.")
	case <-time.After(98 * time.Second):
		log.Println("warning: timeout reached. Forcing shutdown.")
		grpcServer.Stop() // 强制关闭服务器
		close(done)
	}
}
