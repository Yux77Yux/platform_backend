package dispatch

import (
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"

	common "github.com/Yux77Yux/platform_backend/generated/common"
)

const (
	AddView = "InteractionCount"

	LISTENER_CHANNEL_COUNT = 120
	MAX_BATCH_SIZE         = 50
	EXE_CHANNEL_COUNT      = 5
)

var (
	addViewChain *AddViewChain
	insertPool   = sync.Pool{
		New: func() any {
			slice := make([]*common.UserAction, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}
)

func init() {
	// 初始化责任链
	addViewChain = InitialAddViewChain()
}

func HandleRequest(msg protoreflect.ProtoMessage, typeName string) {
	switch typeName {
	case AddView:
		addViewChain.HandleRequest(msg)
	}
}
