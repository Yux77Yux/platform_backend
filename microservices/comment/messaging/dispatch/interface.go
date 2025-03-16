package dispatch

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	pkgDispatch "github.com/Yux77Yux/platform_backend/pkg/dispatch"
)

type CacheMethod interface {
	UpdateCommentsCount(ctx context.Context, creationId int64, count int64) error
}

type SqlMethod interface {
	BatchInsert(ctx context.Context, comments []*generated.Comment) (int64, error)
	BatchUpdateDeleteStatus(ctx context.Context, comments []*common.AfterAuth) (int64, error)
}

type DispatchInterface = pkgDispatch.DispatchInterface

type ChainInterface = pkgDispatch.ChainInterface

type ListenerInterface = pkgDispatch.ListenerInterface
