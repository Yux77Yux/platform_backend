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

	// 队列声明
	queue, err = rabbitMQ.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Printf("rabbitMQ QueueDeclare error: %v", err)
		return
	}

	// 在init中已经声明好交换机了
	// 队列绑定交换机
	err = rabbitMQ.QueueBind(queue.Name, routeKey, exchange, false, nil)
	if err != nil {
		log.Printf("rabbitMQ QueueBind error: %v", err)
		return
	}

	msgs, err = rabbitMQ.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto ack
		true,       // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	if err != nil {
		log.Printf("rabbitMQ Consume error: %v", err)
		return
	}

	for msg := range msgs {
		if err := handler(msg); err != nil {
			log.Printf("error: message processing failed: %v", err)
			msg.Nack(false, false) // Negatively acknowledge
		} else {
			msg.Ack(false) // Acknowledge successful processing
		}
	}
}

func ListenRPC(exchange, queue, routeKey string, handler func(amqp.Delivery) (proto.Message, error)) {
	log.Printf("info: start consume message on exchange %s queue %s with routeKey %s", exchange, queue, routeKey)

	rabbitMQ := GetRabbitMQ()
	defer rabbitMQ.Close()

	if err := rabbitMQ.ExchangeDeclare(exchange, "direct", true, false, false, false, nil); err != nil {
		log.Printf("failed to declare exchange %s : %v", exchange, err)
		return
	}

	// 请求队列声明
	requestQueue, err := rabbitMQ.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		log.Printf("rabbitMQ QueueDeclare error: %v", err)
		return
	}

	// 队列绑定交换机
	err = rabbitMQ.QueueBind(requestQueue.Name, routeKey, exchange, false, nil)
	if err != nil {
		log.Printf("rabbitMQ QueueBind error: %v", err)
		return
	}

	msgs, err := rabbitMQ.Consume(
		requestQueue.Name, // queue
		"",                // consumer
		false,             // auto ack
		false,             // exclusive
		false,             // no local
		false,             // no wait
		nil,               // args
	)
	if err != nil {
		log.Printf("rabbitMQ Consume error: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	for msg := range msgs {
		responseMsg, err := handler(msg)
		if err != nil {
			log.Printf("error: message processing failed: %v", err)
		}
		body, err := proto.Marshal(responseMsg)
		if err != nil {
			log.Printf("failed to marshal request: %v", err)
		}

		err = rabbitMQ.Publish(
			ctx,
			exchange,
			msg.ReplyTo,
			false,
			false,
			amqp.Publishing{
				DeliveryMode:  amqp.Persistent,
				ContentType:   "application/x-protobuf",
				Body:          body,
				CorrelationId: msg.CorrelationId,
			})

		if err != nil {
			log.Printf("error:listen rpc publish : %v", err)
			msg.Nack(false, false)
		} else {
			log.Printf("%s publish success", msg.CorrelationId)
			msg.Ack(false) // Acknowledge successful processing
		}
	}
}
