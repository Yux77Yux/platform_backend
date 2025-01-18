package internal

// import (
// 	"bytes"
// 	"encoding/base64"
// 	"fmt"
// 	"log"

// 	"google.golang.org/protobuf/types/known/timestamppb"

// 	common "github.com/Yux77Yux/platform_backend/generated/common"
// 	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
// 	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
// 	messaging "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"
// 	oss "github.com/Yux77Yux/platform_backend/microservices/interaction/oss"
// 	db "github.com/Yux77Yux/platform_backend/microservices/interaction/repository"
// 	jwt "github.com/Yux77Yux/platform_backend/pkg/jwt"
// )

// func UpdateUser(req *generated.UpdateUserRequest) (*generated.UpdateUserResponse, error) {
// 	space := req.GetUserUpdateSpace()

// 	accessToken := req.GetAccessToken().GetValue()
// 	accessClaims, err := jwt.ParseJWT(accessToken)
// 	if err != nil {
// 		return &generated.UpdateUserResponse{
// 			Msg: &common.ApiResponse{
// 				Status:  common.ApiResponse_ERROR,
// 				Code:    "401",
// 				Message: "Unauthorized",
// 				Details: err.Error(),
// 			},
// 		}, nil
// 	}
// 	token_userId := accessClaims.UserID
// 	if token_userId != space.GetUserDefault().GetUserId() {
// 		return &generated.UpdateUserResponse{
// 			Msg: &common.ApiResponse{
// 				Status:  common.ApiResponse_ERROR,
// 				Code:    "400",
// 				Message: "the userId different between token and request",
// 			},
// 		}, nil
// 	}

// 	reqId := make(chan string, 1)
// 	select {
// 	// 闭包传递
// 	case requestChannel <- func(reqID string) error {
// 		reqId <- reqID
// 		log.Printf("info: handling updateUser request with ID: %s\n", reqID)

// 		err = messaging.SendMessage("updateUser", "updateUser", space)
// 		if err != nil {
// 			return fmt.Errorf("err: request_id: %s ,message: %w", reqID, err)
// 		}

// 		log.Printf("info: send to messaging queue updateUser request with ID: %s\n", reqID)

// 		return nil
// 	}:
// 		log.Println("info: handler completely sent to UserRequestChan")
// 		return &generated.UpdateUserResponse{
// 			Msg: &common.ApiResponse{
// 				Status:  common.ApiResponse_SUCCESS, // 正确：使用常量表示枚举值
// 				Code:    "202",                      // HTTP 状态码，202 请求已接受，但尚未处理，通常用于异步处理
// 				Message: "OK",                       // 标识完成
// 				Details: "UpdateUser success",       // 更详细的成功信息
// 			},
// 		}, nil
// 	default:
// 		log.Println("warning: userChan is full or closed, registration may not complete")
// 		return &generated.UpdateUserResponse{
// 			Msg: &common.ApiResponse{
// 				Status:  common.ApiResponse_ERROR,                             // 使用 ERROR 而不是 SUCCESS，因为发生了错误
// 				Code:    "503",                                                // HTTP 状态码 503，表示服务不可用
// 				Message: "Service Unavailable",                                // 提供用户友好的消息
// 				Details: "User request queue is full, please try again later", // 更详细的错误信息
// 				TraceId: <-reqId,                                              // 跟踪码
// 			},
// 		}, nil
// 	}
// }

// func UpdateUserAvatar(req *generated.UpdateUserAvatarRequest) (*generated.UpdateUserResponse, error) {
// 	accessToken := req.GetAccessToken().GetValue()
// 	accessClaims, err := jwt.ParseJWT(accessToken)
// 	if err != nil {
// 		return &generated.UpdateUserResponse{
// 			Msg: &common.ApiResponse{
// 				Status:  common.ApiResponse_ERROR,
// 				Code:    "401",
// 				Message: "Unauthorized",
// 				Details: err.Error(),
// 			},
// 		}, nil
// 	}
// 	token_userId := accessClaims.UserID
// 	if token_userId != req.GetUserUpdateAvatar().GetUserId() {
// 		return &generated.UpdateUserResponse{
// 			Msg: &common.ApiResponse{
// 				Status:  common.ApiResponse_ERROR,
// 				Code:    "400",
// 				Message: "the userId different between token and request",
// 			},
// 		}, nil
// 	}

