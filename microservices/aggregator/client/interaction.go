package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction" // 你生成的 package 名字
)

type InteractionClient struct {
	connection *grpc.ClientConn
}

func NewInteractionClient() (*InteractionClient, error) {
	// 建立与服务器的连接
	conn, err := grpc.NewClient(service_address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect InteractionClient: %v", err)
	}

	client := &InteractionClient{
		connection: conn,
	}

	return client, nil
}

func (c *InteractionClient) Close() error {
	err := c.connection.Close()
	if err != nil {
		err = fmt.Errorf("error: InteractionClient close %w", err)
		return err
	}
	return nil
}

func (c *InteractionClient) GetActionTag(ctx context.Context, req *generated.GetCreationInteractionRequest) (*generated.GetCreationInteractionResponse, error) {
	// 创建客户端
	client := generated.NewInteractionServiceClient(c.connection)

	response, err := client.GetActionTag(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *InteractionClient) GetHistories(ctx context.Context, req *generated.GetHistoriesRequest) (*generated.GetInteractionsResponse, error) {
	// 创建客户端
	client := generated.NewInteractionServiceClient(c.connection)

	response, err := client.GetHistories(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *InteractionClient) GetCollections(ctx context.Context, req *generated.GetCollectionsRequest) (*generated.GetInteractionsResponse, error) {
	// 创建客户端
	client := generated.NewInteractionServiceClient(c.connection)

	response, err := client.GetCollections(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *InteractionClient) GetRecommendBaseUser(ctx context.Context, req *generated.GetRecommendRequest) (*generated.GetRecommendResponse, error) {
	// 创建客户端
	client := generated.NewInteractionServiceClient(c.connection)

	response, err := client.GetRecommendBaseUser(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *InteractionClient) GetRecommendBaseCreation(ctx context.Context, req *generated.GetRecommendRequest) (*generated.GetRecommendResponse, error) {
	// 创建客户端
	client := generated.NewInteractionServiceClient(c.connection)

	response, err := client.GetRecommendBaseCreation(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *InteractionClient) PostInteraction(ctx context.Context, req *generated.PostInteractionRequest) (*generated.PostInteractionResponse, error) {
	// 创建客户端
	client := generated.NewInteractionServiceClient(c.connection)

	response, err := client.PostInteraction(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}
