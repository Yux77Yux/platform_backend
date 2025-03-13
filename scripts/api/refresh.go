package api

import (
	"context"

	auth "github.com/Yux77Yux/platform_backend/generated/auth"
	client "github.com/Yux77Yux/platform_backend/scripts/client"
)

func Refresh(ctx context.Context, token *auth.RefreshToken) (*auth.RefreshResponse, error) {
	_client, err := client.GetAuthClient()
	if err != nil {
		return nil, err
	}
	req := &auth.RefreshRequest{
		RefreshToken: token,
	}
	return _client.Refresh(ctx, req)
}
