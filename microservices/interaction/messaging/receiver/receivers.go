package receiver

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/anypb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging/dispatch"
	recommend "github.com/Yux77Yux/platform_backend/microservices/interaction/recommend"
)

func computeSimilarProcessor(ctx context.Context, msg *anypb.Any) error {
	const MIN_COUNT = 40
	req := new(common.CreationId)

	err := msg.UnmarshalTo(req)
	if err != nil {
		return err
	}

	id := req.GetId()
	results, err := recommend.RecommendItemBased(ctx, id)
	if err != nil {
		log.Printf("error: RecommendItemBased %v", err)
		return err
	}
	count := len(results)
	if count <= MIN_COUNT {
		ids, err := cache.GetPublicCreations(ctx, MIN_COUNT-count)
		if err != nil {
			return err
		}
		results = append(results, ids...)
	}

	err = cache.SetRecommendBaseItem(ctx, id, results)
	if err != nil {
		return err
	}
	return nil
}

func computeUserProcessor(ctx context.Context, msg *anypb.Any) error {
	const MIN_COUNT = 72
	req := new(common.UserDefault)
	err := msg.UnmarshalTo(req)
	if err != nil {
		return err
	}

	id := req.GetUserId()
	results, err := recommend.Recommend(ctx, id)
	if err != nil {
		return err
	}

	count := len(results)
	if count <= MIN_COUNT {
		ids, err := cache.GetPublicCreations(ctx, MIN_COUNT-count)
		if err != nil {
			return err
		}
		results = append(results, ids...)
	}

	err = cache.SetRecommendBaseUser(ctx, id, results)
	if err != nil {
		return err
	}
	return nil
}

func updateDbInteraction(ctx context.Context, msg *anypb.Any) error {
	req := new(generated.OperateInteraction)
	err := msg.UnmarshalTo(req)
	if err != nil {
		return err
	}

	go dispatcher.HandleRequest(req, dispatch.DbInteraction)
	return nil
}

func addViewProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(generated.OperateInteraction)
	err := msg.UnmarshalTo(req)
	if err != nil {
		return err
	}

	go dispatcher.HandleRequest(req, dispatch.ViewCache)
	err = messaging.SendMessage(ctx, EXCHANGE_UPDATE_DB, KEY_UPDATE_DB, req)
	if err != nil {
		return err
	}
	return nil
}

func addCollectionProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(generated.OperateInteraction)
	err := msg.UnmarshalTo(req)
	if err != nil {
		return err
	}

	go dispatcher.HandleRequest(req, dispatch.CollectionCache)
	err = messaging.SendMessage(ctx, EXCHANGE_UPDATE_DB, KEY_UPDATE_DB, req)
	if err != nil {
		return err
	}
	return nil
}

func addLikeProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(generated.OperateInteraction)
	err := msg.UnmarshalTo(req)
	if err != nil {
		return err
	}

	go dispatcher.HandleRequest(req, dispatch.LikeCache)
	err = messaging.SendMessage(ctx, EXCHANGE_UPDATE_DB, KEY_UPDATE_DB, req)
	if err != nil {
		return err
	}
	return nil
}

func cancelLikeProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(generated.OperateInteraction)
	err := msg.UnmarshalTo(req)
	if err != nil {
		return err
	}

	go dispatcher.HandleRequest(req.GetBase(), dispatch.CancelLikeCache)
	err = messaging.SendMessage(ctx, EXCHANGE_UPDATE_DB, KEY_UPDATE_DB, req)
	if err != nil {
		return err
	}
	return nil
}

func batchUpdateDbProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(generated.AnyOperateInteraction)
	err := msg.UnmarshalTo(req)
	if err != nil {
		return err
	}
	anyOp := req.GetOperateInteractions()
	for _, op := range anyOp {
		go dispatcher.HandleRequest(op, dispatch.DbInteraction)
	}
	return nil
}
