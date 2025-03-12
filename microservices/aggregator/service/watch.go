package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	internal "github.com/Yux77Yux/platform_backend/microservices/aggregator/internal"
)

func (s *Server) WatchCreation(ctx context.Context, req *generated.WatchCreationRequest) (*generated.WatchCreationResponse, error) {
	return internal.WatchCreation(ctx, req)
}

func (s *Server) SimilarCreations(ctx context.Context, req *generated.SimilarCreationsRequest) (*generated.GetCardsResponse, error) {
	return internal.SimilarCreations(ctx, req)
}

func (s *Server) InitialComments(ctx context.Context, req *generated.InitialCommentsRequest) (*generated.InitialCommentsResponse, error) {
	return internal.InitialComments(ctx, req)
}

func (s *Server) GetTopComments(ctx context.Context, req *generated.GetTopCommentsRequest) (*generated.GetTopCommentsResponse, error) {
	return internal.GetTopComments(ctx, req)
}

func (s *Server) GetSecondComments(ctx context.Context, req *generated.GetSecondCommentsRequest) (*generated.GetSecondCommentsResponse, error) {
	return internal.GetSecondComments(ctx, req)
}
