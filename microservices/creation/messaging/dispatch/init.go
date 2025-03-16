package dispatch

import (
	"sync"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	UpdateCount = "InteractionCount"

	LISTENER_CHANNEL_COUNT = 120
	MAX_BATCH_SIZE         = 50
	EXE_CHANNEL_COUNT      = 5
)

func init() {
	// 初始化责任链
	updateCountChain = InitialUpdateCountChain()
}

func HandleRequest(msg protoreflect.ProtoMessage, typeName string) {
	copy := proto.Clone(msg)
	switch typeName {
	case UpdateCount:
		updateCountChain.HandleRequest(copy)
	}
}

var (
	db        SqlMethod
	messaging MessageQueueMethod
	cache     CacheMethod
)

func InitDb(_db SqlMethod) {
	db = _db
}

func InitMQ(_messaging MessageQueueMethod) {
	messaging = _messaging
}

func InitCache(_cache CacheMethod) {
	cache = _cache
}
