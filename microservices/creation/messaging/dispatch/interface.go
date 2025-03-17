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
	UpdateCreationCount(ctx context.Context, creationId int64, saveCount, likeCount, viewCount int32) error
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
}
