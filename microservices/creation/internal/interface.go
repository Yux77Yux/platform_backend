package internal

import (
	"context"

	"google.golang.org/protobuf/proto"
)

type SqlMethod interface {
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
}

type CacheMethod interface {
}
