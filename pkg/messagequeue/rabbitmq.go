package messagequeue

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	retries = 3
	delay   = 2 * time.Second
)

type RabbitMQClass struct {
	rabbitmqClient *amqp.Connection
	channel        *amqp.Channel
}

func (r *RabbitMQClass) Open(connStr string) error {
	var (
		err error
	)

	for i := 0; i < retries; i++ {
		r.rabbitmqClient, err = amqp.Dial(connStr)
		if err == nil {
			r.channel, err = r.rabbitmqClient.Channel()
			if err != nil {
				return fmt.Errorf("can't open the rabbitmq with channel because %w", err)
			}

			return nil
		}

		wiredErr := fmt.Errorf("failed to connect to RabbitMQ at %s: %w (attempt %d/%d). Retrying in %d seconds", connStr, err, i+1, retries, int(delay.Seconds()))
		log.Printf("error: %v\n", wiredErr)
		time.Sleep(delay)
	}

	return fmt.Errorf("can't connect the rabbitmq because %w", err)
}

func (r *RabbitMQClass) Close() error {
	if err := r.channel.Close(); err != nil {
		return fmt.Errorf("can't close channel because %w", err)
	}
	return nil
}

func (r *RabbitMQClass) ExchangeDeclare(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args amqp.Table) error {
	switch kind {
	case "", "fanout", "direct", "topic":
	default:
		return fmt.Errorf("please choose one of \"\" fanout direct topic")
	}

	if err := r.channel.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args); err != nil {
		return fmt.Errorf("can't decalre a exchange %s kind %s because %w", name, kind, err)
	}

	return nil
}

func (r *RabbitMQClass) QueueDeclare(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) (*amqp.Queue, error) {
	q, err := r.channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue %s: %w", name, err)
	}

	return &q, nil
}

func (r *RabbitMQClass) ExchangeBind(destination string, key string, source string, noWait bool, args amqp.Table) error {
	if err := r.channel.ExchangeBind(destination, key, source, noWait, args); err != nil {
		return fmt.Errorf("can't bind with destination %s key %s source %s because %w", destination, key, source, err)
	}

	return nil
}

func (r *RabbitMQClass) QueueBind(name string, key string, exchange string, noWait bool, args amqp.Table) error {
	if err := r.channel.QueueBind(name, key, exchange, noWait, args); err != nil {
		return fmt.Errorf("failed with queue %s bind the routeKey %s of exchange %s: %w", name, key, exchange, err)
	}

	return nil
}

func (r *RabbitMQClass) Publish(ctx context.Context, exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing) error {
	if err := r.channel.PublishWithContext(ctx, exchange, key, false, false, msg); err != nil {
		return fmt.Errorf("can't publish with exchange %s key %s because %w", exchange, key, err)
	}

	return nil
}

func (r *RabbitMQClass) Consume(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
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
		return nil, fmt.Errorf("failed to consume with queue %s: %w", queue, err)
	}

	return msg, nil
}
