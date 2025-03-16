package receiver

import (
	"context"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue/rabbitmq"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type HandlerFunc = pkgMQ.HandlerFunc

type MessageQueueMethod interface {
	ListenToQueue(exchange, queueName, routeKey string, handler HandlerFunc)
}

type DispatchInterface interface {
	HandleRequest(msg protoreflect.ProtoMessage, typeName string)
}

type SqlMethod interface {
	GetCreationIdInTransaction(ctx context.Context, comment_id int32) (int64, int64, error)
}
