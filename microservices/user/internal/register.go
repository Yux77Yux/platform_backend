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
)

func Register(req *generated.RegisterRequest) (*generated.RegisterResponse, error) {
	user_credentials := req.GetUserCredentials()
	// 检查空值
	if user_credentials.GetUsername() == "" || user_credentials.GetPassword() == "" {
		err := status.Errorf(codes.InvalidArgument, "username and password cannot be empty")
		log.Printf("warning: %v", err)
		return &generated.RegisterResponse{
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
		return &generated.RegisterResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "409",                                                                       // HTTP 状态码，409 表示冲突，用户名已存在
				Message: "Username already exists",                                                   // 用户友好的消息
				Details: "The username you entered is already taken. Please choose a different one.", // 更详细的错误信息
			},
		}, fmt.Errorf("info: username already exists")
	}

	// redis查询邮箱是否存在，是否唯一
	if user_credentials.GetUserEmail() != "" {
		exist, err = cache.ExistsEmail(user_credentials.GetUserEmail())
		if err != nil {
			log.Printf("error: failed to use redis client: %v", err)
		}
		if exist {
			log.Printf("info: email already exists")
			return &generated.RegisterResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "409",                                                                    // HTTP 状态码，409 表示冲突，用户名已存在
					Message: "email already exists",                                                   // 用户友好的消息
					Details: "The email you entered is already taken. Please choose a different one.", // 更详细的错误信息
				},
			}, fmt.Errorf("info: email already exists")
		}
	}

	reqId := make(chan string, 1)
	select {
	// 闭包传递
	case requestChannel <- func(reqID string) error {
		reqId <- reqID
		log.Printf("info: handling request with ID: %s\n", reqID)

		err = userMQ.SendMessage("register", "register", user_credentials)
		if err != nil {
			return fmt.Errorf("err: request_id: %s ,message: %w", reqID, err)
		}
		return nil
	}:
		log.Println("info: handler completely sent to UserRequestChan")
		return &generated.RegisterResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_SUCCESS, // 正确：使用常量表示枚举值
				Code:    "201",                      // HTTP 状态码，201 表示创建成功
				Message: "OK",                       // 标识完成
				Details: "Register success",         // 更详细的成功信息
			},
		}, nil
	default:
		log.Println("warning: userChan is full or closed, registration may not complete")
		return &generated.RegisterResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,                             // 使用 ERROR 而不是 SUCCESS，因为发生了错误
				Code:    "503",                                                // HTTP 状态码 503，表示服务不可用
				Message: "Service Unavailable",                                // 提供用户友好的消息
				Details: "User request queue is full, please try again later", // 更详细的错误信息
				TraceId: <-reqId,                                              // 跟踪码
			},
		}, status.Errorf(codes.Unavailable, "user request queue is full, please try again later")
	}
}
