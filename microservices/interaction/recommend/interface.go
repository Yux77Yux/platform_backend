package recommend

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
)

type CacheMethod interface {
	GetHistories(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error)
	GetUsers(ctx context.Context, creationId int64) ([]int64, error)
	ScanZSetsByCreationId(ctx context.Context) ([]string, error)
	ScanZSetsByHistories(ctx context.Context) ([]string, error)
	GetAllItemUsers(ctx context.Context, idStrs []string) (map[int64]map[int64]float64, error)
	GetAllInteractions(ctx context.Context, idStrs []string) (map[int64]map[int64]float64, error)
}
