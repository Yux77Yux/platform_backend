package messaging

import (
	"context"

	"google.golang.org/protobuf/proto"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue/rabbitmq"
)

type HandlerFunc = pkgMQ.HandlerFunc

type MessageQueueInterface interface {
	Close(ctx context.Context)
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
	ListenToQueue(exchange, queueName, routeKey string, handler HandlerFunc)
}
