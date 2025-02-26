package messaging

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/review"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/review/messaging/dispatch"
	snow "github.com/Yux77Yux/platform_backend/pkg/snow"
)

func NewReviewProcessor(msg amqp.Delivery) error {
	req := new(generated.NewReview)

	body := msg.Body
	// 反序列化
	err := proto.Unmarshal(body, req)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("NewReviewProcessor error: %w", err)
	}

	go dispatch.HandleRequest(req, dispatch.Insert)

	switch req.GetTargetType() {
	case generated.TargetType_COMMENT:
		err = SendProtoMessage(Comment_review, Comment_review, body)
	case generated.TargetType_USER:
		err = SendProtoMessage(User_review, User_review, body)
	case generated.TargetType_CREATION:
		err = SendProtoMessage(Creation_review, Creation_review, body)
	}

	return err
}

func PendingCreationProcessor(msg amqp.Delivery) error {
	req := new(common.CreationId)

	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("NewReviewProcessor error: %w", err)
	}

	creationId := req.GetId()

	id := snow.GetId()
	newReview := &generated.NewReview{
		Id:         id,
		TargetId:   creationId,
		TargetType: generated.TargetType_CREATION,
	}

	err = SendMessage(New_review, New_review, newReview)
	return err
}
