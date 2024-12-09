package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/auth"
	internal "github.com/Yux77Yux/platform_backend/microservices/auth/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

func ServerRun(done chan struct{}) func() {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(refreshTokenInterceptor()),
	)
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

	go func() {
		generated.RegisterAuthServiceServer(grpcServer, &Server{}) // 注册 Auth 服务

		log.Println("info: server is running on port ", addr)
		if err = grpcServer.Serve(lis); err != nil {
			err = fmt.Errorf("error: failed to serve: %w", err)
			log.Fatalf("%v", err)
		}
	}()

	forever := make(chan struct{})
	<-forever
}

// 拦截cookie，使其呈现
func refreshTokenInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		// 获取请求的元数据（Headers）
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			log.Println("failed to get metadata")
			return handler(ctx, req)
		}

		// 从 headers 中获取 Cookie 信息
		cookies := md.Get("cookie")
		if len(cookies) == 0 {
			log.Println("missing cookies")
			return handler(ctx, req)
		}

		// 假设 cookies 数组只有一个元素，包含所有 cookies
		cookieStr := cookies[0]

		// 使用分号拆分多个 cookie
		cookieParts := strings.Split(cookieStr, ";")
		// 去除每个 cookie 部分的空白字符
		refreshToken := ""
		for _, part := range cookieParts {
			part = strings.TrimSpace(part) // 去除前后的空白字符
			if strings.HasPrefix(part, "refreshToken=") {
				// 提取 refreshToken 的值
				refreshToken = strings.Split(part, "=")[1]
				break // 找到后退出循环
			}
		}

		if refreshToken == "" {
			log.Println("refresh token not found")
			return handler(ctx, req)
		}

		// 将 refreshToken 放入 ctx 中
		ctx = context.WithValue(ctx, internal.RefreshTokenKey, refreshToken)

		// 继续执行 gRPC handler
		return handler(ctx, req)
	}
}
