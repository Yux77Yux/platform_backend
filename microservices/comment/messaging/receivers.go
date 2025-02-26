package messaging

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/comment/messaging/dispatch"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
)

func JoinCommentProcessor(msg amqp.Delivery) error {
	comment := new(generated.Comment)
	err := proto.Unmarshal(msg.Body, comment)
	if err != nil {
		return err
	}
	// 传递至责任链
	dispatch.HandleRequest(comment, dispatch.Insert)
	return nil
}

func DeleteCommentProcessor(msg amqp.Delivery) error {
	comment := new(common.AfterAuth)
	// 反序列化
	err := proto.Unmarshal(msg.Body, comment)
	if err != nil {
		log.Printf("error: DeleteCommentProcessor unmarshaling message: %v", err)
		return fmt.Errorf("deleteCommentProcessor processor error: %w", err)
	}

	creationId, userId, err := db.GetCreationIdInTransaction(comment.GetCommentId())
	if err != nil {
		return err
	}

	comment.CreationId = creationId
	if comment.GetUserId() != -403 && comment.GetUserId() != userId {
		return nil
	}

	// 发送集中处理
	dispatch.HandleRequest(comment, dispatch.Delete)
	return nil
}
