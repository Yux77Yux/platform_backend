package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"

	user "github.com/Yux77Yux/platform_backend/generated/user" // 你生成的 package 名字
)

type UserClient struct {
	connection *grpc.ClientConn
}

func NewUserClient() (*UserClient, error) {
	// 建立与服务器的连接
	conn, err := grpc.NewClient(user_service_address)
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	client := &UserClient{
		connection: conn,
	}

	return client, nil
}

func (c *UserClient) Login(credentials *user.UserCredentials) (*user.LoginResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := user.NewUserServiceClient(c.connection)

	// 创建请求
	req := &user.LoginRequest{UserCredentials: credentials}

	// 调用 gRPC 方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := client.Login(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}
