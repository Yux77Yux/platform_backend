package messaging

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/comment/messaging/dispatch"
)

func JoinCommentProcessor(msg amqp.Delivery) error {
	data := new(generated.Comment)
	err := proto.Unmarshal(msg.Body, data)
	if err != nil {
		return err
	}
	// 传递至责任链
	dispatch.HandleRequest(data, dispatch.Insert)
	return nil
}

func DeleteCommentProcessor(msg amqp.Delivery) error {
	data := new(generated.AfterAuth)
	// 反序列化
	err := proto.Unmarshal(msg.Body, data)
	if err != nil {
		log.Printf("error: DeleteCommentProcessor unmarshaling message: %v", err)
		return fmt.Errorf("deleteCommentProcessor processor error: %w", err)
	}

	// 发送集中处理
	dispatch.HandleRequest(data, dispatch.Delete)
	return nil
}
