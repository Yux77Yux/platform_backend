package messaging

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/review/messaging/dispatch"
)

func NewReviewProcessor(msg amqp.Delivery) error {
	req := new(generated.NewReview)

	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("NewReviewProcessor error: %w", err)
	}

	go dispatch.HandleRequest(req, dispatch.Insert)

	switch req.TargetType {
	case generated.TargetType_COMMENT:
		err = SendProtoMessage(Comment_review, Comment_review, msg.Body)
	case generated.TargetType_USER:
		err = SendProtoMessage(User_review, User_review, msg.Body)
	case generated.TargetType_CREATION:
		err = SendProtoMessage(Creation_review, Creation_review, msg.Body)
	}

	return err
}
