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
	if (user_credentials.GetUsername() == "" && user_credentials.GetUserEmail() == "") || user_credentials.GetPassword() == "" {
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
	var user_part_info *generated.UserCredentials
	user_part_info, err := cache.GetUserCredentials(user_credentials)
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
	if user_part_info == nil {
		user_part_info, err = db.UserVerifyInTranscation(user_credentials)
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
	}

	if user_part_info == nil {
		return &generated.LoginResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "400",
				Message: "Username and Password may not match",
			},
		}, nil
	}

	user_id := user_part_info.GetUserId()
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
	fields := []string{"user_name", "user_avatar"}
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
			UserAvatar: result["user_avatar"],
			UserRole:   user_info.GetUserRole(),
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

		if result == nil {
			return &generated.LoginResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "404",
					Message: "user not found",
					Details: "No user found with the given ID",
				},
			}, nil
		}

		user_info = &generated.UserLogin{
			UserDefault: &common.UserDefault{
				UserId:   result["id"].(int64),
				UserName: result["name"].(string),
			},
			UserAvatar: result["avatar"].(string),
			UserRole:   user_part_info.GetUserRole(),
		}

		go userMQ.SendMessage("storeUser", "storeUser", user_info)
	}

	return &generated.LoginResponse{
		UserLogin: user_info,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}
