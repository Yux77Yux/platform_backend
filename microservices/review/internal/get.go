package internal

import (
	"context"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/review"
	errMap "github.com/Yux77Yux/platform_backend/pkg/error"
)

func GetReviews(ctx context.Context, req *generated.GetReviewsRequest) (*generated.GetReviewsResponse, error) {
	response := new(generated.GetReviewsResponse)

	reviewerId := req.GetReviewerId()
	reviews, count, err := db.GetReviews(ctx, reviewerId, req.GetType(), req.GetStatus(), req.GetPage())
	if err != nil {
		if errMap.IsServerError(err) {
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    errMap.GrpcCodeToHTTPStatusString(err),
				Details: err.Error(),
			}
			return response, err
		}
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    errMap.GrpcCodeToHTTPStatusString(err),
			Details: err.Error(),
		}
		return response, nil
	}

	response.Count = count
	response.Reviews = reviews
	response.Msg = &common.ApiResponse{
		Code:   "200",
		Status: common.ApiResponse_SUCCESS,
	}
	return response, nil
}

func GetNewReviews(ctx context.Context, req *generated.GetNewReviewsRequest) (*generated.GetReviewsResponse, error) {
	response := new(generated.GetReviewsResponse)

	reviewerId := req.GetReviewerId()
	reviews, err := messaging.GetPendingReviews(ctx, reviewerId, req.GetType())
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}

	response.Reviews = reviews
	response.Msg = &common.ApiResponse{
		Code:   "200",
		Status: common.ApiResponse_SUCCESS,
	}
	return response, nil
}
