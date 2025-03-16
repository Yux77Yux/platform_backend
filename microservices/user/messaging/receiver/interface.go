package receiver

import (
	"context"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue/rabbitmq"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type HandlerFunc = pkgMQ.HandlerFunc

type SqlMethod interface {
	DelReviewer(ctx context.Context, reviewerId int64) (string, string, error)
}

type MessageQueueMethod interface {
	ListenToQueue(exchange, queueName, routeKey string, handler HandlerFunc)
}

type CacheMethod interface {
	DelCredentials(ctx context.Context, username string) error
}

type DispatchInterface interface {
	HandleRequest(msg protoreflect.ProtoMessage, typeName string)
}
