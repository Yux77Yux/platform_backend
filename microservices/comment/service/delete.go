package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	internal "github.com/Yux77Yux/platform_backend/microservices/comment/internal"
)

func (s *Server) DeleteComment(ctx context.Context, req *generated.DeleteCommentRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, internal.DeleteComment(ctx, req)
}
