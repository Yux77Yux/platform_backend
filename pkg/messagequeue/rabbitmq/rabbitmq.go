package rabbitmq

import (
	"context"
	"fmt"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue"
	utils "github.com/Yux77Yux/platform_backend/pkg/utils"
)

var (
	connStr string
	once    sync.Once
)

type RabbitMQClient struct {
}

func GetClient(str string) RabbitMQInterface {
	once.Do(
		func() {
			connStr = str
		},
	)
	return &RabbitMQClient{}
}

func GetRabbitMQ() MessageQueueInterface {
	var messageQueue MessageQueueInterface = &pkgMQ.RabbitMQClass{}
	if err := messageQueue.Open(connStr); err != nil {
		err := fmt.Errorf("failed to connect the rabbit client: %w", err)
		utils.LogSuperError(err)

		return nil
	}

	return messageQueue
}

func (r *RabbitMQClient) SendProtoMessage(ctx context.Context, exchange string, routeKey string, body []byte) error {
	const retries = 3
	var err error

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

		time.Sleep(time.Second * 1) // Wait before retrying
	}
	return fmt.Errorf("failed to publish request: %w", err)
}

func (r *RabbitMQClient) SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error {
	traceId := utils.GetMetadataValue(ctx, "trace-id")
	fullName := utils.GetMetadataValue(ctx, "full-name")
	payload, err := anypb.New(req)
	if err != nil {
		return fmt.Errorf("failed to convert message to anypb.Any: %w", err)
	}

	reqWithCtx := &common.Wrapper{
		TraceId:  traceId,
		FullName: fullName,
		Payload:  payload,
	}

	body, err := proto.Marshal(reqWithCtx)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	return r.SendProtoMessage(ctx, exchange, routeKey, body)
}

func (r *RabbitMQClient) PreSendProtoMessage(ctx context.Context, exchange, queueName, routeKey string, body []byte) error {
	const retries = 3
	var err error

	rabbitMQ := GetRabbitMQ()
	defer rabbitMQ.Close()

	// 队列声明
	queue, err := rabbitMQ.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	// 在init中已经声明好交换机了
	// 队列绑定交换机
	err = rabbitMQ.QueueBind(queue.Name, routeKey, exchange, false, nil)
	if err != nil {
		return err
	}

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

		time.Sleep(time.Second * 1) // Wait before retrying
	}
	return fmt.Errorf("failed to publish request: %w", err)
}

func (r *RabbitMQClient) PreSendMessage(ctx context.Context, exchange, queueName, routeKey string, req proto.Message) error {
	traceId := utils.GetMetadataValue(ctx, "trace-id")
	fullName := utils.GetMetadataValue(ctx, "full-name")
	payload, err := anypb.New(req)
	if err != nil {
		return fmt.Errorf("failed to convert message to anypb.Any: %w", err)
	}

	reqWithCtx := &common.Wrapper{
		TraceId:  traceId,
		FullName: fullName,
		Payload:  payload,
	}

	body, err := proto.Marshal(reqWithCtx)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	return r.PreSendProtoMessage(ctx, exchange, queueName, routeKey, body)
}

func (r *RabbitMQClient) GetMsgs(exchange, queueName, routeKey string, count int) []amqp.Delivery {
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
		err = fmt.Errorf("rabbitMQ QueueDeclare error: %w", err)
		utils.LogSuperError(err)
		return nil
	}

	// 在init中已经声明好交换机了
	// 队列绑定交换机
	err = rabbitMQ.QueueBind(queue.Name, routeKey, exchange, false, nil)
	if err != nil {
		err = fmt.Errorf("rabbitMQ QueueBind error: %w", err)
		utils.LogSuperError(err)
		return nil
	}

	msgs, err = rabbitMQ.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	if err != nil {
		err = fmt.Errorf("rabbitMQ Consume error: %w", err)
		utils.LogSuperError(err)
		return nil
	}

	values := make([]amqp.Delivery, 0, count)
	for i := 0; i < count; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		select {
		case msg, ok := <-msgs:
			if !ok { // 如果通道关闭，提前退出
				cancel()
				return values
			}
			msg.Ack(false)
			values = append(values, msg)
			cancel()
		case <-ctx.Done():
			cancel()
			return values
		}
	}
	return values
}

