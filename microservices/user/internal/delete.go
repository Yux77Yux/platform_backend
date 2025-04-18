package internal

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
	errMap "github.com/Yux77Yux/platform_backend/pkg/error"
)

func CancelFollow(ctx context.Context, req *generated.FollowRequest) error {
	token := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("update", "user", token)
	if err != nil {
		return err
	}
	if !pass {
		return nil
	}
	follow := req.GetFollow()
	follow.FollowerId = userId

	err = db.CancelFollow(ctx, follow)
	if errMap.IsServerError(err) {
		return err
	}

	err = cache.CancelFollow(ctx, follow)
	if err != nil {
		return err
	}

	return nil
}
