package api

import (
	"context"

	user "github.com/Yux77Yux/platform_backend/generated/user"
	client "github.com/Yux77Yux/platform_backend/scripts/client"
)

func Register(ctx context.Context, testId string) (*user.RegisterResponse, error) {
	_client, err := client.GetUserClient()
	if err != nil {
		return nil, err
	}
	req := &user.RegisterRequest{
		UserCredentials: &user.UserCredentials{
			Username: testId,
			Password: testId,
		},
	}
	return _client.Register(ctx, req)
}
