package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	generated "github.com/Yux77Yux/platform_backend/generated/comment" // 你生成的 package 名字
)

type CommentClient struct {
	connection *grpc.ClientConn
}

func NewCommentClient() (*CommentClient, error) {
	// 建立与服务器的连接
	conn, err := grpc.NewClient(service_address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	client := &CommentClient{
		connection: conn,
	}

	return client, nil
}

func (c *CommentClient) GetComment(ctx context.Context, req *generated.GetCommentRequest) (*generated.GetCommentResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := generated.NewCommentServiceClient(c.connection)

	response, err := client.GetComment(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *CommentClient) InitialComments(ctx context.Context, req *generated.InitialCommentsRequest) (*generated.InitialCommentsResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := generated.NewCommentServiceClient(c.connection)

	response, err := client.InitialComments(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *CommentClient) GetComments(ctx context.Context, req *generated.GetCommentsRequest) (*generated.GetCommentsResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := generated.NewCommentServiceClient(c.connection)

	response, err := client.GetComments(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *CommentClient) GetTopComments(ctx context.Context, req *generated.GetTopCommentsRequest) (*generated.GetCommentsResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := generated.NewCommentServiceClient(c.connection)

	response, err := client.GetTopComments(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *CommentClient) GetSecondComments(ctx context.Context, req *generated.GetSecondCommentsRequest) (*generated.GetCommentsResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := generated.NewCommentServiceClient(c.connection)

	response, err := client.GetSecondComments(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}

func (c *CommentClient) GetReplyComments(ctx context.Context, req *generated.GetReplyCommentsRequest) (*generated.GetCommentsResponse, error) {
	defer c.connection.Close()
	// 创建客户端
	client := generated.NewCommentServiceClient(c.connection)

	response, err := client.GetReplyComments(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not greet: %v", err)
	}

	return response, nil
}
