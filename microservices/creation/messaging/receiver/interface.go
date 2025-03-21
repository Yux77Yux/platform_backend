package receiver

import (
	"context"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue/rabbitmq"
)

type HandlerFunc = pkgMQ.HandlerFunc

type SqlMethod interface {
	GetDetailInTransaction(ctx context.Context, creationId int64) (*generated.CreationInfo, error)
	GetAuthorIdInTransaction(ctx context.Context, creationId int64) (int64, error)
	UpdateCreationInTransaction(ctx context.Context, creation *generated.CreationUpdated) error
	UpdateCreationStatusInTransaction(ctx context.Context, creation *generated.CreationUpdateStatus) error
	PublishCreationInTransaction(ctx context.Context, creationId int64, publishTime *timestamppb.Timestamp) error
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
	ListenToQueue(exchange, queueName, routeKey string, handler HandlerFunc)
}

type CacheMethod interface {
	CreationAddInCache(ctx context.Context, creationInfo *generated.CreationInfo) error
	AddSpaceCreations(ctx context.Context, authorId, creationId int64, publishTime *timestamppb.Timestamp) error
	UpdateCreationStatus(ctx context.Context, creation *generated.CreationUpdateStatus) error
	UpdateCreationCount(ctx context.Context, actions []*common.UserAction) error
	AddPublicCreations(ctx context.Context, creationId int64) error
}

type DispatchInterface interface {
	HandleRequest(msg protoreflect.ProtoMessage, typeName string)
}
