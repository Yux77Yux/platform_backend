package internal

import (
	"context"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
)

type SqlMethod interface {
	GetReviews(ctx context.Context, reviewId int64, reviewType generated.TargetType, status generated.ReviewStatus, page int32) ([]*generated.Review, int32, error)
	GetTarget(ctx context.Context, id int64) (int64, *generated.TargetType, error)
}

type MessageQueueMethod interface {
	SendMessage(ctx context.Context, exchange string, routeKey string, req proto.Message) error
	GetMsgs(ctx context.Context, exchange, queueName, routeKey string, count int) [][]byte
}
