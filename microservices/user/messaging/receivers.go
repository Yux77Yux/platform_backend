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
	dispatch "github.com/Yux77Yux/platform_backend/microservices/user/messaging/dispatch"
	db "github.com/Yux77Yux/platform_backend/microservices/user/repository"
	tools "github.com/Yux77Yux/platform_backend/microservices/user/tools"
)

// 补缓存
func storeUserProcessor(msg amqp.Delivery) error {
	user_info := new(generated.User)
	// 反序列化
	err := proto.Unmarshal(msg.Body, user_info)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("register processor error: %w", err)
	}

	// 写入缓存
	go dispatch.HandleRequest(user_info, dispatch.InsertUserCache)

	return nil
}

// 补缓存
func storeCredentialsProcessor(msg amqp.Delivery) error {
	credentials := new(generated.UserCredentials)

	// 反序列化
	err := proto.Unmarshal(msg.Body, credentials)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return fmt.Errorf("register processor error: %w", err)
	}

	go dispatch.HandleRequest(credentials, dispatch.RegisterCache)
	return nil
}

func registerProcessor(msg amqp.Delivery) error {
	credentials := new(generated.UserCredentials)

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
		return fmt.Errorf("failed to create snowflake node: %w", err)
	}

	id := node.Generate().Int64()
	// 生成唯一的ID,确保不为0
	for id == 0 {
		id = node.Generate().Int64()
	}

	credentials.UserId = id
	go dispatch.HandleRequest(credentials, dispatch.Register)
	go dispatch.HandleRequest(credentials, dispatch.RegisterCache)

	user_info := &generated.User{
		UserDefault: &common.UserDefault{
			UserId: id,
		},
		UserStatus:    generated.UserStatus_INACTIVE,
		UserGender:    generated.UserGender_UNDEFINED,
		UserRole:      credentials.GetUserRole(),
		UserUpdatedAt: timestamppb.Now(),
		UserCreatedAt: timestamppb.Now(),
	}

	go dispatch.HandleRequest(user_info, dispatch.InsertUser)
	go dispatch.HandleRequest(user_info, dispatch.InsertUserCache)

	return nil
}

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

func getUserProcessor(msg amqp.Delivery) (proto.Message, error) {
	req := new(generated.GetUserRequest)
	// 反序列化
	err := proto.Unmarshal(msg.Body, req)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return &generated.GetUserResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "cache client error occur",
				Details: err.Error(),
			},
		}, fmt.Errorf("register processor error: %w", err)
	}

	user_id := req.GetUserId()
	// 用于后来的黑名单,尚未开发
	// accessToken := req.GetAccessToken()
	block := false

	// 判断redis有无存有
	exist, err := cache.ExistsUserInfo(user_id)
	if err != nil {
		return &generated.GetUserResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "cache client error occur",
				Details: err.Error(),
			},
		}, fmt.Errorf("fail to get user info in redis: %w", err)
	}

	var user_info *generated.User
	if exist {
		// 先从redis取信息
		result, err := cache.GetUserInfo(user_id, nil)
		if err != nil {
			return &generated.GetUserResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "500",
					Message: "redis client error occur",
					Details: err.Error(),
				},
			}, fmt.Errorf("fail to get user info in redis: %w", err)
		}

		// 调用函数，传递转换后的 map
		user_info = tools.MapUserByString(result)
	} else {
		// redis未存有，则从数据库取信息
		result, err := db.UserGetInfoInTransaction(user_id, nil)
		if err != nil {
			return &generated.GetUserResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "500",
					Message: "mysql client error occur",
					Details: err.Error(),
				},
			}, fmt.Errorf("fail to get user info in db: %w", err)
		}

		if result == nil {
			return &generated.GetUserResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "404",
					Message: "user not found",
					Details: "No user found with the given ID",
				},
			}, nil
		}

		user_info = tools.MapUser(result)
		go SendMessage("storeUser", "storeUser", user_info)
	}

	user_info.UserDefault.UserId = user_id

	return &generated.GetUserResponse{
		User:  user_info,
		Block: block,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}
