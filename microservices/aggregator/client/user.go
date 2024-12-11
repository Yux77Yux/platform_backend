package client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user" // 你生成的 package 名字
)

type UserClient struct {
	connection *grpc.ClientConn
}

func NewUserClient() (*UserClient, error) {
	// 建立与服务器的连接
	conn, err := grpc.NewClient(user_service_address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	client := &UserClient{
		connection: conn,
	}

	return client, nil
}

func (c *UserClient) Login(credentials *generated.UserCredentials) (*generated.LoginResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := generated.NewUserServiceClient(c.connection)

	// 创建请求
	req := &generated.LoginRequest{UserCredentials: credentials}

	// 调用 gRPC 方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := client.Login(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *UserClient) GetUser(userId int64, accessToken *common.AccessToken) (*generated.GetUserResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := generated.NewUserServiceClient(c.connection)

	// 创建请求
	req := &generated.GetUserRequest{UserId: userId, AccessToken: accessToken}

	// 调用 gRPC 方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	response, err := client.GetUser(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}
