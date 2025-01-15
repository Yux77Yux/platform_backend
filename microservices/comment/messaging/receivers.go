package messaging

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	cache "github.com/Yux77Yux/platform_backend/microservices/comment/cache"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
)

func JoinCommentProcessor(msg amqp.Delivery) error {
	comment := new(generated.Comment)
	// 反序列化
	err := proto.Unmarshal(msg.Body, comment)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("joinCommentProcessor processor error: %w", err)
	}

	// 传递至责任链
	chain.HandleCommentRequest(comment)

	return nil
}
