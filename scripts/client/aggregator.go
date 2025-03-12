package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	aggregator "github.com/Yux77Yux/platform_backend/generated/aggregator" // 你生成的 package 名字
	"github.com/google/uuid"
)

type AggregatorClient struct {
	connection *grpc.ClientConn
}

func NewAggregatorClient() (*AggregatorClient, error) {
	unaryInterceptor := grpc.WithUnaryInterceptor(TraceIDInterceptor)
	// 建立与服务器的连接
	conn, err := grpc.NewClient(service_address, grpc.WithTransportCredentials(insecure.NewCredentials()), unaryInterceptor)
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
	traceId := uuid.New().String()
	ctx = metadata.AppendToOutgoingContext(ctx, "TraceId", traceId)
	client := aggregator.NewAggregatorServiceClient(c.connection)
	return client.Login(ctx, req)
}
