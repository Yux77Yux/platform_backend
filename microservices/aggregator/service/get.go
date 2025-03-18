package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	internal "github.com/Yux77Yux/platform_backend/microservices/aggregator/internal"
)

func (s *Server) Collections(ctx context.Context, req *generated.CollectionsRequest) (*generated.GetCardsResponse, error) {
	return internal.Collections(ctx, req)
}

func (s *Server) History(ctx context.Context, req *generated.HistoryRequest) (*generated.GetCardsResponse, error) {
	return internal.History(ctx, req)
}

func (s *Server) HomePage(ctx context.Context, req *generated.HomeRequest) (*generated.GetCardsResponse, error) {
	return internal.HomePage(ctx, req)
}

func (s *Server) Search(ctx context.Context, req *generated.SearchCreationsRequest) (*generated.SearchCreationsResponse, error) {
	return internal.Search(ctx, req)
}
