package dispatch

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
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
	chainMap := make(map[string]ChainInterface)
	chainMap[UpdateCount] = InitialUpdateCountChain()

	_dispatch := &Dispatch{
		chainMap: chainMap,
	}
	return _dispatch
}
