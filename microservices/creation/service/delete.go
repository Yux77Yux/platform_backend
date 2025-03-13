package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	internal "github.com/Yux77Yux/platform_backend/microservices/creation/internal"
)

func (s *Server) DeleteCreation(ctx context.Context, req *generated.DeleteCreationRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, internal.DeleteCreation(ctx, req)
}
