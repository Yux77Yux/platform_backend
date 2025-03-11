package api

import (
	"context"

	"github.com/Yux77Yux/platform_backend/generated/common"
	user "github.com/Yux77Yux/platform_backend/generated/user"
	client "github.com/Yux77Yux/platform_backend/scripts/client"
)

func UpdateUserAvatar(ctx context.Context, avatar string, token *common.AccessToken) (*user.UpdateUserAvatarResponse, error) {
	_client, err := client.GetUserClient()
	if err != nil {
		return nil, err
	}
	req := &user.UpdateUserAvatarRequest{
		UserUpdateAvatar: &user.UserUpdateAvatar{},
		AccessToken:      token,
	}
	return _client.UpdateUserAvatar(ctx, req)
}
