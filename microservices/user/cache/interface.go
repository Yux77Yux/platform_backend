package cache

import (
	"context"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	"github.com/go-redis/redis/v8"
)

type CacheInterface interface {
	Open(connStr string, password string) error
	Close() error

	AddToSet(ctx context.Context, kind string, unique string, values ...interface{}) error
	ExistsInSet(ctx context.Context, kind string, unique string, value interface{}) (bool, error)

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
	TxPipeline() redis.Pipeliner
}

type CacheMethod interface {
	ExistsUsername(ctx context.Context, username string) (bool, error)
	ExistsEmail(ctx context.Context, email string) (bool, error)
	ExistsUserInfo(ctx context.Context, user_id int64) (bool, error)
	StoreEmail(ctx context.Context, credentials []*generated.UserCredentials) error
	StoreUsername(ctx context.Context, credentials []*generated.UserCredentials) error
	StoreUserInfo(ctx context.Context, users []*generated.User) error
	Follow(ctx context.Context, subs []*generated.Follow) error
	UpdateUserSpace(ctx context.Context, users []*generated.UserUpdateSpace) error
	UpdateUserAvatar(ctx context.Context, users []*generated.UserUpdateAvatar) error
	UpdateUserBio(ctx context.Context, users []*generated.UserUpdateBio) error
	UpdateUserStatus(ctx context.Context, users []*generated.UserUpdateStatus) error
	GetUserInfo(ctx context.Context, user_id int64, fields []string) (map[string]string, error)
	GetUserCredentials(ctx context.Context, userCrdentials *generated.UserCredentials) (*generated.UserCredentials, error)
	GetUserCards(ctx context.Context, userIds []int64) ([]*common.UserCreationComment, error)
	GetFolloweesByTime(ctx context.Context, userId int64, page int32) ([]int64, error)
	GetFolloweesByView(ctx context.Context, userId int64, page int32) ([]int64, error)
	GetFollowers(ctx context.Context, userId int64, page int32) ([]int64, error)
	CancelFollow(ctx context.Context, follow *generated.Follow) error
	DelCredentials(ctx context.Context, username string) error
}
