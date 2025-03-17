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
	typeName := ""
	switch reviewType {
	case generated.TargetType_COMMENT:
		typeName = EXCHANGE_COMMENT_REVIEW
	case generated.TargetType_USER:
		typeName = EXCHANGE_USER_REVIEW
	case generated.TargetType_CREATION:
		typeName = EXCHANGE_CREATION_REVIEW
	}

	news := messaging.GetMsgs(typeName, typeName, typeName, LIMIT)

	length := len(news)
	reviews := make([]*generated.Review, length)
	for i, val := range news {
		body := val.Body
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
	go func() {
		anyReview := &generated.AnyReview{
			Reviews: reviews,
		}
		err := messaging.SendMessage(ctx, EXCHANGE_BATCH_UPDATE, KEY_BATCH_UPDATE, anyReview)
		if err != nil {
			log.Printf("error: BatchUpdate SendMessage %v", err)
		}
	}()

	return reviews, nil
}
