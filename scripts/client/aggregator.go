package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	aggregator "github.com/Yux77Yux/platform_backend/generated/aggregator" // 你生成的 package 名字
)

type AggregatorClient struct {
	connection *grpc.ClientConn
}

func NewAggregatorClient() (*AggregatorClient, error) {
	// 建立与服务器的连接
	conn, err := grpc.NewClient(service_address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	client := &AggregatorClient{
		connection: conn,
	}

	return client, nil
}

func (c *AggregatorClient) Close() {
	err := c.connection.Close()
	if err != nil {
		log.Printf("error: grpc client close %v", err)
	}
}

func (c *AggregatorClient) Login(ctx context.Context, req *aggregator.LoginRequest) (*aggregator.LoginResponse, error) {
	// 创建客户端
	client := aggregator.NewAggregatorServiceClient(c.connection)
	return client.Login(ctx, req)
}
