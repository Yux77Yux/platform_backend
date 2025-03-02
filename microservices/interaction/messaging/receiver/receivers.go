package receiver

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	messaging "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging/dispatch"
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

func updateDbInteraction(msg amqp.Delivery) error {
	req := new(generated.OperateInteraction)
	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("error: Unmarshal %v", err)
		return err
	}

	go dispatch.HandleRequest(req, dispatch.DbInteraction)
	return nil
}

func addViewProcessor(msg amqp.Delivery) error {
	req := new(generated.OperateInteraction)
	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("error: Unmarshal %v", err)
		return err
	}

	go dispatch.HandleRequest(req, dispatch.ViewCache)
	err = messaging.SendMessage(messaging.UpdateDb, messaging.UpdateDb, req)
	if err != nil {
		log.Printf("error: SendMessage %v", err)
		return err
	}
	return nil
}

func addCollectionProcessor(msg amqp.Delivery) error {
	req := new(generated.OperateInteraction)
	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("error: Unmarshal %v", err)
		return err
	}

	go dispatch.HandleRequest(req, dispatch.CollectionCache)
	err = messaging.SendMessage(messaging.UpdateDb, messaging.UpdateDb, req)
	if err != nil {
		log.Printf("error: SendMessage %v", err)
		return err
	}
	return nil
}

func addLikeProcessor(msg amqp.Delivery) error {
	req := new(generated.OperateInteraction)
	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("error: Unmarshal %v", err)
		return err
	}

	go dispatch.HandleRequest(req, dispatch.LikeCache)
	err = messaging.SendMessage(messaging.UpdateDb, messaging.UpdateDb, req)
	if err != nil {
		log.Printf("error: SendMessage %v", err)
		return err
	}
	return nil
}

func cancelLikeProcessor(msg amqp.Delivery) error {
	req := new(generated.OperateInteraction)
	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("error: Unmarshal %v", err)
		return err
	}

	go dispatch.HandleRequest(req.GetBase(), dispatch.CancelLikeCache)
	err = messaging.SendMessage(messaging.UpdateDb, messaging.UpdateDb, req)
	if err != nil {
		log.Printf("error: SendMessage %v", err)
		return err
	}
	return nil
}

func batchUpdateDbProcessor(msg amqp.Delivery) error {
	req := new(generated.AnyOperateInteraction)
	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("error: Unmarshal %v", err)
		return err
	}
	go dispatch.HandleRequest(req, dispatch.DbBatchInteraction)
	return nil
}
