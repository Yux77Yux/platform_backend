package dispatch

import "google.golang.org/protobuf/reflect/protoreflect"

type DispatchInterface interface {
	HandleRequest(msg protoreflect.ProtoMessage, typeName string)
	Close()
}

type ChainInterface interface {
	FindListener(protoreflect.ProtoMessage) ListenerInterface
	DestroyListener(ListenerInterface)
	CreateListener(protoreflect.ProtoMessage) ListenerInterface
	HandleRequest(protoreflect.ProtoMessage)
	ExecuteBatch()
	Close(signal chan any)
	GetPoolObj() any
}

type ListenerInterface interface {
	GetId() int64
	StartListening()
	Dispatch(protoreflect.ProtoMessage)
	RestartUpdateIntervalTimer()
	RestartTimeoutTimer()
	SendBatch()
	Cleanup()
}
