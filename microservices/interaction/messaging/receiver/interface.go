package receiver

import (
	"context"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue/rabbitmq"
)

type HandlerFunc = pkgMQ.HandlerFunc

type SqlMethod interface {
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
	ListenToQueue(exchange, queueName, routeKey string, handler HandlerFunc)
}

type CacheMethod interface {
	SetRecommendBaseUser(ctx context.Context, id int64, ids []int64) error
	SetRecommendBaseItem(ctx context.Context, id int64, ids []int64) error
	GetPublicCreations(ctx context.Context, count int) ([]int64, error)
}

type DispatchInterface interface {
	HandleRequest(msg protoreflect.ProtoMessage, typeName string)
}
