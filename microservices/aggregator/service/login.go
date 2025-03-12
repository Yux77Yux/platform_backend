package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	internal "github.com/Yux77Yux/platform_backend/microservices/aggregator/internal"
)

func (s *Server) Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {
	return internal.Login(ctx, req)
}
