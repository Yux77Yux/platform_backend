package internal

import (
	"context"

	"google.golang.org/protobuf/proto"

	common "github.com/Yux77Yux/platform_backend/generated/common"
)

type CacheMethod interface {
	AddIpInSet(ctx context.Context, req *common.ViewCreation) error
	ExistIpInSet(ctx context.Context, req *common.ViewCreation) (bool, error)
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
}
