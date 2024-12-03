package messaging

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

func SendMessage(exchange string, routeKey string, req proto.Message) error {
	log.Printf("info: start send message to exchange %s with routeKey %s", exchange, routeKey)
	const retries = 3
	var (
		err  error
		body []byte
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err = proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	rabbitMQ := GetRabbitMQ()
	defer rabbitMQ.Close()

	for i := 0; i < retries; i++ {
		err := rabbitMQ.Publish(
			ctx,
			exchange,
			routeKey,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/x-protobuf",
				Body:         body,
			})

		if err == nil {
			return nil
		}

		log.Printf("error: failed to publish message, retrying... (%d/3)\n", i+1)
		time.Sleep(time.Second * 2) // Wait before retrying
	}

	return fmt.Errorf("failed to publish request: %w", err)
}

func ListenToQueue(exchange, queueName, routeKey string, handler func(d amqp.Delivery) error) {
	log.Printf("info: start consume message on exchange %s queue %s with routeKey %s", exchange, queueName, routeKey)
	var (
		queue *amqp.Queue
		msgs  <-chan amqp.Delivery
		err   error
	)

	rabbitMQ := GetRabbitMQ()
	defer rabbitMQ.Close()

	queue, err = rabbitMQ.QueueDeclare(queueName, true, false, true, false, nil)
	log.Printf("error: %v", err)

	err = rabbitMQ.QueueBind(queue.Name, routeKey, exchange, false, nil)
	log.Printf("error: %v", err)

	msgs, err = rabbitMQ.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto ack
		true,       // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	log.Printf("error: %v", err)

	go func() {
		for msg := range msgs {
			if err := handler(msg); err != nil {
				log.Printf("error: message processing failed: %v", err)
				msg.Nack(false, false) // Negatively acknowledge
			} else {
				msg.Ack(false) // Acknowledge successful processing
			}
		}
	}()
}
