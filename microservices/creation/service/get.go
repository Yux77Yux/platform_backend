package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	internal "github.com/Yux77Yux/platform_backend/microservices/creation/internal"
)

func (s *Server) GetCreation(ctx context.Context, req *generated.GetCreationRequest) (*generated.GetCreationResponse, error) {
	return internal.GetCreation(ctx, req)
}

func (s *Server) GetCreationPrivate(ctx context.Context, req *generated.GetCreationPrivateRequest) (*generated.GetCreationResponse, error) {
	return internal.GetCreationPrivate(ctx, req)
}

func (s *Server) GetCreationList(ctx context.Context, req *generated.GetCreationListRequest) (*generated.GetCreationListResponse, error) {
	return internal.GetCreationList(ctx, req)
}

func (s *Server) GetUserCreations(ctx context.Context, req *generated.GetUserCreationsRequest) (*generated.GetCreationListResponse, error) {
	return internal.GetUserCreations(ctx, req)
}

func (s *Server) GetPublicCreationList(ctx context.Context, req *generated.GetCreationListRequest) (*generated.GetCreationListResponse, error) {
	return internal.GetPublicCreationList(ctx, req)
}

func (s *Server) GetSpaceCreations(ctx context.Context, req *generated.GetSpaceCreationsRequest) (*generated.GetCreationListResponse, error) {
	return internal.GetSpaceCreations(ctx, req)
}
func (s *Server) SearchCreation(ctx context.Context, req *generated.SearchCreationRequest) (*generated.GetCreationListResponse, error) {
	return internal.SearchCreation(ctx, req)
}
