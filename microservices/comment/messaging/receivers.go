package messaging

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
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

func DeleteCommentProcessor(msg amqp.Delivery) error {
	req := new(generated.DeleteCommentRequest)
	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("DeleteCommentProcessor processor error: %w", err)
	}

	accessToken := req.GetAccessToken().GetValue()
	// 取作品信息，鉴权
	pass, user_id, err := auth.Auth("delete", "creation", accessToken)
	if err != nil {
		return err
	}
	if !pass {
		return fmt.Errorf("no pass")
	}
	// 以上为鉴权

	comment_id := req.GetCommentId()

	// 取发布者id
	var author_id int64 = -1
	author_id, err = db.GetPublisherIdInTransaction(comment_id)
	if err != nil {
		return err
	}
	if author_id != user_id {
		// 评论发布者与token中ID不一致
		return fmt.Errorf("error: author %v not match the token", author_id)
	}

	return nil
}
