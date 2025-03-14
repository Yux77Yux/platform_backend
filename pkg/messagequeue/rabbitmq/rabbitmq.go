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
	utils "github.com/Yux77Yux/platform_backend/pkg/utils"
)

const (
	_COUNT    = 20
	MAX_RETRY = 3
)

var (
	_client *RabbitMQClient
)

func GetClient(str string) MessageQueueInterface {
	if _client != nil {
		return _client
	}

	var err error
	connStr := str
	client := &RabbitMQClient{
		listenCh: make(map[string]*amqp.Channel),
	}
	for i := 0; i < MAX_RETRY; i++ {
		client.rabbitmqClient, err = amqp.Dial(connStr)
		if err == nil {
			client.chPool = sync.Pool{
				New: func() any {
					ch := client.newChannel()
					return ch
				},
			}
			_client = client
			break
		}
		if i == MAX_RETRY {
			utils.LogSuperError(err)
		}
		time.Sleep(MAX_RETRY * time.Second)
	}
	return _client
}

type RabbitMQClient struct {
	rabbitmqClient *amqp.Connection
	chPool         sync.Pool
	listenCh       map[string]*amqp.Channel
}

func (r *RabbitMQClient) Close(ctx context.Context) {
	traceID := utils.GetMainValue(ctx)
	for _, ch := range r.listenCh {
		if err := ch.Close(); err != nil {
			utils.LogError(traceID, "RabbitMQChannel.Close", err)
		}
	}
	if err := r.rabbitmqClient.Close(); err != nil {
		utils.LogError(traceID, "RabbitMQClient.Close", err)
	}
}

// 把 channel 放回池子
func (r *RabbitMQClient) putChannel(ch *amqp.Channel) {
	if ch != nil {
		r.chPool.Put(ch)
	}
}

func (r *RabbitMQClient) getChannel() *amqp.Channel {
	ch := r.chPool.Get()
	if ch == nil {
		utils.LogSuperError(fmt.Errorf("failed to get channel from pool"))
		return nil
	}

	channel, ok := ch.(*amqp.Channel)
	if !ok {
		utils.LogSuperError(fmt.Errorf("invalid channel type retrieved from pool"))
		return nil
	}

	return channel
}

func (r *RabbitMQClient) newChannel() *amqp.Channel {
	ch, err := r.rabbitmqClient.Channel()
	if err != nil {
		utils.LogSuperError(err)
		return nil
	}
	if ch == nil {
		utils.LogSuperError(fmt.Errorf("failed to new channel"))
	}
	return ch
}

