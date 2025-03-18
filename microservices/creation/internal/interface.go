package internal

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	"google.golang.org/protobuf/proto"
)

type SqlMethod interface {
	CreationAddInTransaction(ctx context.Context, creation *generated.Creation) error
	GetDetailInTransaction(ctx context.Context, creationId int64) (*generated.CreationInfo, error)
	GetUserCreations(ctx context.Context, req *generated.GetUserCreationsRequest) ([]*generated.CreationInfo, int32, error)
	GetCreationCardInTransaction(ctx context.Context, ids []int64) ([]*generated.CreationInfo, error)
	UpdateCreationStatusInTransaction(ctx context.Context, creation *generated.CreationUpdateStatus) error
	SearchCreations(ctx context.Context, title string, page int32) ([]*generated.CreationInfo, int32, error)
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
}

type CacheMethod interface {
	GetCreationInfo(ctx context.Context, creation_id int64) (*generated.CreationInfo, error)
	GetSpaceCreationList(ctx context.Context, user_id int64, page int32, typeStr string) ([]int64, int32, error)
}
