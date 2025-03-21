package receiver

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	creation "github.com/Yux77Yux/platform_backend/generated/creation"
	generated "github.com/Yux77Yux/platform_backend/generated/review"
	user "github.com/Yux77Yux/platform_backend/generated/user"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/review/messaging/dispatch"
	tools "github.com/Yux77Yux/platform_backend/microservices/review/tools"
)

func NewReviewProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(generated.NewReview)
	err := msg.UnmarshalTo(req)
	if err != nil {
		return fmt.Errorf("NewReviewProcessor error: %w", err)
	}

	go dispatcher.HandleRequest(req, dispatch.Insert)

	switch req.GetTargetType() {
	case generated.TargetType_COMMENT:
		err = messaging.PreSendMessage(ctx, EXCHANGE_COMMENT_REVIEW, QUEUE_COMMENT_REVIEW, KEY_COMMENT_REVIEW, req)
	case generated.TargetType_USER:
		err = messaging.PreSendMessage(ctx, EXCHANGE_USER_REVIEW, QUEUE_USER_REVIEW, KEY_USER_REVIEW, req)
	case generated.TargetType_CREATION:
		err = messaging.PreSendMessage(ctx, EXCHANGE_CREATION_REVIEW, QUEUE_CREATION_REVIEW, KEY_CREATION_REVIEW, req)
	}

	return err
}

func PendingCreationProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(common.CreationId)

	err := msg.UnmarshalTo(req)
	if err != nil {
		return fmt.Errorf("PendingCreationProcessor error: %w", err)
	}

	creationId := req.GetId()

	id := tools.GetSnowId()
	newReview := &generated.NewReview{
		Id:         id,
		TargetId:   creationId,
		TargetType: generated.TargetType_CREATION,
		CreatedAt:  timestamppb.Now(),
		Msg:        "状态变更",
	}

	err = messaging.SendMessage(ctx, EXCHANGE_NEW_REVIEW, KEY_NEW_REVIEW, newReview)
	return err
}

func BatchUpdateProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(generated.AnyReview)

	err := msg.UnmarshalTo(req)
	if err != nil {
		return fmt.Errorf("BatchUpdateProcessor error: %w", err)
	}
	reviews := req.GetReviews()
	return db.UpdateReviews(ctx, reviews)
}

func UpdateProcessor(ctx context.Context, msg *anypb.Any) error {
	review := new(generated.Review)

	err := msg.UnmarshalTo(review)
	if err != nil {
		return err
	}

	targetId, targetType, err := db.GetTarget(ctx, review.New.GetId())
	if err != nil {
		return err
	}

	review.New.TargetId = targetId
	review.New.TargetType = *targetType

	err = db.UpdateReview(ctx, review)
	if err != nil {
		return err
	}
	var reviewErr error
	switch review.New.GetTargetType() {
	case generated.TargetType_COMMENT:
		reviewErr = UpdateCommentReview(ctx, review)
	case generated.TargetType_CREATION:
		reviewErr = UpdateCreationReview(ctx, review)
	case generated.TargetType_USER:
		reviewErr = UpdateUserReview(ctx, review)
	}

	if reviewErr != nil {
		return reviewErr
	}

	return nil
}

func UpdateCommentReview(ctx context.Context, review *generated.Review) error {
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
		err = messaging.SendMessage(ctx, EXCHANGE_DELETE_COMMENT, KEY_DELETE_COMMENT, commentObj)
	}

	return err
}

func UpdateCreationReview(ctx context.Context, review *generated.Review) error {
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
		err = messaging.SendMessage(ctx, EXCHANGE_UPDATE_CREATION_STATUS, KEY_UPDATE_CREATION_STATUS, creationObj)
	case generated.ReviewStatus_APPROVED:
		creationObj.Status = creation.CreationStatus_PUBLISHED
		err = messaging.SendMessage(ctx, EXCHANGE_UPDATE_CREATION_STATUS, KEY_UPDATE_CREATION_STATUS, creationObj)
	case generated.ReviewStatus_DELETED:
		creationObj.Status = creation.CreationStatus_DELETE
		err = messaging.SendMessage(ctx, EXCHANGE_DELETE_CREATION, KEY_DELETE_CREATION, creationObj)
	}

	return err
}

func UpdateUserReview(ctx context.Context, review *generated.Review) error {
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
		err = messaging.SendMessage(ctx, EXCHANGE_UPDATE_USER_STATUS, KEY_UPDATE_USER_STATUS, updateUser)
	case generated.ReviewStatus_APPROVED:
		updateUser.UserStatus = user.UserStatus_INACTIVE
		err = messaging.SendMessage(ctx, EXCHANGE_UPDATE_USER_STATUS, KEY_UPDATE_USER_STATUS, updateUser)
	}

	return err
}
