package messaging

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
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

type DispatcherInterface interface {
	Start()
	GetChannel() chan RequestHandlerFunc
	Shutdown()
}
