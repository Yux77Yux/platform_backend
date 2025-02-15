package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	generated "github.com/Yux77Yux/platform_backend/generated/review" // 你生成的 package 名字
)

type ReviewClient struct {
	connection *grpc.ClientConn
}

func NewReviewClient() (*ReviewClient, error) {
	// 建立与服务器的连接
	conn, err := grpc.NewClient(service_address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	client := &ReviewClient{
		connection: conn,
	}

	return client, nil
}

func (c *ReviewClient) GetReviews(ctx context.Context, req *generated.GetReviewsRequest) (*generated.GetReviewsResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := generated.NewReviewServiceClient(c.connection)

	response, err := client.GetReviews(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *ReviewClient) GetNewReviews(ctx context.Context, req *generated.GetNewReviewsRequest) (*generated.GetReviewsResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := generated.NewReviewServiceClient(c.connection)

	response, err := client.GetNewReviews(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}
