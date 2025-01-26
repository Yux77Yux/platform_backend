package messaging

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"fmt"
	"log"

	"github.com/bwmarrin/snowflake"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging/dispatch"
	db "github.com/Yux77Yux/platform_backend/microservices/interaction/repository"
	tools "github.com/Yux77Yux/platform_backend/microservices/interaction/tools"
)

func updateUserSpaceProcessor(msg amqp.Delivery) error {
	log.Println("Update User Info Start !!!")
	user := new(generated.UserUpdateSpace)

	// 反序列化
	err := proto.Unmarshal(msg.Body, user)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("update user processor error: %w", err)
	}

	// 更新 数据库用户表
	go dispatch.HandleRequest(user, dispatch.UpdateUserSpace)
	go dispatch.HandleRequest(user, dispatch.UpdateUserSpaceCache)

	log.Println("UpdateUserProcessor success")
	return nil
}
