package internal

import (
	"context"
	"log"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	cache "github.com/Yux77Yux/platform_backend/microservices/user/cache"
	messaging "github.com/Yux77Yux/platform_backend/microservices/user/messaging"
	tools "github.com/Yux77Yux/platform_backend/microservices/user/tools"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func AddReviewer(ctx context.Context, req *generated.AddReviewerRequest) (*generated.AddReviewerResponse, error) {
	response := new(generated.AddReviewerResponse)
	token := req.GetAccessToken().GetValue()
	pass, _, err := auth.Auth("post", "user_credentials", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Code:   "403",
			Status: common.ApiResponse_ERROR,
		}
		return response, err
	}

	user_credentials := req.GetUserCredentials()
	// 检查
	username := user_credentials.GetUsername()
	if err := tools.CheckStringLength(username, USERNAME_MIN_LENGTH, USERNAME_MAX_LENGTH); err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: err.Error(),
		}
		return response, err
	}
	password := user_credentials.GetPassword()
	if err := tools.CheckStringLength(password, PASSWORD_MIN_LENGTH, PASSWORD_MAX_LENGTH); err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: err.Error(),
		}
		return response, err
	}

	// redis查询账号是否唯一
	exist, err := cache.ExistsUsername(ctx, user_credentials.GetUsername())
	if err != nil {
		log.Printf("error: failed to use redis client: %v", err)
	}
	if exist {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "409",
			Message: "Username already exists",
			Details: "The username you entered is already taken. Please choose a different one.",
		}
		return response, err
	}

	// redis查询邮箱是否存在，是否唯一
	if user_credentials.GetUserEmail() != "" {
		exist, err = cache.ExistsEmail(ctx, user_credentials.GetUserEmail())
		if err != nil {
			log.Printf("error: failed to use redis client: %v", err)
		}
		if exist {
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "409",
				Message: "email already exists",
				Details: "The email you entered is already taken. Please choose a different one.",
			}
			return response, err
		}
	}

	user_credentials.UserRole = generated.UserRole_ADMIN
	err = messaging.SendMessage(ctx, messaging.Register, messaging.Register, user_credentials)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	return &generated.AddReviewerResponse{
		Msg: &common.ApiResponse{
			Status:  common.ApiResponse_SUCCESS,
			Code:    "201",
			Message: "OK",
			Details: "Register success",
		},
	}, nil
}

func Register(ctx context.Context, req *generated.RegisterRequest) (*generated.RegisterResponse, error) {
	response := new(generated.RegisterResponse)
	user_credentials := req.GetUserCredentials()
	// 检查
	username := user_credentials.GetUsername()
	if err := tools.CheckStringLength(username, USERNAME_MIN_LENGTH, USERNAME_MAX_LENGTH); err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: err.Error(),
		}
		return response, err
	}
	password := user_credentials.GetPassword()
	if err := tools.CheckStringLength(password, PASSWORD_MIN_LENGTH, PASSWORD_MAX_LENGTH); err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: err.Error(),
		}
		return response, err
	}

	// redis查询账号是否唯一
	exist, err := cache.ExistsUsername(ctx, user_credentials.GetUsername())
	if err != nil {
		log.Printf("error: failed to use redis client: %v", err)
	}
	if exist {
		log.Printf("info: username already exists")
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "409",
			Message: "Username already exists",
			Details: "Sorry, that username you've entered is unavailable. Please pick a different one.",
		}
		return response, err
	}

	// redis查询邮箱是否存在，是否唯一
	if user_credentials.GetUserEmail() != "" {
		exist, err = cache.ExistsEmail(ctx, user_credentials.GetUserEmail())
		if err != nil {
			log.Printf("error: failed to use redis client: %v", err)
		}
		if exist {
			log.Printf("info: email already exists")
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "409",
				Message: "email already exists",
				Details: "The email you entered is already taken. Please choose a different one.",
			}
			return response, err
		}
	}

	err = messaging.SendMessage(ctx, messaging.Register, messaging.Register, user_credentials)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}
	return response, err
}

func Follow(ctx context.Context, req *generated.FollowRequest) (*generated.FollowResponse, error) {
	response := new(generated.FollowResponse)
	follow := req.GetFollow()
	token := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("post", "user_credentials", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Code:   "403",
			Status: common.ApiResponse_ERROR,
		}
		return response, nil
	}
	follow.FollowerId = userId

	err = messaging.SendMessage(ctx, messaging.Follow, messaging.Follow, follow)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	response.Msg = &common.ApiResponse{
		Status:  common.ApiResponse_SUCCESS,
		Code:    "202",
		Message: "OK",
		Details: "Follow procossing",
	}
	return response, nil
}
