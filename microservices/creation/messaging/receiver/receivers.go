package receiver

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	cache "github.com/Yux77Yux/platform_backend/microservices/creation/cache"
	messaging "github.com/Yux77Yux/platform_backend/microservices/creation/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/creation/repository"
)

func storeCreationProcessor(msg amqp.Delivery) error {
	creation_info := new(generated.CreationInfo)
	// 反序列化
	err := proto.Unmarshal(msg.Body, creation_info)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("register processor error: %w", err)
	}

	// 写入缓存
	err = cache.CreationAddInCache(creation_info)
	if err != nil {
		log.Printf("cache CreationAddInCache occur error: %v", err)
	}

	return nil
}

func updateCreationDbProcessor(msg amqp.Delivery) error {
	creation := new(generated.CreationUpdated)
	// 反序列化
	err := proto.Unmarshal(msg.Body, creation)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("register processor error: %w", err)
	}

	// 更新数据库
	err = db.UpdateCreationInTransaction(creation)
	if err != nil {
		log.Printf("db CreationAddInTransaction occur error: %v", err)
		return err
	}

	creationId := creation.GetCreationId()
	err = messaging.SendMessage(PendingCreation, PendingCreation, &common.CreationId{
		Id: creationId,
	})

	return err
}

func updateCreationCacheProcessor(msg amqp.Delivery) error {
	creationId := new(common.CreationId)
	// 反序列化
	err := proto.Unmarshal(msg.Body, creationId)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("register processor error: %w", err)
	}
	id := creationId.GetId()
	if id <= 0 {
		log.Printf("creationId not exist")
		return fmt.Errorf("creationId not exist")
	}

	// 更新缓存
	creation, err := db.GetDetailInTransaction(context.Background(), id)
	if err != nil {
		log.Printf("error: creation %v", err)
		return err
	}

	go func(creation *generated.CreationInfo) {
		err := messaging.SendMessage(StoreCreationInfo, StoreCreationInfo, creation)
		if err != nil {
			log.Printf("error: GetCreation messaging.SendMessage %v", err)
		}
	}(creation)

	return nil
}

func updateCreationStatusProcessor(msg amqp.Delivery) error {
	creation := new(generated.CreationUpdateStatus)
	// 反序列化
	err := proto.Unmarshal(msg.Body, creation)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("updateCreationStatusProcessor processor error: %w", err)
	}

	// 更新数据库
	err = db.UpdateCreationStatusInTransaction(creation)
	if err != nil {
		log.Printf("db UpdateCreationStatusInTransaction occur error: %v", err)
		return err
	}

	status := creation.GetStatus()
	// 作者想发布
	if status == generated.CreationStatus_PENDING {
		err = messaging.SendMessage(PendingCreation, PendingCreation, &common.CreationId{
			Id: creation.GetCreationId(),
		})
		log.Printf("error: %v", err)
		return err
	}

	// 已经是发布状态
	if status == generated.CreationStatus_PUBLISHED {
		err = messaging.SendMessage(UpdateCacheCreation, UpdateCacheCreation, &common.CreationId{
			Id: creation.GetCreationId(),
		})
		log.Printf("error: %v", err)
		return err
	}

	return nil
}

func deleteCreationProcessor(msg amqp.Delivery) error {
	deleteInfo := new(generated.CreationUpdateStatus)
	// 反序列化
	err := proto.Unmarshal(msg.Body, deleteInfo)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("deleteCreationProcessor processor error: %w", err)
	}

	// 删除数据库中作品
	err = db.UpdateCreationStatusInTransaction(deleteInfo)
	if err != nil {
		return fmt.Errorf("error:deleteCreationProcessor UpdateCreationStatusInTransaction error %w", err)
	}

	// 删除缓存中作品
	err = cache.UpdateCreationStatus(deleteInfo)
	if err != nil {
		return fmt.Errorf("error:deleteCreationProcessor UpdateCreationStatus error %w", err)
	}

	return nil
}

func addInteractionCount(msg amqp.Delivery) error {
	actions := new(common.AnyUserAction)
	// 反序列化
	err := proto.Unmarshal(msg.Body, actions)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("addInteractionCount processor error: %w", err)
	}

	return nil
}
