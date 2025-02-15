package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	generated "github.com/Yux77Yux/platform_backend/generated/creation" // 你生成的 package 名字
)

type CreationClient struct {
	connection *grpc.ClientConn
}

func NewCreationClient() (*CreationClient, error) {
	// 建立与服务器的连接
	conn, err := grpc.NewClient(service_address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	client := &CreationClient{
		connection: conn,
	}

	return client, nil
}

func (c *CreationClient) GetCreation(ctx context.Context, req *generated.GetCreationRequest) (*generated.GetCreationResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := generated.NewCreationServiceClient(c.connection)

	response, err := client.GetCreation(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *CreationClient) GetCreationList(ctx context.Context, req *generated.GetCreationListRequest) (*generated.GetCreationListResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := generated.NewCreationServiceClient(c.connection)

	response, err := client.GetCreationList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *CreationClient) GetPublicCreationList(ctx context.Context, req *generated.GetCreationListRequest) (*generated.GetCreationListResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := generated.NewCreationServiceClient(c.connection)

	response, err := client.GetPublicCreationList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *CreationClient) GetSpaceCreations(ctx context.Context, req *generated.GetSpaceCreationsRequest) (*generated.GetCreationListResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := generated.NewCreationServiceClient(c.connection)

	response, err := client.GetSpaceCreations(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}
