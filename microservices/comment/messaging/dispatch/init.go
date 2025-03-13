package dispatch

import (
	"sync"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
)

const (
	Insert = "Insert"
	Delete = "Delete"

	LISTENER_CHANNEL_COUNT = 120
	MAX_BATCH_SIZE         = 50
	EXE_CHANNEL_COUNT      = 5
)

var (
	deleteChain     *DeleteChain
	delCommentsPool = sync.Pool{
		New: func() any {
			slice := make([]*common.AfterAuth, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	insertChain        *InsertChain
	insertCommentsPool = sync.Pool{
		New: func() any {
			slice := make([]*generated.Comment, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}
)

func init() {
	// 初始化责任链
	insertChain = InitialInsertChain()
	deleteChain = InitialDeleteChain()
}

func HandleRequest(msg protoreflect.ProtoMessage, typeName string) {
	copy := proto.Clone(msg)

	switch typeName {
	case Insert:
		insertChain.HandleRequest(copy)
	case Delete:
		deleteChain.HandleRequest(copy)
	}
}
