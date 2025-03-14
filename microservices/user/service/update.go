package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

func (s *Server) DelReviewer(ctx context.Context, req *generated.DelReviewerRequest) (*generated.DelReviewerResponse, error) {
	return internal.DelReviewer(ctx, req)
}

func (s *Server) UpdateUserSpace(ctx context.Context, req *generated.UpdateUserSpaceRequest) (*generated.UpdateUserResponse, error) {
	return internal.UpdateUserSpace(ctx, req)
}

func (s *Server) UpdateUserAvatar(ctx context.Context, req *generated.UpdateUserAvatarRequest) (*generated.UpdateUserAvatarResponse, error) {
	return internal.UpdateUserAvatar(ctx, req)
}

func (s *Server) UpdateUserStatus(ctx context.Context, req *generated.UpdateUserStatusRequest) (*generated.UpdateUserResponse, error) {
	return internal.UpdateUserStatus(ctx, req)
}

func (s *Server) UpdateUserBio(ctx context.Context, req *generated.UpdateUserBioRequest) (*generated.UpdateUserResponse, error) {
	return internal.UpdateUserBio(ctx, req)
}