func (r *RabbitMQClient) SendProtoMessage(ctx context.Context, exchange string, routeKey string, body []byte) error {
	ch := r.getChannel()
	defer r.putChannel(ch)
	var err error
	for i := 0; i < MAX_RETRY; i++ {
		err = ch.PublishWithContext(
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

		time.Sleep(time.Second * MAX_RETRY) // Wait before retrying
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
	ch := r.getChannel()
	defer r.putChannel(ch)

	// 队列声明
	queue, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	// 在init中已经声明好交换机了
	// 队列绑定交换机
	err = ch.QueueBind(queue.Name, routeKey, exchange, false, nil)
	if err != nil {
		return err
	}

	for i := 0; i < MAX_RETRY; i++ {
		err := ch.PublishWithContext(
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

		time.Sleep(time.Second * MAX_RETRY) // Wait before retrying
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
	ch := r.getChannel()
	defer r.putChannel(ch)

	// 队列声明
	queue, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		utils.LogSuperError(fmt.Errorf("rabbitMQ QueueDeclare error: %w", err))
		return nil
	}

	// 在init中已经声明好交换机了
	// 队列绑定交换机
	err = ch.QueueBind(queue.Name, routeKey, exchange, false, nil)
	if err != nil {
		err = fmt.Errorf("rabbitMQ QueueBind error: %w", err)
		utils.LogSuperError(err)
		return nil
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	if err != nil {
		utils.LogSuperError(fmt.Errorf("rabbitMQ Consume error: %w", err))
		return nil
	}

	values := make([]amqp.Delivery, 0, count)
	for i := 0; i < count; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
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

func (r *RabbitMQClient) storeInMap(key string, ch *amqp.Channel) {
	r.listenCh[key] = ch
}

func (r *RabbitMQClient) ListenToQueue(exchange, queueName, routeKey string, handler HandlerFunc) {
	ch := r.getChannel()
	r.storeInMap(fmt.Sprintf("%s_%s_%s", exchange, queueName, routeKey), ch)

	if err := ch.ExchangeDeclare(exchange, "direct", true, false, false, false, nil); err != nil {
		utils.LogSuperError(fmt.Errorf("failed to declare exchange %s : %w", exchange, err))
		return
	}

	// 队列声明
	queue, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		utils.LogSuperError(fmt.Errorf("rabbitMQ QueueDeclare error: %w", err))
		return
	}

	// 在init中已经声明好交换机了
	// 队列绑定交换机
	err = ch.QueueBind(queue.Name, routeKey, exchange, false, nil)
	if err != nil {
		utils.LogSuperError(fmt.Errorf("rabbitMQ QueueBind error: %w", err))
		return
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	if err != nil {
		utils.LogSuperError(fmt.Errorf("rabbitMQ Consume error: %w", err))
		return
	}

	go func() {
		for msg := range msgs {
			var wrap common.Wrapper
			err := proto.Unmarshal(msg.Body, &wrap)
			if err != nil {
				err = fmt.Errorf("error: message processing failed: %w", err)
				utils.LogError("", "", err)
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
				utils.LogError(traceId, fullName, err)
			} else {
				msg.Ack(false)
				utils.LogInfo(traceId, fullName)
			}
		}
	}()
}

func (r *RabbitMQClient) ListenRPC(exchange, queue, routeKey string, handler HandlerFuncWithReturn) {
	ch := r.getChannel()
	r.storeInMap(fmt.Sprintf("%s_%s_%s", exchange, queue, routeKey), ch)

	if err := ch.ExchangeDeclare(exchange, "direct", true, false, false, false, nil); err != nil {
		utils.LogSuperError(fmt.Errorf("failed to declare exchange %s : %w", exchange, err))
		return
	}

	// 请求队列声明
	requestQueue, err := ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		utils.LogSuperError(fmt.Errorf("rabbitMQ QueueDeclare error: %w", err))
		return
	}

	// 队列绑定交换机
	err = ch.QueueBind(requestQueue.Name, routeKey, exchange, false, nil)
	if err != nil {
		utils.LogSuperError(fmt.Errorf("rabbitMQ QueueDeclare error: %w", err))
		return
	}

	msgs, err := ch.Consume(
		requestQueue.Name, // queue
		"",                // consumer
		false,             // auto ack
		true,              // exclusive
		false,             // no local
		false,             // no wait
		nil,               // args
	)
	if err != nil {
		utils.LogSuperError(fmt.Errorf("rabbitMQ Consume error: %w", err))
		return
	}

	for msg := range msgs {
		var wrap common.Wrapper
		err := proto.Unmarshal(msg.Body, &wrap)
		if err != nil {
			err = fmt.Errorf("error: message processing failed: %w", err)
			utils.LogError("", "proto.Unmarshal", err)
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
			utils.LogError(traceId, utils.GetFunctionName(handler), err)
		}
		body, err := proto.Marshal(responseMsg)
		if err != nil {
			err = fmt.Errorf("failed to marshal request: %w", err)
			utils.LogError(traceId, fullName, err)
		}

		err = ch.PublishWithContext(
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
			utils.LogError(traceId, fullName, err)
		} else {
			msg.Ack(false) // Acknowledge successful processing
		}
	}
}
