package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

type MessageQueueInterface interface {
	Open(connStr string) error
	Close() error
	ExchangeDeclare(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args amqp.Table) error
	QueueDeclare(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) (*amqp.Queue, error)
	ExchangeBind(destination string, key string, source string, noWait bool, args amqp.Table) error
	QueueBind(name string, key string, exchange string, noWait bool, args amqp.Table) error
	Publish(ctx context.Context, exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing) error
	Consume(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
}

type RabbitMQInterface interface {
	SendProtoMessage(ctx context.Context, exchange string, routeKey string, body []byte) error
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
	PreSendProtoMessage(ctx context.Context, exchange, queueName, routeKey string, body []byte) error
	PreSendMessage(ctx context.Context, exchange, queueName, routeKey string, req proto.Message) error
	GetMsgs(exchange, queueName, routeKey string, count int) []amqp.Delivery
	ListenToQueue(exchange, queueName, routeKey string, handler HandlerFunc)
	ListenRPC(exchange, queue, routeKey string, handler HandlerFuncWithReturn)
}
