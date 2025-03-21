package receiver

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/creation/messaging/dispatch"
)

func storeCreationProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(generated.CreationInfo)
	// 反序列化
	err := msg.UnmarshalTo(req)
	if err != nil {
		return err
	}

	// 写入缓存
	err = cache.CreationAddInCache(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func updateCreationDbProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(generated.CreationUpdated)
	err := msg.UnmarshalTo(req)
	// 反序列化
	if err != nil {
		return err
	}

	// 更新数据库
	err = db.UpdateCreationInTransaction(ctx, req)
	if err != nil {
		return err
	}

	status := req.GetStatus()
	// 不需要通知审核服务则返回
	if status != generated.CreationStatus_PENDING {
		return nil
	}

	creationId := req.GetCreationId()
	return messaging.SendMessage(ctx, EXCHANGE_PEND_CREATION, KEY_PEND_CREATION, &common.CreationId{
		Id: creationId,
	})
}

func updateCreationCacheProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(common.CreationId)
	err := msg.UnmarshalTo(req)
	// 反序列化
	if err != nil {
		return fmt.Errorf("updateCreationCacheProcessor error: %w", err)
	}
	id := req.GetId()
	if id <= 0 {
		return fmt.Errorf("reqId not exist")
	}

	// 更新缓存
	creationInfo, err := db.GetDetailInTransaction(context.Background(), id)
	if err != nil {
		return err
	}

	return messaging.SendMessage(ctx, EXCHANGE_STORE_CREATION, KEY_STORE_CREATION, creationInfo)
}

func updateCreationStatusProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(generated.CreationUpdateStatus)
	// 反序列化
	err := msg.UnmarshalTo(req)
	if err != nil {
		return fmt.Errorf("updateCreationStatusProcessor processor error: %w", err)
	}

	// 更新数据库
	err = db.UpdateCreationStatusInTransaction(ctx, req)
	if err != nil {
		return fmt.Errorf("db UpdateCreationStatusInTransaction occur : %w", err)
	}

	reqId := req.GetCreationId()
	status := req.GetStatus()
	// 已经是发布状态，这里的_PUBLISHED只能由审核员选择
	if status == generated.CreationStatus_PUBLISHED {
		// 更改发布时间
		publishedTime := timestamppb.Now()
		err = db.PublishCreationInTransaction(ctx, reqId, publishedTime)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		// 更改作品的缓存
		err = messaging.SendMessage(ctx, EXCHANGE_UPDATE_CACHE_CREATION, KEY_UPDATE_CACHE_CREATION, &common.CreationId{
			Id: reqId,
		})
		if err != nil {
			return err
		}

		// 将作品加入可观看视频
		err = cache.AddPublicCreations(ctx, reqId)
		if err != nil {
			log.Printf("add public %v", err)
			return err
		}

		// 获取作者id
		authorId, err := db.GetAuthorIdInTransaction(ctx, reqId)
		if err != nil {
			log.Printf("GetAuthorIdInTransaction %v", err)
			return err
		}

		// 将作品id加入空间
		err = cache.AddSpaceCreations(ctx, authorId, reqId, publishedTime)
		if err != nil {
			log.Printf("AddSpaceCreations %v", err)
			return err
		}
	}

	// 作者选择发布，这里的PENDING是作者能控制的
	if status == generated.CreationStatus_PENDING {
		return messaging.SendMessage(ctx, EXCHANGE_PEND_CREATION, KEY_PEND_CREATION, &common.CreationId{
			Id: reqId,
		})
	}

	return nil
}

func deleteCreationProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(generated.CreationUpdateStatus)
	// 反序列化
	err := msg.UnmarshalTo(req)
	if err != nil {
		return fmt.Errorf("deleteCreationProcessor processor error: %w", err)
	}

	// 删除数据库中作品
	err = db.UpdateCreationStatusInTransaction(ctx, req)
	if err != nil {
		return fmt.Errorf("error: deleteCreationProcessor UpdateCreationStatusInTransaction error %w", err)
	}

	// 删除缓存中作品
	err = cache.UpdateCreationStatus(ctx, req)
	if err != nil {
		return fmt.Errorf("error: deleteCreationProcessor UpdateCreationStatus error %w", err)
	}

	return nil
}

// 从aggrator interaction过来
func addInteractionCount(ctx context.Context, msg *anypb.Any) error {
	anyAction := new(common.AnyUserAction)
	// 反序列化
	err := msg.UnmarshalTo(anyAction)
	if err != nil {
		return fmt.Errorf("addInteractionCount processor error: %w", err)
	}

	actions := anyAction.GetActions()
	err = cache.UpdateCreationCount(ctx, actions)
	if err != nil {
		// 入死信，没做
		return err
	}

	for _, action := range actions {
		go dispatcher.HandleRequest(action, dispatch.UpdateCount)
	}

	return nil
}
