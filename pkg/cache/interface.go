package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisMethods interface {
	Open(connStr string, password string) error
	Close() error

	DelRelatedKeys(ctx context.Context, kindPrefix string, kindType string) error
	DelKey(ctx context.Context, kind string, unique string) error

	SetString(ctx context.Context, kind string, unique string, value interface{}) error
	ModifyString(ctx context.Context, kind string, unique string, value interface{}) error
	GetString(ctx context.Context, kind string, unique string) (string, error)
	ExistsString(ctx context.Context, kind string, unique string) (bool, error)

	ScanHash(ctx context.Context, kind string, unique string, fliter string, cursor uint64, count int64) ([]string, uint64, error)
	ExistsHash(ctx context.Context, kind string, unique string) (bool, error)
	IncrHash(ctx context.Context, kind string, unique string, field string, incr int64) error
	SetFieldsHash(ctx context.Context, kind string, unique string, fieldValues ...interface{}) error
	SetFieldHash(ctx context.Context, kind string, unique string, field string, value interface{}) error
	ModifyFieldHash(ctx context.Context, kind string, unique string, field string, value interface{}) error
	GetHash(ctx context.Context, kind string, unique string, field string) (string, error)
	GetAnyHash(ctx context.Context, kind string, unique string, fields ...string) ([]interface{}, error)
	GetAllHash(ctx context.Context, kind string, unique string) (map[string]string, error)
	GetFieldsHash(ctx context.Context, kind string, unique string) ([]string, error)
	GetValuesHash(ctx context.Context, kind string, unique string) ([]string, error)
	ExistsHashField(ctx context.Context, kind string, unique string, field string) (bool, error)
	GetLenHash(ctx context.Context, kind string, unique string, field string) (int64, error)
	DelHashField(ctx context.Context, kind string, unique string, fields ...string) (int64, error)

	LPushList(ctx context.Context, kind string, unique string, value ...interface{}) error
	RPushList(ctx context.Context, kind string, unique string, value ...interface{}) error
	LPopList(ctx context.Context, kind string, unique string) (string, error)
	RPopList(ctx context.Context, kind string, unique string) (string, error)
	LRangeList(ctx context.Context, kind string, unique string, start, stop int64) ([]string, error)
	InsertBeforeList(ctx context.Context, kind string, unique string, pivot, value interface{}) error
	InsertAfterList(ctx context.Context, kind string, unique string, pivot, value interface{}) error
	IndexList(ctx context.Context, kind string, unique string, index int64) (string, error)
	FindElementList(ctx context.Context, kind string, unique string, value string) (int64, error)
	FindElementsList(ctx context.Context, kind string, unique string, value string, count int64) ([]int64, error)
	LTrimList(ctx context.Context, kind string, unique string, start, stop int64) error
	GetLenList(ctx context.Context, kind string, unique string) (int64, error)
	GetElementsList(ctx context.Context, kind string, unique string, start, stop int64) ([]string, error)

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

	SetBit(ctx context.Context, kind string, unique string, offset int64, value bool) error
	GetBit(ctx context.Context, kind string, unique string, offset int64) (bool, error)
	ClearBit(ctx context.Context, kind string, unique string, offset int64) error
	ClearRangeBits(ctx context.Context, kind string, unique string, startOffset, endOffset int64) error
	CountBit(ctx context.Context, kind string, unique string, startOffset, endOffset int64) (int64, error)
	OpertorBit(ctx context.Context, operation string, destKind string, unique string, srcKinds []string) error
	ModifyBit(ctx context.Context, kind string, unique string, offset int64, newValue bool) error
	ResetBitmap(ctx context.Context, kind string, unique string, length int64) error
	FindPositionBit(ctx context.Context, kind string, unique string, bit bool) (int64, error)
	CombineBitmaps(ctx context.Context, destKind string, unique string, srcKinds []string) error
	AddToBloomFilter(ctx context.Context, kind string, unique string) error
	CheckBloomFilter(ctx context.Context, kind string, unique string) (bool, error)

	BitField(ctx context.Context, kind string, unique string, command ...interface{}) ([]int64, error)

	Pipeline() redis.Pipeliner
	TxPipeline() redis.Pipeliner
}
