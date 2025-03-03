package internal

import (
	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/review"
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

	targetId, err := db.GetTargetId(review.New.GetId())
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}

	review.ReviewerId = reviewerId
	review.New.TargetId = targetId
	err = messaging.SendMessage(messaging.Update, messaging.Update, review)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}

	return &generated.UpdateReviewResponse{
		Msg: &common.ApiResponse{
			Code:   "202",
			Status: common.ApiResponse_PENDING,
		},
	}, nil
}
