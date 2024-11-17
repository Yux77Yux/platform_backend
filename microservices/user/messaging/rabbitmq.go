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

	queue, err = rabbitMQ.QueueDeclare(queueName, true, false, true, false, nil)
	wiredErr := fmt.Errorf("failed to declare a queue %s: %w", queueName, err)
	log.Printf("error: %v", wiredErr)

	err = rabbitMQ.QueueBind(queue.Name, routeKey, exchange, false, nil)
	wiredErr = fmt.Errorf("failed with queue %s bind the routeKey %s of exchange %s: %w", queue.Name, routeKey, exchange, err)
	log.Printf("error: %v", wiredErr)

	msgs, err = rabbitMQ.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto ack
		true,       // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	wiredErr = fmt.Errorf("failed to consume with queue %s: %w", queue.Name, err)
	log.Printf("error: %v", wiredErr)

	go func() {
		for d := range msgs {
			if err := handler(d); err != nil {
				log.Printf("error: message processing failed: %v", err)
				d.Nack(false, false) // Negatively acknowledge
			} else {
				d.Ack(false) // Acknowledge successful processing
			}
		}
	}()
}
