package rabbitmq

import (
	"context"

	"google.golang.org/protobuf/proto"
)

type MessageQueueMethod interface {
	Close(ctx context.Context)
	SendProtoMessage(ctx context.Context, exchange string, routeKey string, body []byte) error
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
	PreSendProtoMessage(ctx context.Context, exchange, queueName, routeKey string, body []byte) error
	PreSendMessage(ctx context.Context, exchange, queueName, routeKey string, req proto.Message) error
	GetMsgs(ctx context.Context, exchange, queueName, routeKey string, count int) [][]byte
	ListenToQueue(exchange, queueName, routeKey string, handler HandlerFunc)
	ListenRPC(exchange, queue, routeKey string, handler HandlerFuncWithReturn)
}
