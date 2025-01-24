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
)

func Init() {
	// 初始化责任链

	insertUsersCacheChain = InitialInsertCacheChain()
	insertUsersChain = InitialInsertChain()

	registerCacheChain = InitialRegisterCacheChain()
	registerChain = InitialRegisterChain()
}

func HandleRequest(msg protoreflect.ProtoMessage, typeName string) {
	switch typeName {
	case "insert":
		insertUsersChain.HandleRequest(msg)
	case "register":
		registerChain.HandleRequest(msg)
	}
}
