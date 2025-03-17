package dispatch

import (
	"context"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
	pkgDispatch "github.com/Yux77Yux/platform_backend/pkg/dispatch"
)

type DispatchInterface = pkgDispatch.DispatchInterface

type ChainInterface = pkgDispatch.ChainInterface

type ListenerInterface = pkgDispatch.ListenerInterface

type SqlMethod interface {
	PostReviews(ctx context.Context, reviews []*generated.NewReview) error
	UpdateReviews(ctx context.Context, reviews []*generated.Review) error
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
}
