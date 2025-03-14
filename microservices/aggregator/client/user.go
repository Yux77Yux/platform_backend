package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	generated "github.com/Yux77Yux/platform_backend/generated/user" // 你生成的 package 名字
)

type UserClient struct {
	connection *grpc.ClientConn
}

func NewUserClient() (*UserClient, error) {
	// 建立与服务器的连接
	conn, err := grpc.NewClient(service_address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect UserClient: %v", err)
	}

	client := &UserClient{
		connection: conn,
	}

	return client, nil
}

func (c *UserClient) Close() error {
	err := c.connection.Close()
	if err != nil {
		err = fmt.Errorf("error: UserClient close %w", err)
		return err
	}
	return nil
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
