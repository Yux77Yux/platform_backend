package messaging

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue"
)

type MessagequeueInterface interface {
	Open(connStr string) error
	Close()
	ExchangeDeclare(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args amqp.Table) error
	QueueDeclare(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) (*amqp.Queue, error)
	ExchangeBind(destination string, key string, source string, noWait bool, args amqp.Table) error
	QueueBind(name string, key string, exchange string, noWait bool, args amqp.Table) error
	Publish(ctx context.Context, exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing) error
	Consume(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
}

var (
	exchangesConfig = map[string]string{
		"register_exchange": "direct",
		"login_exchange":    "direct",
		// Add more exchanges here
	}
)

var connStr string

func InitConnStr(_str string) {
	connStr = _str
}

func GetRabbitMQ() MessagequeueInterface {
	var rabbitMQ MessagequeueInterface = &pkgMQ.RabbitMQClass{}
	err := rabbitMQ.Open(connStr)
	wiredErr := fmt.Errorf("failed to connect the rabbit client: %w", err)
	log.Printf("error: %v", wiredErr)

	return rabbitMQ
}

func init() {
	rabbitMQ := GetRabbitMQ()

	for exchange, kind := range exchangesConfig {
		err := rabbitMQ.ExchangeDeclare(exchange, kind, true, false, false, false, nil)
		wiredErr := fmt.Errorf("failed to declare exchange %s : %w", exchange, err)
		log.Printf("error: %v", wiredErr)
	}
}
