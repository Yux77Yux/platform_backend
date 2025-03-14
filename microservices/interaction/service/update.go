package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	internal "github.com/Yux77Yux/platform_backend/microservices/interaction/internal"
)

func (s *Server) CancelCollections(ctx context.Context, req *generated.UpdateInteractionsRequest) (*generated.UpdateInteractionResponse, error) {
	return internal.CancelCollections(ctx, req)
}

func (s *Server) CancelLike(ctx context.Context, req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
	return internal.CancelLike(ctx, req)
}

func (s *Server) ClickCollection(ctx context.Context, req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
	return internal.ClickCollection(ctx, req)
}

func (s *Server) ClickLike(ctx context.Context, req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
	return internal.ClickLike(ctx, req)
}

func (s *Server) DelHistories(ctx context.Context, req *generated.UpdateInteractionsRequest) (*generated.UpdateInteractionResponse, error) {
	return internal.DelHistories(ctx, req)
}
