package internal

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	db "github.com/Yux77Yux/platform_backend/microservices/user/repository"
	errMap "github.com/Yux77Yux/platform_backend/pkg/error"
)

func CancelFollow(ctx context.Context, req *generated.FollowRequest) error {
	follow := req.GetFollow()
	err := db.CancelFollow(ctx, follow)

	if errMap.IsServerError(err) {
		return err
	}

	return nil
}
