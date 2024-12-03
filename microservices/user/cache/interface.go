package cache

import (
	"context"
)

type CacheInterface interface {
	Open(connStr string, password string) error
	Close() error

	AddToSet(ctx context.Context, kind string, unique string, value interface{}) error
	ExistsInSet(ctx context.Context, kind string, unique string, value interface{}) (bool, error)

	GetAllHash(ctx context.Context, kind string, unique string) (map[string]string, error)
	ExistsHash(ctx context.Context, kind string, unique string) (bool, error)
	SetFieldsHash(ctx context.Context, kind string, unique string, fieldValues ...interface{}) error
	SetFieldHash(ctx context.Context, kind string, unique string, field string, value interface{}) error
	ModifyFieldHash(ctx context.Context, kind string, unique string, field string, value interface{}) error
	GetHash(ctx context.Context, kind string, unique string, field string) (string, error)
	GetFieldsHash(ctx context.Context, kind string, unique string) ([]string, error)
	GetValuesHash(ctx context.Context, kind string, unique string) ([]string, error)
	ExistsHashField(ctx context.Context, kind string, unique string, field string) (bool, error)
	DelHash(ctx context.Context, kind string, unique string, fields ...string) (int64, error)
}