func (r *RabbitMQClient) ListenToQueue(exchange, queueName, routeKey string, handler HandlerFunc) {
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
		err = fmt.Errorf("rabbitMQ QueueDeclare error: %w", err)
		utils.LogSuperError(err)
		return
	}

	// 在init中已经声明好交换机了
	// 队列绑定交换机
	err = rabbitMQ.QueueBind(queue.Name, routeKey, exchange, false, nil)
	if err != nil {
		err = fmt.Errorf("rabbitMQ QueueBind error: %w", err)
		utils.LogSuperError(err)
		return
	}

	msgs, err = rabbitMQ.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	if err != nil {
		err = fmt.Errorf("rabbitMQ Consume error: %w", err)
		utils.LogSuperError(err)
		return
	}

	for msg := range msgs {
		var wrap common.Wrapper
		err := proto.Unmarshal(msg.Body, &wrap)
		if err != nil {
			err = fmt.Errorf("error: message processing failed: %w", err)
			utils.LogError("", "", err.Error())
		}

		fullName := wrap.GetFullName()
		traceId := wrap.GetTraceId()
		utils.LogInfo(traceId, fullName)
		payload := wrap.GetPayload()

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
		err = handler(ctx, payload)
		cancel()
		if err != nil {
			msg.Nack(false, false)
			utils.LogError(traceId, fullName, err.Error())
		} else {
			msg.Ack(false)
			utils.LogInfo(traceId, fullName)
		}
	}
}

func (r *RabbitMQClient) ListenRPC(exchange, queue, routeKey string, handler HandlerFuncWithReturn) {
	rabbitMQ := GetRabbitMQ()
	defer rabbitMQ.Close()

	if err := rabbitMQ.ExchangeDeclare(exchange, "direct", true, false, false, false, nil); err != nil {
		err = fmt.Errorf("failed to declare exchange %s : %w", exchange, err)
		utils.LogSuperError(err)
		return
	}

	// 请求队列声明
	requestQueue, err := rabbitMQ.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		err = fmt.Errorf("rabbitMQ QueueDeclare error: %w", err)
		utils.LogSuperError(err)
		return
	}

	// 队列绑定交换机
	err = rabbitMQ.QueueBind(requestQueue.Name, routeKey, exchange, false, nil)
	if err != nil {
		err = fmt.Errorf("rabbitMQ QueueDeclare error: %w", err)
		utils.LogSuperError(err)
		return
	}

	msgs, err := rabbitMQ.Consume(
		requestQueue.Name, // queue
		"",                // consumer
		false,             // auto ack
		true,              // exclusive
		false,             // no local
		false,             // no wait
		nil,               // args
	)
	if err != nil {
		err = fmt.Errorf("rabbitMQ Consume error: %w", err)
		utils.LogSuperError(err)
		return
	}

	for msg := range msgs {
		var wrap common.Wrapper
		err := proto.Unmarshal(msg.Body, &wrap)
		if err != nil {
			err = fmt.Errorf("error: message processing failed: %w", err)
			utils.LogError("", "", err.Error())
		}

		fullName := wrap.GetFullName()
		traceId := wrap.GetTraceId()
		utils.LogInfo(traceId, fullName)
		payload := wrap.GetPayload()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		responseMsg, err := handler(ctx, payload)
		cancel()
		if err != nil {
			err = fmt.Errorf("error: message processing failed: %w", err)
			utils.LogError("", "", err.Error())
		}
		body, err := proto.Marshal(responseMsg)
		if err != nil {
			err = fmt.Errorf("failed to marshal request: %w", err)
			utils.LogError(traceId, fullName, err.Error())
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
			msg.Nack(false, false)
			err = fmt.Errorf("failed to marshal request: %w", err)
			utils.LogError(traceId, fullName, err.Error())
		} else {
			msg.Ack(false) // Acknowledge successful processing
		}
	}
}
