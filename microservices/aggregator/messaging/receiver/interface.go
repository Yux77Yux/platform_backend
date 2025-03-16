package receiver

import (
	"context"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue/rabbitmq"
)

type HandlerFunc = pkgMQ.HandlerFunc

type CacheMethod interface {
	AddIpInSet(ctx context.Context, req *common.ViewCreation) error
	ExistIpInSet(ctx context.Context, req *common.ViewCreation) (bool, error)
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
	ListenToQueue(exchange, queueName, routeKey string, handler HandlerFunc)
}

type DispatchInterface interface {
	HandleRequest(msg protoreflect.ProtoMessage, typeName string)
}
