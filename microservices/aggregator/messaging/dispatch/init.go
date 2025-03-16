package dispatch

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	LISTENER_CHANNEL_COUNT = 120
	MAX_BATCH_SIZE         = 50
	EXE_CHANNEL_COUNT      = 5
)

var (
	messaging MessageQueueMethod
)

type Dispatch struct {
	chainMap map[string]ChainInterface
}

func (d *Dispatch) HandleRequest(msg protoreflect.ProtoMessage, typeName string) {
	copy := proto.Clone(msg)
	d.chainMap[typeName].HandleRequest(copy)
}

func (d *Dispatch) Close() {
	for _, chain := range d.chainMap {
		s := make(chan any, 1)
		chain.Close(s)
		<-s
	}
}

func Run() DispatchInterface {
	_dispatch := &Dispatch{
		chainMap: make(map[string]ChainInterface),
	}

	_dispatch.chainMap[AddView] = InitialAddViewChain()

	return _dispatch
}
