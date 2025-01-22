package messaging

import (
	"google.golang.org/protobuf/reflect/protoreflect"

	chain "github.com/Yux77Yux/platform_backend/microservices/user/messaging/chain"
)

type Dispatcher struct {
	chains []chain.ChainInterface
}

func (d *Dispatcher) Dispatch(msg protoreflect.ProtoMessage, typeName string) error {
	var target chain.ChainInterface = nil
	for _, val := range d.chains {
		if typeName == target.GetChainType() {
			target = val
			break
		}
	}
	if target == nil {
		target = chain.InitialChain(typeName)
		d.chains = append(d.chains, target)
	}

	target.HandleRequest(msg)

	return nil
}
