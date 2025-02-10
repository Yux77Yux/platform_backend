package messaging

import (
	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/review/messaging/dispatch"
)

// 这里拿取请求之后，立马将reviewerId填充至NewReview,使其责任化
func GetReviews(reviewerId int64, reviewType generated.TargetType) ([]*generated.Review, error) {
	typeName := reviewType.String()

	news := GetMsgs(typeName, typeName, typeName, 8)

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
		}

		reviews[i] = review
		go dispatch.HandleRequest(review, dispatch.Update)
	}

	return reviews, nil
}
