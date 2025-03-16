package dispatch

import (
	"sync"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
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

var (
	db    SqlMethod
	cache CacheMethod
)

func InitDb(_db SqlMethod) {
	db = _db
}

func InitCache(_cache CacheMethod) {
	cache = _cache
}

func Run() DispatchInterface {
	chainMap := make(map[string]ChainInterface)
	chainMap[]
	_dispatch := &Dispatch{
		chainMap: chainMap,
	}

	return _dispatch
}
