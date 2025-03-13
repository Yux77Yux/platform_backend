package dispatch

import (
	"sync"

	"google.golang.org/protobuf/proto"
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
	// 用 msg 的类型创建一个新对象
	copy := proto.Clone(msg)

	switch typeName {
	case Register:
		registerChain.HandleRequest(copy)
	case RegisterCache:
		registerCacheChain.HandleRequest(copy)

	case InsertUser:
		insertUsersChain.HandleRequest(copy)
	case InsertUserCache:
		insertUsersCacheChain.HandleRequest(copy)

	case UpdateUserAvatar:
		userAvatarChain.HandleRequest(copy)
	case UpdateUserAvatarCache:
		userAvatarCacheChain.HandleRequest(copy)

	case UpdateUserSpace:
		userSpaceChain.HandleRequest(copy)
	case UpdateUserSpaceCache:
		userSpaceCacheChain.HandleRequest(copy)

	case UpdateUserBio:
		userBioChain.HandleRequest(copy)
	case UpdateUserBioCache:
		userBioCacheChain.HandleRequest(copy)

	case UpdateUserStatus:
		userStatusChain.HandleRequest(copy)
	case UpdateUserStatusCache:
		userStatusCacheChain.HandleRequest(copy)

	case Follow:
		followChain.HandleRequest(copy)
	case FollowCache:
		followCacheChain.HandleRequest(copy)
	}
}
