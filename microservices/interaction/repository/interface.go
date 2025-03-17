package repository

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	pkgDb "github.com/Yux77Yux/platform_backend/pkg/database"
)

type SqlInterface = pkgDb.SqlInterface

type SqlMethod interface {
	GetActionTag(ctx context.Context, req *generated.BaseInteraction) (*generated.Interaction, error)
	GetCollections(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error)
	GetHistories(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error)
	GetOtherUserHistories(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error)
	UpdateInteractions(ctx context.Context, req []*generated.OperateInteraction) error
}
