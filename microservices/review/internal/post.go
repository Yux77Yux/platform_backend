package internal

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/review"
	tools "github.com/Yux77Yux/platform_backend/microservices/review/tools"
)

func NewReview(ctx context.Context, req *generated.NewReviewRequest) (*generated.NewReviewResponse, error) {
	review := req.GetNew()

	// 将id发到
	id := tools.GetSnowId()
	if id <= 0 {
		return &generated.NewReviewResponse{
			Msg: &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: "snow id error",
			},
		}, nil
	}
	review.Id = id
	review.CreatedAt = timestamppb.Now()

	go func(review *generated.NewReview, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err := messaging.SendMessage(ctx, EXCHANGE_NEW_REVIEW, KEY_NEW_REVIEW, review)
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(review, ctx)

	return &generated.NewReviewResponse{
		Msg: &common.ApiResponse{
			Code:   "202",
			Status: common.ApiResponse_SUCCESS,
		},
	}, nil
}
