package messaging

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue/rabbitmq"
)

var _client MessageQueueInterface

func GetMsgs(exchange, queueName, routeKey string, count int) []amqp.Delivery {
	return _client.GetMsgs(exchange, queueName, routeKey, count)
}

func PreSendMessage(ctx context.Context, exchange, queueName, routeKey string, req proto.Message) error {
	return _client.PreSendMessage(ctx, exchange, queueName, routeKey, req)
}

func PreSendProtoMessage(ctx context.Context, exchange, queueName, routeKey string, req []byte) error {
	return _client.PreSendProtoMessage(ctx, exchange, queueName, routeKey, req)
}

func SendProtoMessage(ctx context.Context, exchange string, routeKey string, req []byte) error {
	return _client.SendProtoMessage(ctx, exchange, routeKey, req)
}

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
