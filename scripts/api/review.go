package api

import (
	"context"

	"github.com/Yux77Yux/platform_backend/generated/common"
	review "github.com/Yux77Yux/platform_backend/generated/review"
	client "github.com/Yux77Yux/platform_backend/scripts/client"
)

func GetNewReviews(ctx context.Context, id int64, Type review.TargetType) (*review.GetReviewsResponse, error) {
	_client, err := client.GetReviewClient()
	if err != nil {
		return nil, err
	}
	req := &review.GetNewReviewsRequest{
		ReviewerId: id,
		Type:       Type,
	}
	return _client.GetNewReviews(ctx, req)
}

func GetReviews(ctx context.Context, req *review.GetReviewsRequest) (*review.GetReviewsResponse, error) {
	_client, err := client.GetReviewClient()
	if err != nil {
		return nil, err
	}
	return _client.GetReviews(ctx, req)
}

func UpdateReview(ctx context.Context, token *common.AccessToken, _review *review.Review) (*review.UpdateReviewResponse, error) {
	_client, err := client.GetReviewClient()
	if err != nil {
		return nil, err
	}

	req := &review.UpdateReviewRequest{
		AccessToken: token,
		Review:      _review,
	}

	return _client.UpdateReview(ctx, req)
}
