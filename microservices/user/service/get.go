package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

func (s *Server) Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {
	return internal.Login(ctx, req)
}

func (s *Server) GetUser(ctx context.Context, req *generated.GetUserRequest) (*generated.GetUserResponse, error) {
	return internal.GetUser(ctx, req)
}

func (s *Server) GetUsers(ctx context.Context, req *generated.GetUsersRequest) (*generated.GetUsersResponse, error) {
	return internal.GetUsers(ctx, req)
}

func (s *Server) GetFolloweesByTime(ctx context.Context, req *generated.GetFollowRequest) (*generated.GetFollowResponse, error) {
	return internal.GetFolloweesByTime(ctx, req)
}

func (s *Server) GetFolloweesByViews(ctx context.Context, req *generated.GetFollowRequest) (*generated.GetFollowResponse, error) {
	return internal.GetFolloweesByViews(ctx, req)
}

func (s *Server) GetFollowers(ctx context.Context, req *generated.GetFollowRequest) (*generated.GetFollowResponse, error) {
	return internal.GetFollowers(ctx, req)
}
