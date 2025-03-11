package api

import (
	"context"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	creation "github.com/Yux77Yux/platform_backend/generated/creation"
	client "github.com/Yux77Yux/platform_backend/scripts/client"
)

func GetUserCreations(ctx context.Context, token *common.AccessToken, status creation.CreationStatus, page int32) (*creation.GetCreationListResponse, error) {
	_client, err := client.GetCreationClient()
	if err != nil {
		return nil, err
	}

	req := &creation.GetUserCreationsRequest{
		AccessToken: token,
		Page:        page,
		Status:      status,
	}
	return _client.GetUserCreations(ctx, req)
}

func GetCreationList(ctx context.Context, creationIds []int64) (*creation.GetCreationListResponse, error) {
	_client, err := client.GetCreationClient()
	if err != nil {
		return nil, err
	}

	req := &creation.GetCreationListRequest{
		Ids: creationIds,
	}
	return _client.GetCreationList(ctx, req)
}
