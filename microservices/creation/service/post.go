package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	internal "github.com/Yux77Yux/platform_backend/microservices/creation/internal"
)

func (s *Server) UploadCreation(ctx context.Context, req *generated.UploadCreationRequest) (*generated.UploadCreationResponse, error) {
	return internal.UploadCreation(req)
}
