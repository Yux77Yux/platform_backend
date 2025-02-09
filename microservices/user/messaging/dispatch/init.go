package dispatch

import (
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
)

const (
	LISTENER_CHANNEL_COUNT = 80
	MAX_BATCH_SIZE         = 50
	EXE_CHANNEL_COUNT      = 5

	Register      = "Register"
	RegisterCache = "RegisterCache"

	InsertUser      = "InsertUser"
	InsertUserCache = "InsertUserCache"

	UpdateUserAvatar      = "UpdateUserAvatar"
	UpdateUserAvatarCache = "UpdateUserAvatarCache"

	UpdateUserSpace      = "UpdateUserSpace"
	UpdateUserSpaceCache = "UpdateUserSpaceCache"

	UpdateUserStatus      = "UpdateUserStatus"
	UpdateUserStatusCache = "UpdateUserStatusCache"

	UpdateUserBio      = "UpdateUserBio"
	UpdateUserBioCache = "UpdateUserBioCache"

	Follow      = "Follow"
	FollowCache = "FollowCache"
)

var (
	registerChain             *RegisterChain
	registerCacheChain        *RegisterCacheChain
	insertUserCredentialsPool = sync.Pool{
		New: func() any {
			slice := make([]*generated.UserCredentials, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	insertUsersChain      *InsertChain
	insertUsersCacheChain *InsertCacheChain
	insertUsersPool       = sync.Pool{
		New: func() any {
			slice := make([]*generated.User, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	userAvatarChain      *UserAvatarChain
	userAvatarCacheChain *UserAvatarCacheChain
	userAvatarPool       = sync.Pool{
		New: func() any {
			slice := make([]*generated.UserUpdateAvatar, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	userSpaceChain      *UserSpaceChain
	userSpaceCacheChain *UserSpaceCacheChain
	userSpacePool       = sync.Pool{
		New: func() any {
			slice := make([]*generated.UserUpdateSpace, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	userBioChain      *UserBioChain
	userBioCacheChain *UserBioCacheChain
	userBioPool       = sync.Pool{
		New: func() any {
			slice := make([]*generated.UserUpdateBio, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	userStatusChain      *UserStatusChain
	userStatusCacheChain *UserStatusCacheChain
	userStatusPool       = sync.Pool{
		New: func() any {
			slice := make([]*generated.UserUpdateStatus, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	followChain      *FollowChain
	followCacheChain *FollowCacheChain
	followPool       = sync.Pool{
		New: func() any {
			slice := make([]*generated.Follow, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}
)

func init() {
	// 初始化责任链

	insertUsersCacheChain = InitialInsertCacheChain()
	insertUsersChain = InitialInsertChain()

	registerCacheChain = InitialRegisterCacheChain()
	registerChain = InitialRegisterChain()

	userAvatarChain = InitialUserAvatarChain()
	userAvatarCacheChain = InitialUserAvatarCacheChain()

	userSpaceChain = InitialUserSpaceChain()
	userSpaceCacheChain = InitialUserSpaceCacheChain()

	userBioChain = InitialUserBioChain()
	userBioCacheChain = InitialUserBioCacheChain()

	userStatusChain = InitialUserStatusChain()
	userStatusCacheChain = InitialUserStatusCacheChain()

	followChain = InitialFollowChain()
	followCacheChain = InitialFollowCacheChain()
}

func HandleRequest(msg protoreflect.ProtoMessage, typeName string) {
	switch typeName {
	case Register:
		registerChain.HandleRequest(msg)
	case RegisterCache:
		registerCacheChain.HandleRequest(msg)

	case InsertUser:
		insertUsersChain.HandleRequest(msg)
	case InsertUserCache:
		insertUsersCacheChain.HandleRequest(msg)

	case UpdateUserAvatar:
		userAvatarChain.HandleRequest(msg)
	case UpdateUserAvatarCache:
		userAvatarCacheChain.HandleRequest(msg)

	case UpdateUserSpace:
		userSpaceChain.HandleRequest(msg)
	case UpdateUserSpaceCache:
		userSpaceCacheChain.HandleRequest(msg)

	case UpdateUserBio:
		userBioChain.HandleRequest(msg)
	case UpdateUserBioCache:
		userBioCacheChain.HandleRequest(msg)

	case UpdateUserStatus:
		userStatusChain.HandleRequest(msg)
	case UpdateUserStatusCache:
		userStatusCacheChain.HandleRequest(msg)

	case Follow:
		followChain.HandleRequest(msg)
	case FollowCache:
		followCacheChain.HandleRequest(msg)
	}
}
