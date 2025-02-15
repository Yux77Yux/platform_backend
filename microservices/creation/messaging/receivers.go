package messaging

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	cache "github.com/Yux77Yux/platform_backend/microservices/creation/cache"
	db "github.com/Yux77Yux/platform_backend/microservices/creation/repository"
	snow "github.com/Yux77Yux/platform_backend/pkg/snow"
)

func draftCreationProcessor(msg amqp.Delivery) error {
	creation_info := new(generated.CreationUpload)
	// 反序列化
	err := proto.Unmarshal(msg.Body, creation_info)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("register processor error: %w", err)
	}

	creation := &generated.Creation{
		CreationId: snow.GetId(),
		BaseInfo:   creation_info,
		UploadTime: timestamppb.Now(),
	}

	// 写入数据库

	err = db.CreationAddInTransaction(creation)
	if err != nil {
		log.Printf("db CreationAddInTransaction occur error: %v", err)
	}

	// 写入缓存
	err = cache.CreationAddInCache(&generated.CreationInfo{Creation: creation})
	if err != nil {
		log.Printf("cache CreationAddInCache occur error: %v", err)
	}

	return nil
}

func pendingCreationProcessor(msg amqp.Delivery) error {
	creation_info := new(generated.CreationUpload)
	// 反序列化
	err := proto.Unmarshal(msg.Body, creation_info)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("register processor error: %w", err)
	}

	creation := &generated.Creation{
		CreationId: snow.GetId(),
		BaseInfo:   creation_info,
		UploadTime: timestamppb.Now(),
	}

	// 写入数据库

	err = db.CreationAddInTransaction(creation)
	if err != nil {
		log.Printf("db CreationAddInTransaction occur error: %v", err)
	}

	// 写入缓存
	err = cache.CreationAddInCache(&generated.CreationInfo{Creation: creation})
	if err != nil {
		log.Printf("cache CreationAddInCache occur error: %v", err)
	}

	return nil
}

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

func updateCreationProcessor(msg amqp.Delivery) error {
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

	// 更新缓存
	err = cache.UpdateCreationStatus(creation)
	if err != nil {
		log.Printf("cache UpdateCreationStatus occur error: %v", err)
	}

	return nil
}
