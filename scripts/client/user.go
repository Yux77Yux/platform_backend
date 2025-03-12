package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	generated "github.com/Yux77Yux/platform_backend/generated/user" // 你生成的 package 名字
)

type UserClient struct {
	connection *grpc.ClientConn
}

func NewUserClient() (*UserClient, error) {
	unaryInterceptor := grpc.WithUnaryInterceptor(TraceIDInterceptor)
	// 建立与服务器的连接
	conn, err := grpc.NewClient(service_address, grpc.WithTransportCredentials(insecure.NewCredentials()), unaryInterceptor)
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	client := &UserClient{
		connection: conn,
	}

	return client, nil
}

func (c *UserClient) Close() {
	err := c.connection.Close()
	if err != nil {
		log.Printf("error: grpc client close %v", err)
	}
}

func (c *UserClient) Login(ctx context.Context, credentials *generated.UserCredentials) (*generated.LoginResponse, error) {
	// 创建客户端
	client := generated.NewUserServiceClient(c.connection)

	// 创建请求
	req := &generated.LoginRequest{UserCredentials: credentials}

	response, err := client.Login(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *UserClient) GetUser(ctx context.Context, userId int64) (*generated.GetUserResponse, error) {
	// 创建客户端
	client := generated.NewUserServiceClient(c.connection)

	// 创建请求
	req := &generated.GetUserRequest{UserId: userId}

	response, err := client.GetUser(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *UserClient) GetUsers(ctx context.Context, userIds []int64) (*generated.GetUsersResponse, error) {
	// 创建客户端
	client := generated.NewUserServiceClient(c.connection)

	// 创建请求
	req := &generated.GetUsersRequest{Ids: userIds}

	response, err := client.GetUsers(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *UserClient) Register(ctx context.Context, req *generated.RegisterRequest) (*generated.RegisterResponse, error) {
	// 创建客户端
	client := generated.NewUserServiceClient(c.connection)
	return client.Register(ctx, req)
}

func (c *UserClient) UpdateUserAvatar(ctx context.Context, req *generated.UpdateUserAvatarRequest) (*generated.UpdateUserAvatarResponse, error) {
	// 创建客户端
	client := generated.NewUserServiceClient(c.connection)
	return client.UpdateUserAvatar(ctx, req)
}

func (c *UserClient) UpdateUserSpace(ctx context.Context, req *generated.UpdateUserSpaceRequest) (*generated.UpdateUserResponse, error) {
	// 创建客户端
	client := generated.NewUserServiceClient(c.connection)
	return client.UpdateUserSpace(ctx, req)
}

func (c *UserClient) UpdateUserStatus(ctx context.Context, req *generated.UpdateUserStatusRequest) (*generated.UpdateUserResponse, error) {
	// 创建客户端
	client := generated.NewUserServiceClient(c.connection)
	return client.UpdateUserStatus(ctx, req)
}
