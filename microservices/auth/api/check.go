package api

import (
	"context"

	generatedAuth "github.com/Yux77Yux/platform_backend/generated/auth"
)

func (s *Server) Check(ctx context.Context, req *generatedAuth.CheckRequest) (*generatedAuth.CheckResponse, error) {
	var (
		res *generatedAuth.CheckResponse
		err error
	)

	return res, err
}
