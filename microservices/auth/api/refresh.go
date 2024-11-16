package api

import (
	"context"

	generatedAuth "github.com/Yux77Yux/platform_backend/generated/auth"
)

func (s *Server) Refresh(ctx context.Context, req *generatedAuth.RefreshRequest) (*generatedAuth.RefreshResponse, error) {
	var (
		res *generatedAuth.RefreshResponse
		err error
	)

	return res, err
}
