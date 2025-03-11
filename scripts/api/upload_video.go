package api

import (
	"context"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	creation "github.com/Yux77Yux/platform_backend/generated/creation"
	client "github.com/Yux77Yux/platform_backend/scripts/client"
	data "github.com/Yux77Yux/platform_backend/scripts/data"
)

func ToCreationUpload(newInfo *data.Creation) *creation.CreationUpload {
	return &creation.CreationUpload{
		Src:        newInfo.Src,
		Title:      newInfo.Title,
		Bio:        newInfo.Bio,
		Duration:   newInfo.Duration,
		CategoryId: newInfo.CategoryId,
		Thumbnail:  newInfo.Thumbnail,
	}
}

func UploadCreation(ctx context.Context, token *common.AccessToken, newInfo *data.Creation) (*creation.UploadCreationResponse, error) {
	_client, err := client.GetCreationClient()
	if err != nil {
		return nil, err
	}

	baseInfo := ToCreationUpload(newInfo)
	baseInfo.Status = creation.CreationStatus_DRAFT
	req := &creation.UploadCreationRequest{
		AccessToken: token,
		BaseInfo:    baseInfo,
	}
	return _client.UploadCreation(ctx, req)
}
