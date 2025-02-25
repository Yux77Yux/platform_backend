package internal

import (
	"context"
	"strconv"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/review"
	messaging "github.com/Yux77Yux/platform_backend/microservices/review/messaging"
	snow "github.com/Yux77Yux/platform_backend/pkg/snow"
)

func NewReview(ctx context.Context, req *generated.NewReviewRequest) (*generated.NewReviewResponse, error) {
	review := req.GetNew()

	// 将id发到
	id := snow.GetId()
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

	err := messaging.SendMessage(messaging.New_review, messaging.New_review, review)
	if err != nil {
		return &generated.NewReviewResponse{
			Msg: &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			},
		}, nil
	}

	return &generated.NewReviewResponse{
		Msg: &common.ApiResponse{
			Code:    "202",
			Status:  common.ApiResponse_SUCCESS,
			TraceId: strconv.FormatInt(id, 10),
		},
	}, nil
}
