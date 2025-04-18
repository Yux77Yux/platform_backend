package internal

import (
	"context"

	"google.golang.org/protobuf/proto"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
)

type SqlMethod interface {
	Exists(ctx context.Context, isEmail bool, usernameOrEmail string) (bool, error)
	UserGetInfoInTransaction(ctx context.Context, id int64) (*generated.User, error)
	GetUsers(ctx context.Context, userIds []int64) ([]*common.UserCreationComment, error)
	GetFolloweers(ctx context.Context, userId int64, page int32) ([]*common.UserCreationComment, error)
	GetFolloweesByTime(ctx context.Context, userId int64, page int32) ([]*common.UserCreationComment, error)
	GetFolloweesByViews(ctx context.Context, userId int64, page int32) ([]*common.UserCreationComment, error)
	UserVerifyInTranscation(ctx context.Context, user_credential *generated.UserCredentials) (*generated.UserCredentials, error)
	CancelFollow(ctx context.Context, f *generated.Follow) error
	ExistsFollowee(ctx context.Context, followerId, followeeId int64) (bool, error)
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
}

type CacheMethod interface {
	ExistsEmail(ctx context.Context, email string) (bool, error)
	ExistsUsername(ctx context.Context, username string) (bool, error)
	GetUserInfo(ctx context.Context, user_id int64, fields []string) (map[string]string, error)
	GetUserCredentials(ctx context.Context, userCrdentials *generated.UserCredentials) (*generated.UserCredentials, error)
	CancelFollow(ctx context.Context, follow *generated.Follow) error
}
