package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	internal "github.com/Yux77Yux/platform_backend/microservices/comment/internal"
)

func (s *Server) GetComments(ctx context.Context, req *generated.GetCommentsRequest) (*generated.GetCommentsResponse, error) {
	return internal.GetComments(ctx, req)
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

func (s *Server) GetReplyComments(ctx context.Context, req *generated.GetReplyCommentsRequest) (*generated.GetCommentsResponse, error) {
	return internal.GetReplyComments(ctx, req)
}
