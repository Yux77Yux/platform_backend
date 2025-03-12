package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/auth"
	internal "github.com/Yux77Yux/platform_backend/microservices/auth/internal"
)

func (s *Server) Refresh(ctx context.Context, req *generated.RefreshRequest) (*generated.RefreshResponse, error) {
	return internal.Refresh(req)
}
