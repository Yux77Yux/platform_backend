package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	internal "github.com/Yux77Yux/platform_backend/microservices/aggregator/internal"
)

func (s *Server) GetUserReviews(ctx context.Context, req *generated.GetReviewsRequest) (*generated.GetUserReviewsResponse, error) {
	return internal.GetUserReviews(ctx, req)
}

func (s *Server) GetCreationReviews(ctx context.Context, req *generated.GetReviewsRequest) (*generated.GetCreationReviewsResponse, error) {
	return internal.GetCreationReviews(ctx, req)
}

func (s *Server) GetCommentReviews(ctx context.Context, req *generated.GetReviewsRequest) (*generated.GetCommentReviewsResponse, error) {
	return internal.GetCommentReviews(ctx, req)
}

func (s *Server) GetNewUserReviews(ctx context.Context, req *generated.GetNewReviewsRequest) (*generated.GetUserReviewsResponse, error) {
	return internal.GetNewUserReviews(ctx, req)
}

func (s *Server) GetNewCreationReviews(ctx context.Context, req *generated.GetNewReviewsRequest) (*generated.GetCreationReviewsResponse, error) {
	return internal.GetNewCreationReviews(ctx, req)
}

func (s *Server) GetNewCommentReviews(ctx context.Context, req *generated.GetNewReviewsRequest) (*generated.GetCommentReviewsResponse, error) {
	return internal.GetNewCommentReviews(ctx, req)
}
