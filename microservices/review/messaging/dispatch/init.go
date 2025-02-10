package dispatch

import (
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
)

const (
	Insert = "Insert"
	Update = "Update"

	LISTENER_CHANNEL_COUNT = 120
	MAX_BATCH_SIZE         = 50
	EXE_CHANNEL_COUNT      = 5
)

var (
	insertChain *InsertChain
	insertPool  = sync.Pool{
		New: func() any {
			slice := make([]*generated.NewReview, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}

	updateChain *UpdateChain
	updatePool  = sync.Pool{
		New: func() any {
			slice := make([]*generated.Review, 0, MAX_BATCH_SIZE)
			return &slice
		},
	}
)

func init() {
	// 初始化责任链
	insertChain = InitialInsertChain()
	updateChain = InitialUpdateChain()
}

func HandleRequest(msg protoreflect.ProtoMessage, typeName string) {
	switch typeName {
	case Insert:
		insertChain.HandleRequest(msg)
	case Update:
		updateChain.HandleRequest(msg)
	}
}
