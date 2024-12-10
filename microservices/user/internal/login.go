package internal

import (
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	cache "github.com/Yux77Yux/platform_backend/microservices/user/cache"
	userMQ "github.com/Yux77Yux/platform_backend/microservices/user/messaging"
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
				Details: err.Error(),
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
				Details: err.Error(),
			},
		}, fmt.Errorf("fail to get user info in redis: %w", err)
	}

	var user_info *generated.UserLogin
	fields := []string{"user_name", "user_avator", "user_role"}
	if exist {
		// 先从redis取信息
		result, err := cache.GetUserInfo(user_id, fields)
		if err != nil {
			return &generated.LoginResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "500",
					Message: "redis client error occur",
					Details: err.Error(),
				},
			}, fmt.Errorf("fail to get user info in redis: %w", err)
		}

		user_info = &generated.UserLogin{
			UserDefault: &common.UserDefault{
				UserId:   user_id,
				UserName: result["user_name"],
			},
			UserAvator: result["user_avator"],
			UserRole:   generated.UserRole(generated.UserRole_value[result["user_role"]]),
		}
	} else {
		// redis未存有，则从数据库取信息
		result, err := db.UserGetInfoInTransaction(user_id, nil)
		if err != nil {
			return &generated.LoginResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "500",
					Message: "mysql client error occur",
					Details: err.Error(),
				},
			}, fmt.Errorf("fail to get user info in db: %w", err)
		}

		user_info = &generated.UserLogin{
			UserDefault: &common.UserDefault{
				UserId:   result["user_id"].(int64),
				UserName: result["user_name"].(string),
			},
			UserAvator: result["user_avator"].(string),
			UserRole:   generated.UserRole(generated.UserRole_value[result["user_role"].(string)]),
		}

		go userMQ.SendMessage("storeUserInCache_exchange", "storeUserInCache_route", user_info)
	}

	return &generated.LoginResponse{
		UserLogin: user_info,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}
