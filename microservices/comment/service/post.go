package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	internal "github.com/Yux77Yux/platform_backend/microservices/comment/internal"
)

func (s *Server) PublishComment(ctx context.Context, req *generated.PublishCommentRequest) (*generated.PublishCommentResponse, error) {
	return internal.PublishComment(req)
}
