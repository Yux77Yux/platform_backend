package internal

import (
	"context"
	"log"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
)

// 这里拿取新的审核请求
func GetPendingReviews(ctx context.Context, reviewerId int64, reviewType generated.TargetType) ([]*generated.Review, error) {
	const LIMIT = 8
	var (
		exchange string
		queue    string
		key      string
	)
	switch reviewType {
	case generated.TargetType_COMMENT:
		exchange = EXCHANGE_COMMENT_REVIEW
		queue = QUEUE_COMMENT_REVIEW
		key = KEY_COMMENT_REVIEW
	case generated.TargetType_USER:
		exchange = EXCHANGE_USER_REVIEW
		queue = QUEUE_USER_REVIEW
		key = KEY_USER_REVIEW
	case generated.TargetType_CREATION:
		exchange = EXCHANGE_CREATION_REVIEW
		queue = QUEUE_CREATION_REVIEW
		key = KEY_CREATION_REVIEW
	}

	news := messaging.GetMsgs(exchange, queue, key, LIMIT)

	length := len(news)
	reviews := make([]*generated.Review, length)
	for i, body := range news {
		newReview := new(generated.NewReview)
		err := proto.Unmarshal(body, newReview)
		if err != nil {
			return nil, err
		}

		review := &generated.Review{
			New:        newReview,
			ReviewerId: reviewerId,
			UpdatedAt:  newReview.GetCreatedAt(),
		}

		reviews[i] = review
	}
	go func(reviews []*generated.Review) {
		anyReview := &generated.AnyReview{
			Reviews: reviews,
		}
		err := messaging.SendMessage(ctx, EXCHANGE_BATCH_UPDATE, KEY_BATCH_UPDATE, anyReview)
		if err != nil {
			log.Printf("error: BatchUpdate SendMessage %v", err)
		}
	}(reviews)

	return reviews, nil
}
