package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type CacheInterface interface {
	Open(connStr string, password string) error
	Close() error

	ScanSet(ctx context.Context, kind string, fliter string, cursor uint64, count int64) ([]string, uint64, error)
	AddToSet(ctx context.Context, kind string, unique string, values ...interface{}) error
	RemSet(ctx context.Context, kind string, unique string, values ...interface{}) error
	ExistsInSet(ctx context.Context, kind string, unique string, value interface{}) (bool, error)
	CountSet(ctx context.Context, kind string, unique string) (int64, error)
	GetMembersSet(ctx context.Context, kind string, unique string) ([]string, error)

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
