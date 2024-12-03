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
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	cache "github.com/Yux77Yux/platform_backend/microservices/user/cache"
	db "github.com/Yux77Yux/platform_backend/microservices/user/repository"
)

func RegisterProcessor(msg amqp.Delivery) error {
	credentials := &generated.UserCredentials{}

	// 反序列化
	err := proto.Unmarshal(msg.Body, credentials)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("register processor error: %w", err)
	}

	// 生成id
	node, err := snowflake.NewNode(1) // 传入机器ID，这里假设为1
	if err != nil {
		log.Printf("Failed to create snowflake node: %v", err)
	}

	id := node.Generate().Int64()
	// 生成唯一的ID,确保不为0
	for id == 0 {
		id = node.Generate().Int64()
	}

	// 使用反序列化后的 credentials
	// 写入数据库注册表
	err = db.UserRegisterInTransaction(credentials, id)
	if err != nil {
		return fmt.Errorf("register in db error occur: %w", err)
	}

	user_info := &generated.User{
		UserDefault: &common.UserDefault{
			UserId:   id,
			UserName: "",
		},
		UserAvator:    "",
		UserBio:       "",
		UserStatus:    generated.User_INACTIVE,
		UserGender:    generated.User_UNDEFINED,
		UserEmail:     credentials.GetEmail(),
		UserBday:      nil,
		UserUpdatedAt: timestamppb.Now(),
		UserCreatedAt: timestamppb.Now(),
	}

	// 写入数据库用户表
	err = db.UserAddInfoInTransaction(user_info)
	if err != nil {
		log.Printf("db methods UserAddInfoInTransaction occur error: %v", err)
	}

	// 写入缓存
	err = cache.StoreUserInfo(user_info)
	if err != nil {
		log.Printf("cache methods StoreUserInfo occur error: %v", err)
	}

	if err := cache.StoreUsername(credentials.GetUsername()); err != nil {
		log.Printf("redis methods StoreUsername occur error: %v", err)
	}

	if err := cache.StoreEmail(credentials.GetEmail()); err != nil {
		log.Printf("redis methods StoreUsername occur error: %v", err)
	}

	log.Println("RegisterProcessor success")
	return nil
}
