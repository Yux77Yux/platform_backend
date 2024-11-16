package messagequeue

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQMethods interface {
	Close()
	ExchangeDeclare(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args amqp.Table) error
	QueueDeclare(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) (amqp.Queue, error)
	ExchangeBind(destination string, key string, source string, noWait bool, args amqp.Table) error
	QueueBind(name string, key string, exchange string, noWait bool, args amqp.Table) error
	Publish(ctx context.Context, exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing) error
	Consume(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
}

const (
	retries = 3
	delay   = 2 * time.Second
)

var (
	err error
)

func OpenRabbitMQ(connStr string) (*RabbitMQClass, error) {
	RabbitMQClient := &RabbitMQClass{}

	for i := 0; i < retries; i++ {
		RabbitMQClient.rabbitmqClient, err = amqp.Dial(connStr)
		if err == nil {
			RabbitMQClient.channel, err = RabbitMQClient.rabbitmqClient.Channel()
			if err != nil {
				return nil, err
			}

			return RabbitMQClient, nil
		}

		wiredErr := fmt.Errorf("failed to connect to RabbitMQ at %s: %w (attempt %d/%d). Retrying in %d seconds", connStr, err, i+1, retries, int(delay.Seconds()))
		log.Printf("error: %v\n", wiredErr)
		time.Sleep(delay)
	}

	return nil, err
}

type RabbitMQClass struct {
	rabbitmqClient *amqp.Connection
	channel        *amqp.Channel
}

func (r *RabbitMQClass) Close() {
	if err = r.channel.Close(); err != nil {
		wiredErr := fmt.Errorf("failed to close channel: %w", err)
		log.Printf("error: %v", wiredErr)
	}
}

func (r *RabbitMQClass) ExchangeDeclare(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args amqp.Table) error {
	switch kind {
	case "", "fanout", "direct", "topic":
	default:
		return fmt.Errorf("please choose one of \"\" fanout direct topic")
	}

	if err = r.channel.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args); err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQClass) QueueDeclare(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) (*amqp.Queue, error) {
	q, err := r.channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

func (r *RabbitMQClass) ExchangeBind(destination string, key string, source string, noWait bool, args amqp.Table) error {
	if err = r.channel.ExchangeBind(destination, key, source, noWait, args); err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQClass) QueueBind(name string, key string, exchange string, noWait bool, args amqp.Table) error {
	if err = r.channel.QueueBind(name, key, exchange, noWait, args); err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQClass) Publish(
	ctx context.Context,
	exchange string,
	key string,
	mandatory bool,
	immediate bool,
	msg amqp.Publishing) error {

	if err = r.channel.PublishWithContext(ctx, exchange, key, false, false, msg); err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQClass) Consume(
	queue string,
	consumer string,
	autoAck bool,
	exclusive bool,
	noLocal bool,
	noWait bool,
	args amqp.Table,
) (<-chan amqp.Delivery, error) {
	msg, err := r.channel.Consume(
		queue,     // queue
		consumer,  // consumer
		autoAck,   // auto ack
		exclusive, // exclusive
		noLocal,   // no local
		noWait,    // no wait
		args,      // args
	)

	if err != nil {
		return nil, err
	}

	return msg, nil
}
