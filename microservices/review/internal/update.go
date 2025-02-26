package internal

import (
	common "github.com/Yux77Yux/platform_backend/generated/common"
	creation "github.com/Yux77Yux/platform_backend/generated/creation"
	generated "github.com/Yux77Yux/platform_backend/generated/review"
	user "github.com/Yux77Yux/platform_backend/generated/user"
	messaging "github.com/Yux77Yux/platform_backend/microservices/review/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/review/repository"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func UpdateReview(req *generated.UpdateReviewRequest) (*generated.UpdateReviewResponse, error) {
	response := new(generated.UpdateReviewResponse)
	token := req.GetAccessToken().GetValue()
	pass, reviewerId, err := auth.Auth("update", "review", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Code:   "403",
			Status: common.ApiResponse_ERROR,
		}
		return response, err
	}

	review := req.GetReview()
	review.ReviewerId = reviewerId

	err = db.UpdateReview(review)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}

	// 执行完事务后发送跨服务事件
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
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: reviewErr.Error(),
		}
	}

	return &generated.UpdateReviewResponse{
		Msg: &common.ApiResponse{
			Code:   "202",
			Status: common.ApiResponse_SUCCESS,
		},
	}, nil
}

func UpdateCommentReview(review *generated.Review) error {
	status := review.GetStatus()
	commentId := review.GetNew().GetId()

	commentObj := &common.AfterAuth{
		CommentId: int32(commentId),
		UserId:    -403,
	}

	var err error
	switch status {
	case generated.ReviewStatus_REJECTED, generated.ReviewStatus_DELETED:
		err = messaging.SendMessage(messaging.COMMENT_REJECTED, messaging.COMMENT_REJECTED, commentObj)
	}

	return err
}

func UpdateCreationReview(review *generated.Review) error {
	status := review.GetStatus()
	creationId := review.GetNew().GetId()

	creationObj := &creation.CreationUpdateStatus{
		CreationId: creationId,
		AuthorId:   -403,
	}

	var err error
	switch status {
	case generated.ReviewStatus_REJECTED:
		creationObj.Status = creation.CreationStatus_REJECTED
		err = messaging.SendMessage(messaging.CREATION_REJECTED, messaging.CREATION_REJECTED, creationObj)
	case generated.ReviewStatus_APPROVED:
		creationObj.Status = creation.CreationStatus_PUBLISHED
		err = messaging.SendMessage(messaging.CREATION_APPROVE, messaging.CREATION_APPROVE, creationObj)
	case generated.ReviewStatus_DELETED:
		creationObj.Status = creation.CreationStatus_DELETE
		err = messaging.SendMessage(messaging.CREATION_DELETED, messaging.CREATION_DELETED, creationObj)
	}

	return err
}

func UpdateUserReview(review *generated.Review) error {
	status := review.GetStatus()
	userId := review.GetNew().GetId()

	updateUser := &user.UserUpdateStatus{
		UserId: userId,
	}

	var err error
	switch status {
	case generated.ReviewStatus_REJECTED:
		updateUser.UserStatus = user.UserStatus_LIMITED
		err = messaging.SendMessage(messaging.USER_REJECTED, messaging.USER_REJECTED, updateUser)
	case generated.ReviewStatus_APPROVED:
		updateUser.UserStatus = user.UserStatus_INACTIVE
		err = messaging.SendMessage(messaging.USER_APPROVE, messaging.USER_APPROVE, updateUser)
	case generated.ReviewStatus_DELETED:
		updateUser.UserStatus = user.UserStatus_DELETE
		err = messaging.SendMessage(messaging.USER_DELETED, messaging.USER_DELETED, updateUser)
	}

	return err
}
