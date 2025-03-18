package repository

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	pkgDb "github.com/Yux77Yux/platform_backend/pkg/database"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SqlInterface = pkgDb.SqlInterface

type SqlMethod interface {
	CreationAddInTransaction(ctx context.Context, creation *generated.Creation) error
	GetDetailInTransaction(ctx context.Context, creationId int64) (*generated.CreationInfo, error)
	GetAuthorIdInTransaction(ctx context.Context, creationId int64) (int64, error)
	GetUserCreations(ctx context.Context, req *generated.GetUserCreationsRequest) ([]*generated.CreationInfo, int32, error)
	GetCreationCardInTransaction(ctx context.Context, ids []int64) ([]*generated.CreationInfo, error)
	DeleteCreationInTransaction(ctx context.Context, id int64) error
	UpdateViewsInTransaction(ctx context.Context, creationId int64, changingNum int) error
	UpdateLikesInTransaction(ctx context.Context, creationId int64, changingNum int) error
	UpdateSavesInTransaction(ctx context.Context, creationId int64, changingNum int) error
	UpdateCreationInTransaction(ctx context.Context, creation *generated.CreationUpdated) error
	UpdateCreationStatusInTransaction(ctx context.Context, creation *generated.CreationUpdateStatus) error
	PublishCreationInTransaction(ctx context.Context, creationId int64, publishTime *timestamppb.Timestamp) error
	UpdateCreationCount(ctx context.Context, creationId int64, saveCount, likeCount, viewCount int32) error
	SearchCreations(ctx context.Context, title string, page int32) ([]*generated.CreationInfo, int32, error)
}
