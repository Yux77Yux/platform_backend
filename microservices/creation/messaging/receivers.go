package messaging

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

	// 更新缓存
	err = cache.UpdateCreation(creation)
	if err != nil {
		log.Printf("cache CreationAddInCache occur error: %v", err)
	}

	return nil
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
		err := SendMessage(StoreCreationInfo, StoreCreationInfo, creation)
		if err != nil {
			log.Printf("error: GetCreation SendMessage %v", err)
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
		return fmt.Errorf("register processor error: %w", err)
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
		err = SendMessage(PendingCreation, PendingCreation, &common.CreationId{
			Id: creation.GetCreationId(),
		})
		log.Printf("error: %v", err)
		return err
	}

	// 已经是发布状态
	if status == generated.CreationStatus_PUBLISHED {
		err = SendMessage(UpdateCacheCreation, UpdateCacheCreation, &common.CreationId{
			Id: creation.GetCreationId(),
		})
		log.Printf("error: %v", err)
		return err
	}

	return nil
}
