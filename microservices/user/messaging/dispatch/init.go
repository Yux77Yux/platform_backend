package dispatch

import (
	"sync"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
)

var (
	deleteChain     *DeleteChain
	delListenerPool = sync.Pool{
		New: func() any {
			return &DeleteListener{
				userChannel:     make(chan *generated.User, 30),
				timeoutDuration: 10 * time.Second,
				updateInterval:  3 * time.Second,
			}
		},
	}
	delusersPool = sync.Pool{
		New: func() any {
			slice := make([]*generated.AfterAuth, 0, 50)
			return &slice
		},
	}

	insertChain        *InsertChain
	insertListenerPool = sync.Pool{
		New: func() any {
			return &InsertListener{
				userChannel:     make(chan *generated.User, 50),
				timeoutDuration: 20 * time.Second,
				updateInterval:  3 * time.Second,
			}
		},
	}
	insertusersPool = sync.Pool{
		New: func() any {
			slice := make([]*generated.User, 0, 50)
			return &slice
		},
	}
)

func Init() {
	// 初始化责任链
	insertChain = InitialInsertChain()
	deleteChain = InitialDeleteChain()
}

func HandleRequest(msg protoreflect.ProtoMessage, typeName string) {
	switch typeName {
	case "insert":
		insertChain.HandleRequest(msg)
	case "delete":
		deleteChain.HandleRequest(msg)
	}
}
