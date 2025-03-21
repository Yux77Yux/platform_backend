package repository

import (
	"context"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	pkgDb "github.com/Yux77Yux/platform_backend/pkg/database"
)

type SqlInterface = pkgDb.SqlInterface

type SqlMethod interface {
	UserAddInfoInTransaction(ctx context.Context, users []*generated.User) error
	UserRegisterInTransaction(ctx context.Context, user_credentials []*generated.UserCredentials) error
	Follow(ctx context.Context, subs []*generated.Follow) error
	Exists(ctx context.Context, isEmail bool, usernameOrEmail string) (bool, error)
	UserGetInfoInTransaction(ctx context.Context, id int64) (*generated.User, error)
	GetUsers(ctx context.Context, userIds []int64) ([]*common.UserCreationComment, error)
	GetFolloweers(ctx context.Context, userId int64, page int32) ([]*common.UserCreationComment, error)
	GetFolloweesByTime(ctx context.Context, userId int64, page int32) ([]*common.UserCreationComment, error)
	GetFolloweesByViews(ctx context.Context, userId int64, page int32) ([]*common.UserCreationComment, error)
	UserVerifyInTranscation(ctx context.Context, user_credential *generated.UserCredentials) (*generated.UserCredentials, error)
	UserEmailUpdateInTransaction(ctx context.Context, user_credentials []*generated.UserCredentials) error
	UserUpdateSpaceInTransaction(ctx context.Context, users []*generated.UserUpdateSpace) error
	UserUpdateAvatarInTransaction(ctx context.Context, users []*generated.UserUpdateAvatar) error
	UserUpdateStatusInTransaction(ctx context.Context, users []*generated.UserUpdateStatus) error
	UserUpdateBioInTransaction(ctx context.Context, users []*generated.UserUpdateBio) error
	DelReviewer(ctx context.Context, reviewerId int64) (string, string, error)
	ViewFollowee(ctx context.Context, subs []*generated.Follow) error
	CancelFollow(ctx context.Context, f *generated.Follow) error
	ExistsFollowee(ctx context.Context, followerId, followeeId int64) (bool, error)
}
