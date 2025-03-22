package cache

import (
	"context"

	"github.com/go-redis/redis/v8"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
)

type CacheInterface interface {
	Open(connStr string, password string) error
	Close() error

	ScanSet(ctx context.Context, kind string, fliter string, cursor uint64, count int64) ([]string, uint64, error)
	AddToSet(ctx context.Context, kind string, unique string, values ...interface{}) error
	GetPopNSet(ctx context.Context, kind string, unique string, count int64) ([]string, error)
	RemSet(ctx context.Context, kind string, unique string, values ...interface{}) error
	ExistsInSet(ctx context.Context, kind string, unique string, value interface{}) (bool, error)
	CountSet(ctx context.Context, kind string, unique string) (int64, error)
	GetMembersSet(ctx context.Context, kind string, unique string) ([]string, error)

	GetRandZSetMember(ctx context.Context, kind string, unique string, count int) ([]string, error)
	ScanZSet(ctx context.Context, kind string, fliter string, cursor uint64, count int64) ([]string, uint64, error)
	AddZSet(ctx context.Context, kind string, unique string, member string, score float64) error
	ModifyScoreZSet(ctx context.Context, kind string, unique string, member string, score float64) error
	ZRemMemberZSet(ctx context.Context, kind string, unique string, members ...interface{}) error
	GetRankZSet(ctx context.Context, kind string, unique string, member string) (int64, error)
	GetScoreZSet(ctx context.Context, kind string, unique string, member string) (float64, error)
	RangeZSet(ctx context.Context, kind string, unique string, start, stop int64) ([]string, error)
	RevRangeZSet(ctx context.Context, kind string, unique string, start, stop int64) ([]string, error)
	RevRangeZSetWithScore(ctx context.Context, kind string, unique string, start, stop int64) ([]redis.Z, error)
	RangeByScoreZSet(ctx context.Context, kind, unique, min, max string, offset, count int64) ([]string, error)
	RevRangeByScoreZSet(ctx context.Context, kind, unique, min, max string, offset, count int64) ([]string, error)
	GetCountZSet(ctx context.Context, kind, unique string) (int64, error)
	GetScoreCountZSet(ctx context.Context, kind, unique, min, max string) (int64, error)

	GetAllHash(ctx context.Context, kind string, unique string) (map[string]string, error)
	GetAnyHash(ctx context.Context, kind string, unique string, fields ...string) ([]interface{}, error)
	ExistsHash(ctx context.Context, kind string, unique string) (bool, error)
	SetFieldsHash(ctx context.Context, kind string, unique string, fieldValues ...interface{}) error
	SetFieldHash(ctx context.Context, kind string, unique string, field string, value interface{}) error
	ModifyFieldHash(ctx context.Context, kind string, unique string, field string, value interface{}) error
	GetHash(ctx context.Context, kind string, unique string, field string) (string, error)
	GetFieldsHash(ctx context.Context, kind string, unique string) ([]string, error)
	GetValuesHash(ctx context.Context, kind string, unique string) ([]string, error)
	ExistsHashField(ctx context.Context, kind string, unique string, field string) (bool, error)
	DelHashField(ctx context.Context, kind string, unique string, fields ...string) (int64, error)

	Pipeline() redis.Pipeliner
}

type CacheMethod interface {
	ToBaseInteraction(results []redis.Z) ([]*generated.Interaction, error)
	GetHistories(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error)
	GetCollections(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error)
	GetLikes(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error)
	GetUsers(ctx context.Context, creationId int64) ([]int64, error)
	GetInteraction(ctx context.Context, interaction *generated.BaseInteraction) (*generated.Interaction, error)
	GetRecommendBaseUser(ctx context.Context, id int64) ([]int64, int64, error)
	GetRecommendBaseItem(ctx context.Context, id int64) ([]int64, bool, error)
	SetRecommendBaseUser(ctx context.Context, id int64, ids []int64) error
	SetRecommendBaseItem(ctx context.Context, id int64, ids []int64) error
	UpdateHistories(ctx context.Context, data []*generated.OperateInteraction) error
	ModifyCollections(ctx context.Context, data []*generated.OperateInteraction) error
	ModifyLike(ctx context.Context, data []*generated.OperateInteraction) error
	DelHistories(ctx context.Context, data []*generated.BaseInteraction) error
	DelCollections(ctx context.Context, data []*generated.BaseInteraction) error
	DelLike(ctx context.Context, data []*generated.BaseInteraction) error
	ScanZSetsByHistories(ctx context.Context) ([]string, error)
	ScanZSetsByCreationId(ctx context.Context) ([]string, error)
	GetAllInteractions(ctx context.Context, idStrs []string) (map[int64]map[int64]float64, error)
	GetAllItemUsers(ctx context.Context, idStrs []string) (map[int64]map[int64]float64, error)
	GetPublicCreations(ctx context.Context, count int) ([]int64, error)
}
