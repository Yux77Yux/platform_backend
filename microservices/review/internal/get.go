package internal

import (
	"context"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/review"
	messaging "github.com/Yux77Yux/platform_backend/microservices/review/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/review/repository"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func GetReviews(ctx context.Context, req *generated.GetReviewsRequest) (*generated.GetReviewsResponse, error) {
	token := req.GetAccessToken().GetValue()
	pass, reviewerId, err := auth.Auth("get", "review", token)
	if err != nil {
		return &generated.GetReviewsResponse{
			Msg: &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			},
		}, err
	}
	if !pass {
		return &generated.GetReviewsResponse{
			Msg: &common.ApiResponse{
				Code:   "403",
				Status: common.ApiResponse_ERROR,
			},
		}, nil
	}

	reviews, err := db.GetReviews(ctx, reviewerId, req.GetType(), req.GetStatus(), req.GetPage())
	if err != nil {
		return &generated.GetReviewsResponse{
			Msg: &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			},
		}, err
	}

	return &generated.GetReviewsResponse{
		Reviews: reviews,
		Msg: &common.ApiResponse{
			Code:   "200",
			Status: common.ApiResponse_SUCCESS,
		},
	}, nil
}

func GetNewReviews(ctx context.Context, req *generated.GetNewReviewsRequest) (*generated.GetReviewsResponse, error) {
	token := req.GetAccessToken().GetValue()
	pass, reviewerId, err := auth.Auth("get", "review", token)
	if err != nil {
		return &generated.GetReviewsResponse{
			Msg: &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			},
		}, err
	}
	if !pass {
		return &generated.GetReviewsResponse{
			Msg: &common.ApiResponse{
				Code:   "403",
				Status: common.ApiResponse_ERROR,
			},
		}, nil
	}

	reviews, err := messaging.GetPendingReviews(reviewerId, req.GetType())
	if err != nil {
		return &generated.GetReviewsResponse{
			Msg: &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			},
		}, err
	}

	return &generated.GetReviewsResponse{
		Reviews: reviews,
		Msg: &common.ApiResponse{
			Code:   "200",
			Status: common.ApiResponse_SUCCESS,
		},
	}, nil
}
