package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	internal "github.com/Yux77Yux/platform_backend/microservices/interaction/internal"
)

func (s *Server) GetActionTag(ctx context.Context, req *generated.GetCreationInteractionRequest) (*generated.GetCreationInteractionResponse, error) {
	return internal.GetActionTag(ctx, req)
}

func (s *Server) GetCollections(ctx context.Context, req *generated.GetCollectionsRequest) (*generated.GetInteractionsResponse, error) {
	return internal.GetCollections(ctx, req)
}

func (s *Server) GetHistories(ctx context.Context, req *generated.GetHistoriesRequest) (*generated.GetInteractionsResponse, error) {
	return internal.GetHistories(ctx, req)
}

func (s *Server) GetRecommendBaseUser(ctx context.Context, req *generated.GetRecommendRequest) (*generated.GetRecommendResponse, error) {
	return internal.GetRecommendBaseUser(ctx, req)
}

func (s *Server) GetRecommendBaseCreation(ctx context.Context, req *generated.GetRecommendRequest) (*generated.GetRecommendResponse, error) {
	return internal.GetRecommendBaseCreation(ctx, req)
}
