package repository

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	pkgDb "github.com/Yux77Yux/platform_backend/pkg/database"
)

type SqlInterface = pkgDb.SqlInterface

type SqlMethod interface {
	BatchInsert(ctx context.Context, comments []*generated.Comment) (int64, error)
	GetInitialTopCommentsInTransaction(ctx context.Context, creation_id int64) (*generated.CommentArea, []*generated.TopComment, int32, error)
	GetTopCommentsInTransaction(ctx context.Context, creation_id int64, pageNumber int32) ([]*generated.TopComment, error)
	GetSecondCommentsInTransaction(ctx context.Context, creation_id int64, root, pageNumber int32) ([]*generated.SecondComment, error)
	GetReplyCommentsInTransaction(ctx context.Context, user_id int64, page int32) ([]*generated.Comment, error)
	GetCommentInfo(ctx context.Context, comments []*common.AfterAuth) ([]*common.AfterAuth, error)
	GetComments(ctx context.Context, ids []int32) ([]*generated.Comment, error)
	GetCreationIdInTransaction(ctx context.Context, comment_id int32) (int64, int64, error)
	BatchUpdateDeleteStatus(ctx context.Context, comments []*common.AfterAuth) (int64, error)
}
