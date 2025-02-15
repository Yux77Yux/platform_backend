package cache

import (
	"context"
)

type CacheInterface interface {
	Open(connStr string, password string) error
	Close() error

	DelKey(ctx context.Context, kind string, unique string) error

	AddToSet(ctx context.Context, kind string, unique string, values ...interface{}) error
	ExistsInSet(ctx context.Context, kind string, unique string, value interface{}) (bool, error)

	LPushList(ctx context.Context, kind string, unique string, value ...interface{}) error
	RPushList(ctx context.Context, kind string, unique string, value ...interface{}) error
	LRangeList(ctx context.Context, kind string, unique string, start, stop int64) ([]string, error)
	LTrimList(ctx context.Context, kind string, unique string, start, stop int64) error
	GetLenList(ctx context.Context, kind string, unique string) (int64, error)

	GetAllHash(ctx context.Context, kind string, unique string) (map[string]string, error)
	GetAnyHash(ctx context.Context, kind string, unique string, fields ...string) ([]interface{}, error)
	ExistsHash(ctx context.Context, kind string, unique string) (bool, error)
	IncrHash(ctx context.Context, kind string, unique string, field string, incr int64) error
	SetFieldsHash(ctx context.Context, kind string, unique string, fieldValues ...interface{}) error
	SetFieldHash(ctx context.Context, kind string, unique string, field string, value interface{}) error
	ModifyFieldHash(ctx context.Context, kind string, unique string, field string, value interface{}) error
	GetHash(ctx context.Context, kind string, unique string, field string) (string, error)
	GetFieldsHash(ctx context.Context, kind string, unique string) ([]string, error)
	GetValuesHash(ctx context.Context, kind string, unique string) ([]string, error)
	ExistsHashField(ctx context.Context, kind string, unique string, field string) (bool, error)
	DelHashField(ctx context.Context, kind string, unique string, fields ...string) (int64, error)
}
