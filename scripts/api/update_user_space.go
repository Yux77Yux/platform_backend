package api

import (
	"context"

	"github.com/Yux77Yux/platform_backend/generated/common"
	user "github.com/Yux77Yux/platform_backend/generated/user"
	client "github.com/Yux77Yux/platform_backend/scripts/client"
)

func UpdateUserSpace(ctx context.Context, name, bio string, token *common.AccessToken) (*user.UpdateUserResponse, error) {
	_client, err := client.GetUserClient()
	if err != nil {
		return nil, err
	}
	req := &user.UpdateUserSpaceRequest{
		UserUpdateSpace: &user.UserUpdateSpace{
			UserDefault: &common.UserDefault{
				UserName: name,
			},
			UserBio: bio,
		},
		AccessToken: token,
	}
	return _client.UpdateUserSpace(ctx, req)
}