// 	// 操作oss上传
// 	fileBytesStr := req.GetUserUpdateAvatar().GetUserAvatar()
// 	fileName := fmt.Sprintf("user_avatar_%v_%v.png", req.GetUserUpdateAvatar().GetUserId(), timestamppb.Now())
// 	fileBytes, err := base64.StdEncoding.DecodeString(fileBytesStr)
// 	if err != nil {
// 		log.Fatal("Error decoding Base64 string: ", err)
// 	}
// 	file := bytes.NewReader(fileBytes)
// 	userAvatar, err := oss.Client.UploadFile(file, fileName)
// 	if err != nil {
// 		return &generated.UpdateUserResponse{
// 			Msg: &common.ApiResponse{
// 				Status:  common.ApiResponse_ERROR,
// 				Code:    "500",
// 				Message: "Falied to upload to oss",
// 				Details: err.Error(),
// 			},
// 		}, nil
// 	}
// 	log.Printf("avatar url: %s over", userAvatar)
// 	updateAvatar := &generated.UserUpdateAvatar{
// 		UserId:     req.GetUserUpdateAvatar().GetUserId(),
// 		UserAvatar: userAvatar,
// 	}

// 	go func() {
// 		db.UserUpdateAvatarInTransaction(updateAvatar)
// 	}()
// 	go func() {
// 		cache.UpdateUserAvatar(updateAvatar)
// 	}()

// 	return &generated.UpdateUserResponse{
// 		Msg: &common.ApiResponse{
// 			Status:  common.ApiResponse_SUCCESS, // 正确：使用常量表示枚举值
// 			Code:    "202",                      // HTTP 状态码，202 请求已接受，但尚未处理，通常用于异步处理
// 			Message: "OK",                       // 标识完成
// 			Details: "UpdateUser success",       // 更详细的成功信息
// 		},
// 	}, nil
// }

// func UpdateUserStatus(req *generated.UpdateUserStatusRequest) (*generated.UpdateUserResponse, error) {
// 	accessToken := req.GetAccessToken().GetValue()
// 	accessClaims, err := jwt.ParseJWT(accessToken)
// 	if err != nil {
// 		return &generated.UpdateUserResponse{
// 			Msg: &common.ApiResponse{
// 				Status:  common.ApiResponse_ERROR,
// 				Code:    "401",
// 				Message: "Unauthorized",
// 				Details: err.Error(),
// 			},
// 		}, nil
// 	}
// 	token_userId := accessClaims.UserID
// 	if token_userId != req.GetUserUpdateStatus().GetUserId() {
// 		return &generated.UpdateUserResponse{
// 			Msg: &common.ApiResponse{
// 				Status:  common.ApiResponse_ERROR,
// 				Code:    "400",
// 				Message: "the userId different between token and request",
// 			},
// 		}, nil
// 	}

// 	updateStatus := req.GetUserUpdateStatus()
// 	go func() {
// 		db.UserUpdateStatusInTransaction(updateStatus)
// 	}()
// 	go func() {
// 		cache.UpdateUserStatus(updateStatus)
// 	}()

// 	return &generated.UpdateUserResponse{
// 		Msg: &common.ApiResponse{
// 			Status:  common.ApiResponse_SUCCESS, // 正确：使用常量表示枚举值
// 			Code:    "202",                      // HTTP 状态码，202 请求已接受，但尚未处理，通常用于异步处理
// 			Message: "OK",                       // 标识完成
// 			Details: "UpdateUser success",       // 更详细的成功信息
// 		},
// 	}, nil
// }
