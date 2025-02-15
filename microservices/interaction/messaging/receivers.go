package messaging

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	recommend "github.com/Yux77Yux/platform_backend/microservices/interaction/recommend"
)

func computeSimilarProcessor(msg amqp.Delivery) error {
	req := new(common.CreationId)
	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("error: Unmarshal %v", err)
		return err
	}

	id := req.GetId()
	results, err := recommend.RecommendItemBased(id)
	if err != nil {
		log.Printf("error: RecommendItemBased %v", err)
		return err
	}

	err = cache.SetRecommendBaseItem(id, results)
	if err != nil {
		log.Printf("error: cache SetRecommendBaseItem %v", err)
		return err
	}
	return nil
}

func computeUserProcessor(msg amqp.Delivery) error {
	req := new(common.UserDefault)
	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("error: Unmarshal %v", err)
		return err
	}

	id := req.GetUserId()
	results, err := recommend.Recommend(id)
	if err != nil {
		log.Printf("error: RecommendItemBased %v", err)
		return err
	}

	err = cache.SetRecommendBaseUser(id, results)
	if err != nil {
		log.Printf("error: cache SetRecommendBaseUser %v", err)
		return err
	}
	return nil
}
