package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

func (s *Server) CancelFollow(ctx context.Context, req *generated.FollowRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, internal.CancelFollow(ctx, req)
}
