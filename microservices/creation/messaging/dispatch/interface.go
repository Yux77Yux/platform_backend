package dispatch

import (
	"context"

	pkgDispatch "github.com/Yux77Yux/platform_backend/pkg/dispatch"
	"google.golang.org/protobuf/proto"
)

type DispatchInterface = pkgDispatch.DispatchInterface

type ChainInterface = pkgDispatch.ChainInterface

type ListenerInterface = pkgDispatch.ListenerInterface

type SqlMethod interface {
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
}

type CacheMethod interface {
}
