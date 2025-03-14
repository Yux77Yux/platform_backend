package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
	internal "github.com/Yux77Yux/platform_backend/microservices/review/internal"
)

func (s *Server) UpdateReview(ctx context.Context, req *generated.UpdateReviewRequest) (*generated.UpdateReviewResponse, error) {
	return internal.UpdateReview(ctx, req)
}
