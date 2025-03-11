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
	tools "github.com/Yux77Yux/platform_backend/microservices/user/tools"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func AddReviewer(req *generated.AddReviewerRequest) (*generated.AddReviewerResponse, error) {
	token := req.GetAccessToken().GetValue()
	pass, _, err := auth.Auth("post", "user_credentials", token)
	if err != nil {
		return &generated.AddReviewerResponse{
			Msg: &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			},
		}, err
	}
	if !pass {
		return &generated.AddReviewerResponse{
			Msg: &common.ApiResponse{
				Code:    "403",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			},
		}, err
	}

	user_credentials := req.GetUserCredentials()
	// 检查空值
	if user_credentials.GetUsername() == "" || user_credentials.GetPassword() == "" {
		err := status.Errorf(codes.InvalidArgument, "username and password cannot be empty")
		log.Printf("warning: %v", err)
		return &generated.AddReviewerResponse{
			Msg: &common.ApiResponse{},
		}, err
	}

	// redis查询账号是否唯一
	exist, err := cache.ExistsUsername(user_credentials.GetUsername())
	if err != nil {
		log.Printf("error: failed to use redis client: %v", err)
	}
	if exist {
		log.Printf("info: username already exists")
		return &generated.AddReviewerResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "409",                                                                       // HTTP 状态码，409 表示冲突，用户名已存在
				Message: "Username already exists",                                                   // 用户友好的消息
				Details: "The username you entered is already taken. Please choose a different one.", // 更详细的错误信息
			},
		}, nil
	}

	// redis查询邮箱是否存在，是否唯一
	if user_credentials.GetUserEmail() != "" {
		exist, err = cache.ExistsEmail(user_credentials.GetUserEmail())
		if err != nil {
			log.Printf("error: failed to use redis client: %v", err)
		}
		if exist {
			log.Printf("info: email already exists")
			return &generated.AddReviewerResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "409",                                                                    // HTTP 状态码，409 表示冲突，用户名已存在
					Message: "email already exists",                                                   // 用户友好的消息
					Details: "The email you entered is already taken. Please choose a different one.", // 更详细的错误信息
				},
			}, nil
		}
	}

	user_credentials.UserRole = generated.UserRole_ADMIN
	reqIdCh := make(chan string, 1)
	requestChannel <- func(reqID string) error {
		reqIdCh <- reqID

		err = userMQ.SendMessage(userMQ.Register, userMQ.Register, user_credentials)
		if err != nil {
			return fmt.Errorf("err: request_id: %s ,message: %w", reqID, err)
		}
		return nil
	}
	reqId := <-reqIdCh

	return &generated.AddReviewerResponse{
		Msg: &common.ApiResponse{
			Status:  common.ApiResponse_SUCCESS, // 正确：使用常量表示枚举值
			Code:    "201",                      // HTTP 状态码，201 表示创建成功
			Message: "OK",                       // 标识完成
			Details: "Register success",         // 更详细的成功信息
			TraceId: reqId,
		},
	}, nil
}

func Register(req *generated.RegisterRequest) (*generated.RegisterResponse, error) {
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
	exist, err := cache.ExistsUsername(user_credentials.GetUsername())
	if err != nil {
		log.Printf("error: failed to use redis client: %v", err)
	}
	if exist {
		log.Printf("info: username already exists")
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "409",                                                                       // HTTP 状态码，409 表示冲突，用户名已存在
			Message: "Username already exists",                                                   // 用户友好的消息
			Details: "The username you entered is already taken. Please choose a different one.", // 更详细的错误信息
		}
		return response, err
	}

	// redis查询邮箱是否存在，是否唯一
	if user_credentials.GetUserEmail() != "" {
		exist, err = cache.ExistsEmail(user_credentials.GetUserEmail())
		if err != nil {
			log.Printf("error: failed to use redis client: %v", err)
		}
		if exist {
			log.Printf("info: email already exists")
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "409",                                                                    // HTTP 状态码，409 表示冲突，用户名已存在
				Message: "email already exists",                                                   // 用户友好的消息
				Details: "The email you entered is already taken. Please choose a different one.", // 更详细的错误信息
			}
			return response, err
		}
	}

	reqIdCh := make(chan string, 1)
	requestChannel <- func(reqID string) error {
		reqIdCh <- reqID

		err = userMQ.SendMessage(userMQ.Register, userMQ.Register, user_credentials)
		if err != nil {
			return fmt.Errorf("err: request_id: %s ,message: %w", reqID, err)
		}
		return nil
	}
	reqId := <-reqIdCh

	return &generated.RegisterResponse{
		Msg: &common.ApiResponse{
			Status:  common.ApiResponse_SUCCESS, // 正确：使用常量表示枚举值
			Code:    "202",                      // HTTP 状态码，201 表示创建成功
			Message: "OK",                       // 标识完成
			Details: "Register procossing",      // 更详细的成功信息
			TraceId: reqId,
		},
	}, nil
}

func Follow(req *generated.FollowRequest) (*generated.FollowResponse, error) {
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

	err = userMQ.SendMessage(userMQ.Follow, userMQ.Follow, follow)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	response.Msg = &common.ApiResponse{
		Status:  common.ApiResponse_SUCCESS, // 正确：使用常量表示枚举值
		Code:    "202",                      // HTTP 状态码，201 表示创建成功
		Message: "OK",                       // 标识完成
		Details: "Follow procossing",        // 更详细的成功信息
	}
	return response, nil
}
