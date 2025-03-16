package dispatch

import (
	"context"

	"google.golang.org/protobuf/proto"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	pkgDispatch "github.com/Yux77Yux/platform_backend/pkg/dispatch"
	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue/rabbitmq"
)

type HandlerFunc = pkgMQ.HandlerFunc
type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
	ListenToQueue(exchange, queueName, routeKey string, handler HandlerFunc)
}

type CacheMethod interface {
	AddIpInSet(ctx context.Context, req *common.ViewCreation) error
	ExistIpInSet(ctx context.Context, req *common.ViewCreation) (bool, error)
}

type DispatchInterface = pkgDispatch.DispatchInterface

type ChainInterface = pkgDispatch.ChainInterface

type ListenerInterface = pkgDispatch.ListenerInterface
