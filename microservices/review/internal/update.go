package internal

import (
	"context"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/review"
	tools "github.com/Yux77Yux/platform_backend/microservices/review/tools"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func UpdateReview(ctx context.Context, req *generated.UpdateReviewRequest) (*generated.UpdateReviewResponse, error) {
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

	go func(review *generated.Review, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err = messaging.SendMessage(ctx, EXCHANGE_UPDATE, KEY_UPDATE, review)
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(review, ctx)

	return &generated.UpdateReviewResponse{
		Msg: &common.ApiResponse{
			Code:   "202",
			Status: common.ApiResponse_PENDING,
		},
	}, nil
}
