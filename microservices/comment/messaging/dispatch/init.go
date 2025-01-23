package dispatch

import (
	"sync"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
)

var (
	deleteChain     *DeleteChain
	delListenerPool = sync.Pool{
		New: func() any {
			return &DeleteListener{
				commentChannel:  make(chan *generated.AfterAuth, 120),
				timeoutDuration: 10 * time.Second,
				updateInterval:  3 * time.Second,
			}
		},
	}
	delCommentsPool = sync.Pool{
		New: func() any {
			slice := make([]*generated.AfterAuth, 0, 50)
			return &slice
		},
	}

	insertChain        *InsertChain
	insertListenerPool = sync.Pool{
		New: func() any {
			return &InsertListener{
				commentChannel:  make(chan *generated.Comment, 160),
				timeoutDuration: 12 * time.Second,
				updateInterval:  3 * time.Second,
			}
		},
	}
	insertCommentsPool = sync.Pool{
		New: func() any {
			slice := make([]*generated.Comment, 0, 50)
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
