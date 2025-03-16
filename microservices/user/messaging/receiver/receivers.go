package receiver

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/user/messaging/dispatch"
	tools "github.com/Yux77Yux/platform_backend/microservices/user/tools"
)

// 补缓存
func storeUserProcessor(ctx context.Context, msg *anypb.Any) error {
	user_info := new(generated.User)

	err := msg.UnmarshalTo(user_info)
	if err != nil {
		return fmt.Errorf("storeUser processor error: %w", err)
	}

	// 写入缓存
	go dispatcher.HandleRequest(user_info, dispatch.InsertUserCache)

	return nil
}

// 补缓存
func storeCredentialsProcessor(ctx context.Context, msg *anypb.Any) error {
	credentials := new(generated.UserCredentials)

	err := msg.UnmarshalTo(credentials)
	if err != nil {
		return fmt.Errorf("storeCredentialsProcessor error: %w", err)
	}

	go dispatcher.HandleRequest(credentials, dispatch.RegisterCache)
	return nil
}

func registerProcessor(ctx context.Context, msg *anypb.Any) error {
	credentials := new(generated.UserCredentials)

	err := msg.UnmarshalTo(credentials)
	if err != nil {
		return fmt.Errorf("registerProcessor error: %w", err)
	}

	id := tools.GetSnowId()

	// 对密码进行加密
	pwd, err := tools.HashPassword(credentials.GetPassword())
	if err != nil {
		return fmt.Errorf("decrypt hash password failed because %w", err)
	}

	credentials.Password = pwd
	credentials.UserId = id

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

	go dispatcher.HandleRequest(user_info, dispatch.InsertUser)
	go dispatcher.HandleRequest(credentials, dispatch.Register)
	return nil
}

func updateUserSpaceProcessor(ctx context.Context, msg *anypb.Any) error {
	user := new(generated.UserUpdateSpace)

	err := msg.UnmarshalTo(user)
	if err != nil {
		return fmt.Errorf("updateUserSpaceProcessor error: %w", err)
	}

	// 更新 数据库用户表
	go dispatcher.HandleRequest(user, dispatch.UpdateUserSpace)
	go dispatcher.HandleRequest(user, dispatch.UpdateUserSpaceCache)

	return nil
}

func updateUserBioProcessor(ctx context.Context, msg *anypb.Any) error {
	user := new(generated.UserUpdateBio)

	err := msg.UnmarshalTo(user)
	if err != nil {
		return fmt.Errorf("updateUserBioProcessor error: %w", err)
	}

	// 更新 数据库用户表
	go dispatcher.HandleRequest(user, dispatch.UpdateUserBio)
	go dispatcher.HandleRequest(user, dispatch.UpdateUserBioCache)

	return nil
}

func updateUserAvatarProcessor(ctx context.Context, msg *anypb.Any) error {
	user := new(generated.UserUpdateAvatar)

	err := msg.UnmarshalTo(user)
	if err != nil {
		return fmt.Errorf("updateUserAvatarProcessor error: %w", err)
	}

	// 更新 数据库用户表
	go dispatcher.HandleRequest(user, dispatch.UpdateUserAvatar)
	go dispatcher.HandleRequest(user, dispatch.UpdateUserAvatarCache)

	return nil
}

func delReviewerProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(common.UserDefault)

	err := msg.UnmarshalTo(req)
	if err != nil {
		return fmt.Errorf("delReviewerProcessor error: %w", err)
	}

	// 删除审核员身份
	username, email, err := db.DelReviewer(ctx, req.GetUserId())
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	err = cache.DelCredentials(ctx, username)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	if email != "" {
		err = cache.DelCredentials(ctx, email)
		if err != nil {
			log.Printf("error: %v", err)
			return err
		}
	}

	return nil
}

func updateUserStatusProcessor(ctx context.Context, msg *anypb.Any) error {
	updateStatus := new(generated.UserUpdateStatus)

	err := msg.UnmarshalTo(updateStatus)
	if err != nil {
		return fmt.Errorf("updateUserStatusProcessor error: %w", err)
	}

	// 更新 数据库用户表
	go dispatcher.HandleRequest(updateStatus, dispatch.UpdateUserStatus)
	go dispatcher.HandleRequest(updateStatus, dispatch.UpdateUserStatusCache)

	return nil
}

func followProcessor(ctx context.Context, msg *anypb.Any) error {
	follow := new(generated.Follow)

	err := msg.UnmarshalTo(follow)
	if err != nil {
		return fmt.Errorf("followProcessor error: %w", err)
	}

	// 更新 数据库用户表
	go dispatcher.HandleRequest(follow, dispatch.Follow)
	go dispatcher.HandleRequest(follow, dispatch.FollowCache)

	return nil
}
