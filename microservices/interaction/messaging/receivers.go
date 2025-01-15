package messaging

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	db "github.com/Yux77Yux/platform_backend/microservices/interaction/repository"
	snow "github.com/Yux77Yux/platform_backend/pkg/snow"
)

func draftinteractionProcessor(msg amqp.Delivery) error {
	interaction_info := new(generated.interactionUpload)
	// 反序列化
	err := proto.Unmarshal(msg.Body, interaction_info)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("register processor error: %w", err)
	}

	interaction := &generated.interaction{
		interactionId: snow.GetId(),
		BaseInfo:      interaction_info,
		UploadTime:    timestamppb.Now(),
	}

	// 写入数据库

	err = db.interactionAddInTransaction(interaction)
	if err != nil {
		log.Printf("db interactionAddInTransaction occur error: %v", err)
	}

	// 写入缓存
	err = cache.interactionAddInCache(&generated.interactionInfo{interaction: interaction})
	if err != nil {
		log.Printf("cache interactionAddInCache occur error: %v", err)
	}

	return nil
}

func pendinginteractionProcessor(msg amqp.Delivery) error {
	interaction_info := new(generated.interactionUpload)
	// 反序列化
	err := proto.Unmarshal(msg.Body, interaction_info)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("register processor error: %w", err)
	}

	interaction := &generated.interaction{
		interactionId: snow.GetId(),
		BaseInfo:      interaction_info,
		UploadTime:    timestamppb.Now(),
	}

	// 写入数据库

	err = db.interactionAddInTransaction(interaction)
	if err != nil {
		log.Printf("db interactionAddInTransaction occur error: %v", err)
	}

	// 写入缓存
	err = cache.interactionAddInCache(&generated.interactionInfo{interaction: interaction})
	if err != nil {
		log.Printf("cache interactionAddInCache occur error: %v", err)
	}

	return nil
}
