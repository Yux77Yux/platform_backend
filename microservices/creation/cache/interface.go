package cache

import (
	"context"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CacheInterface interface {
	Open(connStr string, password string) error
	Close() error

	DelKey(ctx context.Context, kind string, unique string) error

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

	Pipeline() redis.Pipeliner
	TxPipeline() redis.Pipeliner
}

type CacheMethod interface {
	CreationAddInCache(ctx context.Context, creationInfo *generated.CreationInfo) error
	AddSpaceCreations(ctx context.Context, authorId, creationId int64, publishTime *timestamppb.Timestamp) error
	GetCreationInfoFields(ctx context.Context, creation_id int64, fields []string) (map[string]string, error)
	GetCreationInfo(ctx context.Context, creation_id int64) (*generated.CreationInfo, error)
	parseIntField(value string, bitSize int) (int64, error)
	mapToCreationInfo(results map[string]string, creation_id int64) (*generated.CreationInfo, error)
	GetSimilarCreationList(ctx context.Context, creation_id int64) ([]int64, error)
	GetSpaceCreationList(ctx context.Context, user_id int64, page int32, typeStr string) ([]int64, int32, error)
	GetHistoryCreationList(ctx context.Context, user_id int64) ([]int64, error)
	GetCollectionCreationList(ctx context.Context, user_id int64) ([]int64, error)
	GetCreationList(ctx context.Context, creationIds []int64) ([]*generated.CreationInfo, []int64, error)
	UpdateCreation(ctx context.Context, creation *generated.CreationUpdated) error
	UpdateCreationStatus(ctx context.Context, creation *generated.CreationUpdateStatus) error
	getAuthorIdMap(ctx context.Context, creationIds []int64) (map[int64]string, error)
	UpdateCreationCount(ctx context.Context, actions []*common.UserAction) error
}
