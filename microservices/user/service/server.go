package service

import (
	"fmt"
	"log"
	"net"
	"time"

	generatedUser "github.com/Yux77Yux/platform_backend/generated/user"
	"google.golang.org/grpc"
)

func ServerRun(done chan struct{}) func() {
	grpcServer := grpc.NewServer()
	go InitServer(grpcServer)

	return func() {
		go func() {
			log.Printf("info: shutting down gracefully...")
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
}

type Server struct {
	generatedUser.UnimplementedUserServiceServer
}

func InitServer(grpcServer *grpc.Server) {
	var (
		lis net.Listener
		err error
	)

	lis, err = net.Listen("tcp", ":8080")
	if err != nil {
		err = fmt.Errorf("error: failed to listen: %w", err)
		log.Fatalf("%v", err)
	}

	go func() {
		generatedUser.RegisterUserServiceServer(grpcServer, &Server{}) // 注册 User 服务

		log.Println("info: server is running on port :50020")
		if err = grpcServer.Serve(lis); err != nil {
			err = fmt.Errorf("error: failed to serve: %w", err)
			log.Fatalf("%v", err)
		}
	}()

	forever := make(chan struct{})
	<-forever
}
