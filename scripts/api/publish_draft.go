package api

import (
	"context"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	creation "github.com/Yux77Yux/platform_backend/generated/creation"
	client "github.com/Yux77Yux/platform_backend/scripts/client"
)

func PublishDraftCreation(ctx context.Context, token *common.AccessToken, creationId int64) (*creation.UpdateCreationResponse, error) {
	_client, err := client.GetCreationClient()
	if err != nil {
		return nil, err
	}

	req := &creation.UpdateCreationStatusRequest{
		AccessToken: token,
		UpdateInfo: &creation.CreationUpdateStatus{
			Status:     creation.CreationStatus_PENDING,
			CreationId: creationId,
		},
	}
	return _client.PublishDraftCreation(ctx, req)
}
