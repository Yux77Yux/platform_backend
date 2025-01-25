package dispatch

import (
	"sync"
	"time"

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
)

var (
	registerChain        *RegisterChain
	registerCacheChain   *RegisterCacheChain
	registerListenerPool = sync.Pool{
		New: func() any {
			return &RegisterListener{
				userCredentialsChannel: make(chan *generated.UserCredentials, LISTENER_CHANNEL_COUNT),
				timeoutDuration:        10 * time.Second,
				updateInterval:         3 * time.Second,
			}
		},
	}
	registerCacheListenerPool = sync.Pool{
		New: func() any {
			return &RegisterCacheListener{
				userCredentialsChannel: make(chan *generated.UserCredentials, LISTENER_CHANNEL_COUNT),
				timeoutDuration:        10 * time.Second,
				updateInterval:         3 * time.Second,
			}
		},
	}
	insertUserCredentialsPool = sync.Pool{
		New: func() any {
			slice := make([]*generated.UserCredentials, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	insertUsersChain        *InsertChain
	insertUsersCacheChain   *InsertCacheChain
	insertUsersListenerPool = sync.Pool{
		New: func() any {
			return &InsertListener{
				usersChannel:    make(chan *generated.User, LISTENER_CHANNEL_COUNT),
				timeoutDuration: 10 * time.Second,
				updateInterval:  3 * time.Second,
			}
		},
	}
	insertUsersCacheListenerPool = sync.Pool{
		New: func() any {
			return &InsertCacheListener{
				usersChannel:    make(chan *generated.User, LISTENER_CHANNEL_COUNT),
				timeoutDuration: 10 * time.Second,
				updateInterval:  3 * time.Second,
			}
		},
	}
	insertUsersPool = sync.Pool{
		New: func() any {
			slice := make([]*generated.User, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	userAvatarChain        *UserAvatarChain
	userAvatarCacheChain   *UserAvatarCacheChain
	userAvatarListenerPool = sync.Pool{
		New: func() any {
			return &UserAvatarListener{
				userUpdateAvatarChannel: make(chan *generated.UserUpdateAvatar, LISTENER_CHANNEL_COUNT),
				timeoutDuration:         10 * time.Second,
				updateInterval:          3 * time.Second,
			}
		},
	}
	userAvatarCacheListenerPool = sync.Pool{
		New: func() any {
			return &UserAvatarCacheListener{
				userUpdateAvatarChannel: make(chan *generated.UserUpdateAvatar, LISTENER_CHANNEL_COUNT),
				timeoutDuration:         10 * time.Second,
				updateInterval:          3 * time.Second,
			}
		},
	}
	userAvatarPool = sync.Pool{
		New: func() any {
			slice := make([]*generated.UserUpdateAvatar, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	userSpaceChain        *UserSpaceChain
	userSpaceCacheChain   *UserSpaceCacheChain
	userSpaceListenerPool = sync.Pool{
		New: func() any {
			return &UserSpaceListener{
				userUpdateSpaceChannel: make(chan *generated.UserUpdateSpace, LISTENER_CHANNEL_COUNT),
				timeoutDuration:        10 * time.Second,
				updateInterval:         3 * time.Second,
			}
		},
	}
	userSpaceCacheListenerPool = sync.Pool{
		New: func() any {
			return &UserSpaceCacheListener{
				userUpdateSpaceChannel: make(chan *generated.UserUpdateSpace, LISTENER_CHANNEL_COUNT),
				timeoutDuration:        10 * time.Second,
				updateInterval:         3 * time.Second,
			}
		},
	}
	userSpacePool = sync.Pool{
		New: func() any {
			slice := make([]*generated.UserUpdateSpace, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	userBioChain        *UserBioChain
	userBioCacheChain   *UserBioCacheChain
	userBioListenerPool = sync.Pool{
		New: func() any {
			return &UserBioListener{
				userUpdateBioChannel: make(chan *generated.UserUpdateBio, LISTENER_CHANNEL_COUNT),
				timeoutDuration:      10 * time.Second,
				updateInterval:       3 * time.Second,
			}
		},
	}
	userBioCacheListenerPool = sync.Pool{
		New: func() any {
			return &UserBioCacheListener{
				userUpdateBioChannel: make(chan *generated.UserUpdateBio, LISTENER_CHANNEL_COUNT),
				timeoutDuration:      10 * time.Second,
				updateInterval:       3 * time.Second,
			}
		},
	}
	userBioPool = sync.Pool{
		New: func() any {
			slice := make([]*generated.UserUpdateBio, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	userStatusChain        *UserStatusChain
	userStatusCacheChain   *UserStatusCacheChain
	userStatusListenerPool = sync.Pool{
		New: func() any {
			return &UserStatusListener{
				userUpdateStatusChannel: make(chan *generated.UserUpdateStatus, LISTENER_CHANNEL_COUNT),
				timeoutDuration:         10 * time.Second,
				updateInterval:          3 * time.Second,
			}
		},
	}
	userStatusCacheListenerPool = sync.Pool{
		New: func() any {
			return &UserStatusCacheListener{
				userUpdateStatusChannel: make(chan *generated.UserUpdateStatus, LISTENER_CHANNEL_COUNT),
				timeoutDuration:         10 * time.Second,
				updateInterval:          3 * time.Second,
			}
		},
	}
	userStatusPool = sync.Pool{
		New: func() any {
			slice := make([]*generated.UserUpdateStatus, 0, MAX_BATCH_SIZE)
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

	}
}
