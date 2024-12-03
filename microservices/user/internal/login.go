package internal

import (
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	"github.com/Yux77Yux/platform_backend/microservices/user/cache"
	db "github.com/Yux77Yux/platform_backend/microservices/user/repository"
)

func Login(req *generated.LoginRequest) (*generated.LoginResponse, error) {
	user_credentials := req.GetUserCredentials()
	// 检查空值
	if user_credentials.GetUsername() == "" || user_credentials.GetPassword() == "" {
		err := status.Errorf(codes.InvalidArgument, "username and password cannot be empty")
		log.Printf("warning: %v", err)
		return &generated.LoginResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "400",
				Message: "Username and Password cannot be empty",
			},
		}, err
	}

	// 验证密码
	user_id, err := db.UserVerifyInTranscation(user_credentials)
	if err != nil {
		return &generated.LoginResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "Server error occur",
				Details: err.Error(),
			},
		}, err
	}

	if user_id == -1 {
		return &generated.LoginResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "400",
				Message: "Username and Password may not match",
			},
		}, nil
	}

	// 判断redis有无存有
	exist, err := cache.ExistsUserInfo(user_id)
	if err != nil {
		return &generated.LoginResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "cache client error occur",
			},
		}, fmt.Errorf("fail to get user info in redis: %w", err)
	}

	var user_info *generated.User
	if exist {
		// 先从redis取信息
		user_info, err = cache.GetUserInfo(user_id)
		if err != nil {
			return &generated.LoginResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "500",
					Message: "redis client error occur",
				},
			}, fmt.Errorf("fail to get user info in redis: %w", err)
		}
	} else {
		// redis未存有，则从数据库取信息
		user_info, err = db.UserGetInfoInTransaction(user_id)
		if err != nil {
			return &generated.LoginResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "500",
					Message: "mysql client error occur",
				},
			}, fmt.Errorf("fail to get user info in db: %w", err)
		}

		go func() {
			if err := cache.StoreUserInfo(user_info); err != nil {
				log.Println("redis StoreUserInfo failed")
			}
		}()
	}

	return &generated.LoginResponse{
		UserLogin: &generated.UserLogin{
			UserDefault: user_info.GetUserDefault(),
			UserAvator:  user_info.GetUserAvator(),
		},
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}
