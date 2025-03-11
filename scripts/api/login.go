package api

import (
	"context"

	aggregator "github.com/Yux77Yux/platform_backend/generated/aggregator"
	user "github.com/Yux77Yux/platform_backend/generated/user"
	client "github.com/Yux77Yux/platform_backend/scripts/client"
)

func Login(ctx context.Context, testId string) (*aggregator.LoginResponse, error) {
	_client, err := client.GetAggregatorClient()
	if err != nil {
		return nil, err
	}
	req := &aggregator.LoginRequest{
		UserCredentials: &user.UserCredentials{
			Username: testId,
			Password: testId,
		},
	}
	return _client.Login(ctx, req)
}
