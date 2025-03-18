package receiver

import (
	"context"

	"google.golang.org/protobuf/reflect/protoreflect"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue/rabbitmq"
)

type HandlerFunc = pkgMQ.HandlerFunc

type SqlMethod interface {
	UpdateReview(ctx context.Context, review *generated.Review) error
	UpdateReviews(ctx context.Context, reviews []*generated.Review) error
}

type MessageQueueMethod = pkgMQ.MessageQueueMethod

type DispatchInterface interface {
	HandleRequest(msg protoreflect.ProtoMessage, typeName string)
}
