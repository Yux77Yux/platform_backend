package messaging

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

// (交换机名，请求队列，响应队列，路由键，请求id，
// 请求消息体)
func RPCPattern(
	exchange, queue, routeKey, correlationId string,
	req proto.Message,
) (*amqp.Delivery, error) {
	log.Printf("info: start rpc with exchange %s with routeKey %s", exchange, routeKey)

	const retries = 3
	var (
		err  error
		body []byte
	)

	ctx := context.Background()

	body, err = proto.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	rabbitMQ := GetRabbitMQ()

	if err := rabbitMQ.ExchangeDeclare(exchange, "direct", true, false, false, false, nil); err != nil {
		log.Printf("failed to declare exchange %s : %v", exchange, err)
		return nil, fmt.Errorf("failed to ExchangeDeclare: %w", err)
	}

	// 响应队列声明
	responseQueue, err := rabbitMQ.QueueDeclare("", false, true, true, false, nil)
	if err != nil {
		log.Printf("rabbitMQ QueueDeclare error: %v", err)
		return nil, fmt.Errorf("failed to QueueDeclare: %w", err)
	}

	for i := 0; i < retries; i++ {
		err := rabbitMQ.Publish(
			ctx,
			exchange,
			routeKey,
			false,
			false,
			amqp.Publishing{
				DeliveryMode:  amqp.Persistent,
				ContentType:   "application/x-protobuf",
				Body:          body,
				ReplyTo:       responseQueue.Name,
				CorrelationId: correlationId,
			})

		if err == nil {
			break
		}

		log.Printf("error: failed to publish message, retrying... (%d/3)\n", i+1)
		time.Sleep(time.Second * 1) // Wait before retrying
	}

	// 队列绑定交换机
	err = rabbitMQ.QueueBind(responseQueue.Name, responseQueue.Name, exchange, false, nil)
	if err != nil {
		log.Printf("rabbitMQ QueueBind error: %v", err)
		return nil, fmt.Errorf("failed to QueueBind: %w", err)
	}

	msgs, err := rabbitMQ.Consume(
		responseQueue.Name, // queue
		"",                 // consumer
		false,              // auto ack
		true,               // exclusive
		false,              // no local
		false,              // no wait
		nil,                // args
	)

	if err != nil {
		log.Println("error: ", responseQueue.Name, err.Error())
	}

	for msg := range msgs {
		if msg.CorrelationId == correlationId {
			msg.Ack(false) // Acknowledge successful processing
			return &msg, nil
		}
	}

	return nil, fmt.Errorf("failed to complete request")
}
