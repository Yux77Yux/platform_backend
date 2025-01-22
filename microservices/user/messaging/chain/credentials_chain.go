package chain

import (
	"log"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
)

func InitialChain(typeName string) ChainInterface {
	switch typeName {

	case "credentials":
		return &CredentialsChain{
			typeName:  typeName,
			head:      &CredentialsListener{},
			listeners: make([]*CredentialsListener, 0, 2),
			length:    0,
			capacity:  2,
		}

	default:
		log.Printf("do not exist %s type", typeName)
		return nil
	}

}

// 责任链
type CredentialsChain struct {
	typeName string
	head     *CredentialsListener // 责任链的头部
}

func (chain *CredentialsChain) GetChainType() string {
	return chain.typeName
}

func (chain *CredentialsChain) HandleRequest(msg protoreflect.ProtoMessage) {
	// listener := chain.findListener()
}
