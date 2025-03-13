package messaging

import (
	"context"

	"google.golang.org/protobuf/proto"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue/rabbitmq"
)

func SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error {
	_client := pkgMQ.GetClient(connStr)
	return _client.SendMessage(ctx, exchange, routeKey, req)
}

func ListenToQueue(exchange, queueName, routeKey string, handler HandlerFunc) {
	_client := pkgMQ.GetClient(connStr)
	_client.ListenToQueue(exchange, queueName, routeKey, handler)
}
