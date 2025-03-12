package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/auth"
	internal "github.com/Yux77Yux/platform_backend/microservices/auth/internal"
)

func (s *Server) Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {
	return internal.Login(req)
}
