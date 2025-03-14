package messaging

import (
	"context"

	"google.golang.org/protobuf/proto"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue/rabbitmq"
)

var _client MessageQueueInterface

func SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error {
	return _client.SendMessage(ctx, exchange, routeKey, req)
}

func ListenToQueue(exchange, queueName, routeKey string, handler HandlerFunc) {
	_client.ListenToQueue(exchange, queueName, routeKey, handler)
}

func Init() {
	_client = pkgMQ.GetClient(connStr)
}

func Close(ctx context.Context) {
	_client.Close(ctx)
}
