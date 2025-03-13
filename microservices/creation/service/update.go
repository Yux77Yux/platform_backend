package service

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	internal "github.com/Yux77Yux/platform_backend/microservices/creation/internal"
)

func (s *Server) UpdateCreation(ctx context.Context, req *generated.UpdateCreationRequest) (*generated.UpdateCreationResponse, error) {
	return internal.UpdateCreation(ctx, req)
}

func (s *Server) PublishDraftCreation(ctx context.Context, req *generated.UpdateCreationStatusRequest) (*generated.UpdateCreationResponse, error) {
	info := req.GetUpdateInfo()
	if info == nil {
		info = new(generated.CreationUpdateStatus)
	}
	info.Status = generated.CreationStatus_PENDING
	req.UpdateInfo = info
	return internal.UpdateCreationStatus(ctx, req)
}
