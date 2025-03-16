package internal

import (
	"context"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
)

type SqlMethod interface {
	GetComments(ctx context.Context, ids []int32) ([]*generated.Comment, error)
	GetInitialTopCommentsInTransaction(ctx context.Context, creation_id int64) (*generated.CommentArea, []*generated.TopComment, int32, error)
	GetTopCommentsInTransaction(ctx context.Context, creation_id int64, pageNumber int32) ([]*generated.TopComment, error)
	GetSecondCommentsInTransaction(ctx context.Context, creation_id int64, root, pageNumber int32) ([]*generated.SecondComment, error)
	GetReplyCommentsInTransaction(ctx context.Context, user_id int64, page int32) ([]*generated.Comment, error)
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
}

type CacheMethod interface {
}
