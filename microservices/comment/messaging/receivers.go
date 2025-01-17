package messaging

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
)

func JoinCommentProcessor(msg amqp.Delivery) error {
	// 传递至责任链
	insertChain.HandleRequest(msg.Body)
	return nil
}

func DeleteCommentProcessor(msg amqp.Delivery) error {
	req := new(generated.AfterAuth)
	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("error: DeleteCommentProcessor unmarshaling message: %v", err)
		return fmt.Errorf("deleteCommentProcessor processor error: %w", err)
	}

	// 发送集中处理
	selectListener.Dispatch(msg.Body)

	return nil
}
