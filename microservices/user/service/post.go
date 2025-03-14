package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

func (s *Server) AddReviewer(ctx context.Context, req *generated.AddReviewerRequest) (*generated.AddReviewerResponse, error) {
	return internal.AddReviewer(ctx, req)
}

func (s *Server) Follow(ctx context.Context, req *generated.FollowRequest) (*generated.FollowResponse, error) {
	return internal.Follow(ctx, req)
}

func (s *Server) Register(ctx context.Context, req *generated.RegisterRequest) (*generated.RegisterResponse, error) {
	return internal.Register(ctx, req)
}
