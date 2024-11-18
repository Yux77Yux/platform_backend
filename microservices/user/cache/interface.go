package cache

import (
	"context"
)

type CacheInterface interface {
	Open(connStr string, password string) error
	Close()

	AddToSet(ctx context.Context, kind string, unique string, value interface{}) error
	ExistsInSet(ctx context.Context, kind string, unique string, value interface{}) (bool, error)
}
