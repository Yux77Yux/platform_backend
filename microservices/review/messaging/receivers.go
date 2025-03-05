package messaging

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	creation "github.com/Yux77Yux/platform_backend/generated/creation"
	generated "github.com/Yux77Yux/platform_backend/generated/review"
	user "github.com/Yux77Yux/platform_backend/generated/user"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/review/messaging/dispatch"
	db "github.com/Yux77Yux/platform_backend/microservices/review/repository"
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
		err = PreSendProtoMessage(Comment_review, Comment_review, Comment_review, body)
	case generated.TargetType_USER:
		err = PreSendProtoMessage(User_review, User_review, User_review, body)
	case generated.TargetType_CREATION:
		err = PreSendProtoMessage(Creation_review, Creation_review, Creation_review, body)
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
		Msg:        "状态变更",
	}

	err = SendMessage(New_review, New_review, newReview)
	return err
}

func BatchUpdateProcessor(msg amqp.Delivery) error {
	req := new(generated.AnyReview)

	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("NewReviewProcessor error: %w", err)
	}

	go dispatch.HandleRequest(req, dispatch.BatchUpdate)
	return nil
}

func UpdateProcessor(msg amqp.Delivery) error {
	review := new(generated.Review)

	// 反序列化
	err := proto.Unmarshal(msg.Body, review)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("NewReviewProcessor error: %w", err)
	}

	err = db.UpdateReview(review)
	if err != nil {
		log.Printf("error: db UpdateReview %v", err)
		return err
	}

	var reviewErr error
	switch review.New.GetTargetType() {
	case generated.TargetType_COMMENT:
		reviewErr = UpdateCommentReview(review)
	case generated.TargetType_CREATION:
		reviewErr = UpdateCreationReview(review)
	case generated.TargetType_USER:
		reviewErr = UpdateUserReview(review)
	}

	if reviewErr != nil {
		return reviewErr
	}

	return nil
}

func UpdateCommentReview(review *generated.Review) error {
	status := review.GetStatus()
	commentId := review.GetNew().GetTargetId()
	if commentId <= 0 {
		return fmt.Errorf("error: targetId is null")
	}

	commentObj := &common.AfterAuth{
		CommentId: int32(commentId),
		UserId:    -403,
	}

	var err error
	switch status {
	case generated.ReviewStatus_REJECTED, generated.ReviewStatus_DELETED:
		err = SendMessage(COMMENT_REJECTED, COMMENT_REJECTED, commentObj)
	}

	return err
}

func UpdateCreationReview(review *generated.Review) error {
	status := review.GetStatus()
	creationId := review.GetNew().GetTargetId()
	if creationId <= 0 {
		return fmt.Errorf("error: targetId is null")
	}

	creationObj := &creation.CreationUpdateStatus{
		CreationId: creationId,
		AuthorId:   -403,
	}

	var err error
	switch status {
	case generated.ReviewStatus_REJECTED:
		creationObj.Status = creation.CreationStatus_REJECTED
		err = SendMessage(CREATION_REJECTED, CREATION_REJECTED, creationObj)
	case generated.ReviewStatus_APPROVED:
		creationObj.Status = creation.CreationStatus_PUBLISHED
		err = SendMessage(CREATION_APPROVE, CREATION_APPROVE, creationObj)
	case generated.ReviewStatus_DELETED:
		creationObj.Status = creation.CreationStatus_DELETE
		err = SendMessage(CREATION_DELETED, CREATION_DELETED, creationObj)
	}

	return err
}

func UpdateUserReview(review *generated.Review) error {
	status := review.GetStatus()
	userId := review.GetNew().GetTargetId()
	if userId <= 0 {
		return fmt.Errorf("error: targetId is null")
	}

	updateUser := &user.UserUpdateStatus{
		UserId: userId,
	}

	var err error
	switch status {
	case generated.ReviewStatus_REJECTED:
		updateUser.UserStatus = user.UserStatus_LIMITED
		err = SendMessage(USER_REJECTED, USER_REJECTED, updateUser)
	case generated.ReviewStatus_APPROVED:
		updateUser.UserStatus = user.UserStatus_INACTIVE
		err = SendMessage(USER_APPROVE, USER_APPROVE, updateUser)
	}

	return err
}
