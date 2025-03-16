package dispatch

import (
	"context"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	pkgDispatch "github.com/Yux77Yux/platform_backend/pkg/dispatch"
)

type DispatchInterface = pkgDispatch.DispatchInterface

type ChainInterface = pkgDispatch.ChainInterface

type ListenerInterface = pkgDispatch.ListenerInterface

type SqlMethod interface {
	UserAddInfoInTransaction(ctx context.Context, users []*generated.User) error
	UserRegisterInTransaction(ctx context.Context, user_credentials []*generated.UserCredentials) error
	Follow(ctx context.Context, subs []*generated.Follow) error
	UserUpdateSpaceInTransaction(ctx context.Context, users []*generated.UserUpdateSpace) error
	UserUpdateAvatarInTransaction(ctx context.Context, users []*generated.UserUpdateAvatar) error
	UserUpdateStatusInTransaction(ctx context.Context, users []*generated.UserUpdateStatus) error
	UserUpdateBioInTransaction(ctx context.Context, users []*generated.UserUpdateBio) error
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
}

type CacheMethod interface {
	StoreEmail(ctx context.Context, credentials []*generated.UserCredentials) error
	StoreUsername(ctx context.Context, credentials []*generated.UserCredentials) error
	StoreUserInfo(ctx context.Context, users []*generated.User) error
	Follow(ctx context.Context, subs []*generated.Follow) error
	UpdateUserSpace(ctx context.Context, users []*generated.UserUpdateSpace) error
	UpdateUserAvatar(ctx context.Context, users []*generated.UserUpdateAvatar) error
	UpdateUserBio(ctx context.Context, users []*generated.UserUpdateBio) error
	UpdateUserStatus(ctx context.Context, users []*generated.UserUpdateStatus) error
}
