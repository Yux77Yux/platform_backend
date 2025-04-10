package internal

import (
	"context"
	"mime/multipart"
	"os"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
)

type SqlMethod interface {
	GetActionTag(ctx context.Context, req *generated.BaseInteraction) (*generated.Interaction, error)
	GetCollections(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error)
	GetHistories(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error)
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
}

type CacheMethod interface {
	GetHistories(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error)
	GetCollections(ctx context.Context, userId int64, page int32) ([]*generated.Interaction, error)
	GetInteraction(ctx context.Context, interaction *generated.BaseInteraction) (*generated.Interaction, error)
	GetRecommendBaseUser(ctx context.Context, id int64) ([]int64, int64, error)
	GetRecommendBaseItem(ctx context.Context, id int64) ([]int64, bool, error)
	DelHistories(ctx context.Context, data []*generated.BaseInteraction) error
	DelCollections(ctx context.Context, data []*generated.BaseInteraction) error
	GetPublicCreations(ctx context.Context, count int) ([]int64, error)

	SetUsingArchive(ctx context.Context, id int64, order string) error
	GetUsingArchive(ctx context.Context, id int64) (string, map[int]bool, error)
	GetArchive(ctx context.Context, id int64, order string) (*os.File, error)
	SetArchive(ctx context.Context, id int64, order string, file multipart.File) error
}
