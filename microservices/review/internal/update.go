package internal

import (
	"github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/review"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/review/messaging/dispatch"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func UpdateReview(req *generated.UpdateReviewRequest) (*generated.UpdateReviewResponse, error) {
	token := req.GetAccessToken().GetValue()
	pass, reviewerId, err := auth.Auth("update", "review", token)
	if err != nil {
		return &generated.UpdateReviewResponse{
			Msg: &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			},
		}, err
	}
	if !pass {
		return &generated.UpdateReviewResponse{
			Msg: &common.ApiResponse{
				Code:   "403",
				Status: common.ApiResponse_ERROR,
			},
		}, nil
	}

	review := req.GetReview()
	review.ReviewerId = reviewerId

	dispatch.HandleRequest(review, dispatch.Update)
	return &generated.UpdateReviewResponse{
		Msg: &common.ApiResponse{
			Code:   "202",
			Status: common.ApiResponse_SUCCESS,
		},
	}, nil
}
