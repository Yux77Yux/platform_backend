package repository

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
	pkgDb "github.com/Yux77Yux/platform_backend/pkg/database"
)

type SqlInterface = pkgDb.SqlInterface

type SqlMethod interface {
	GetReviews(ctx context.Context, reviewId int64, reviewType generated.TargetType, status generated.ReviewStatus, page int32) ([]*generated.Review, int32, error)
	GetTarget(ctx context.Context, id int64) (int64, *generated.TargetType, error)
	PostReviews(ctx context.Context, reviews []*generated.NewReview) error
	UpdateReviews(ctx context.Context, reviews []*generated.Review) error
	UpdateReview(ctx context.Context, review *generated.Review) error
}
