package dispatch

import (
	"context"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	pkgDispatch "github.com/Yux77Yux/platform_backend/pkg/dispatch"
)

type DispatchInterface = pkgDispatch.DispatchInterface

type ChainInterface = pkgDispatch.ChainInterface

type ListenerInterface = pkgDispatch.ListenerInterface

type SqlMethod interface {
	UpdateInteractions(ctx context.Context, req []*generated.OperateInteraction) error
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
}

type CacheMethod interface {
	UpdateHistories(ctx context.Context, data []*generated.OperateInteraction) error
	ModifyCollections(ctx context.Context, data []*generated.OperateInteraction) error
	ModifyLike(ctx context.Context, data []*generated.OperateInteraction) error
	DelLike(ctx context.Context, data []*generated.BaseInteraction) error
}
