package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
	internal "github.com/Yux77Yux/platform_backend/microservices/review/internal"
)

func (s *Server) NewReview(ctx context.Context, req *generated.NewReviewRequest) (*generated.NewReviewResponse, error) {
	return internal.NewReview(ctx, req)
}
