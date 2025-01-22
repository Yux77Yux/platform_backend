package chain

import "google.golang.org/protobuf/reflect/protoreflect"

type ListenerInterface interface {
}

type ChainInterface interface {
	GetChainType() string
	HandleRequest(protoreflect.ProtoMessage)
}
