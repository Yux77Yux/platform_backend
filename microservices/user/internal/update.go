package internal

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	userMQ "github.com/Yux77Yux/platform_backend/microservices/user/messaging"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/user/messaging/dispatch"
	oss "github.com/Yux77Yux/platform_backend/microservices/user/oss"
	jwt "github.com/Yux77Yux/platform_backend/pkg/jwt"
)

func DelReviewer(req *generated.DelReviewerRequest) (*generated.DelReviewerResponse, error) {
	reqIdCh := make(chan string, 1)
	requestChannel <- func(reqID string) error {
		reqIdCh <- reqID
		log.Printf("info: handling updateUser request with ID: %s\n", reqID)

		err := userMQ.SendMessage("delReviewer", "delReviewer", req)
		if err != nil {
			return fmt.Errorf("err: request_id: %s ,message: %w", reqID, err)
		}

		log.Printf("info: send to messaging queue updateUser request with ID: %s\n", reqID)

		return nil
	}
	reqId := <-reqIdCh

	return &generated.DelReviewerResponse{
		Msg: &common.ApiResponse{
			Status:  common.ApiResponse_SUCCESS, // 正确：使用常量表示枚举值
			Code:    "202",                      // HTTP 状态码，202 请求已接受，但尚未处理，通常用于异步处理
			Details: "DelReviewer processing",   // 更详细的成功信息
			TraceId: reqId,
		},
	}, nil
}

func UpdateUserSpace(req *generated.UpdateUserSpaceRequest) (*generated.UpdateUserResponse, error) {
	space := req.GetUserUpdateSpace()

	accessToken := req.GetAccessToken().GetValue()
	accessClaims, err := jwt.ParseJWT(accessToken)
	if err != nil {
		return &generated.UpdateUserResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "401",
				Message: "Unauthorized",
				Details: err.Error(),
			},
		}, nil
	}

	space.UserDefault.UserId = accessClaims.UserID

	reqIdCh := make(chan string, 1)
	requestChannel <- func(reqID string) error {
		reqIdCh <- reqID
		log.Printf("info: handling updateUser request with ID: %s\n", reqID)

		err = userMQ.SendMessage("updateUserSpace", "updateUserSpace", space)
		if err != nil {
			return fmt.Errorf("err: request_id: %s ,message: %w", reqID, err)
		}

		log.Printf("info: send to messaging queue updateUser request with ID: %s\n", reqID)

		return nil
	}
	reqId := <-reqIdCh

	return &generated.UpdateUserResponse{
		Msg: &common.ApiResponse{
			Status:  common.ApiResponse_SUCCESS, // 正确：使用常量表示枚举值
			Code:    "202",                      // HTTP 状态码，202 请求已接受，但尚未处理，通常用于异步处理
			Message: "OK",                       // 标识完成
			Details: "UpdateUser success",       // 更详细的成功信息
			TraceId: reqId,
		},
	}, nil
}

func UpdateUserAvatar(req *generated.UpdateUserAvatarRequest) (*generated.UpdateUserAvatarResponse, error) {
	accessToken := req.GetAccessToken().GetValue()
	accessClaims, err := jwt.ParseJWT(accessToken)
	if err != nil {
		return &generated.UpdateUserAvatarResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "401",
				Message: "Unauthorized",
				Details: err.Error(),
			},
		}, nil
	}

	// 操作oss上传
	fileBytesStr := req.GetUserUpdateAvatar().GetUserAvatar()
	fileName := fmt.Sprintf("avatar/%v_%v.png", req.GetUserUpdateAvatar().GetUserId(), timestamppb.Now())
	fileBytes, err := base64.StdEncoding.DecodeString(fileBytesStr)
	if err != nil {
		log.Fatal("Error decoding Base64 string: ", err)
	}
	file := bytes.NewReader(fileBytes)
	userAvatar, err := oss.Client.UploadFile(file, fileName)
	if err != nil {
		return &generated.UpdateUserAvatarResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "Falied to upload to oss",
				Details: err.Error(),
			},
		}, nil
	}
	log.Printf("avatar url: %s over", userAvatar)
	userId := accessClaims.UserID

	updateAvatar := &generated.UserUpdateAvatar{
		UserId:     userId,
		UserAvatar: userAvatar,
	}

	go dispatch.HandleRequest(updateAvatar, dispatch.UpdateUserAvatar)
	go dispatch.HandleRequest(updateAvatar, dispatch.UpdateUserAvatarCache)

	return &generated.UpdateUserAvatarResponse{
		Msg: &common.ApiResponse{
			Status:  common.ApiResponse_SUCCESS, // 正确：使用常量表示枚举值
			Code:    "202",                      // HTTP 状态码，202 请求已接受，但尚未处理，通常用于异步处理
			Message: "OK",                       // 标识完成
			Details: "UpdateUser success",       // 更详细的成功信息
		},
		UserUpdateAvatar: updateAvatar,
	}, nil
}

func UpdateUserStatus(req *generated.UpdateUserStatusRequest) (*generated.UpdateUserResponse, error) {
	accessToken := req.GetAccessToken().GetValue()
	accessClaims, err := jwt.ParseJWT(accessToken)
	if err != nil {
		return &generated.UpdateUserResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "401",
				Message: "Unauthorized",
				Details: err.Error(),
			},
		}, nil
	}

	updateStatus := req.GetUserUpdateStatus()
	userId := accessClaims.UserID
	updateStatus.UserId = userId

	go dispatch.HandleRequest(updateStatus, dispatch.UpdateUserStatus)
	go dispatch.HandleRequest(updateStatus, dispatch.UpdateUserStatusCache)

	return &generated.UpdateUserResponse{
		Msg: &common.ApiResponse{
			Status:  common.ApiResponse_SUCCESS, // 正确：使用常量表示枚举值
			Code:    "202",                      // HTTP 状态码，202 请求已接受，但尚未处理，通常用于异步处理
			Message: "OK",                       // 标识完成
			Details: "UpdateUser success",       // 更详细的成功信息
		},
	}, nil
}

func UpdateUserBio(req *generated.UpdateUserBioRequest) (*generated.UpdateUserResponse, error) {
	accessToken := req.GetAccessToken().GetValue()
	accessClaims, err := jwt.ParseJWT(accessToken)
	if err != nil {
		return &generated.UpdateUserResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "401",
				Message: "Unauthorized",
				Details: err.Error(),
			},
		}, nil
	}

	updateBio := req.GetUserUpdateBio()
	userId := accessClaims.UserID
	updateBio.UserId = userId

	go dispatch.HandleRequest(updateBio, dispatch.UpdateUserBio)
	go dispatch.HandleRequest(updateBio, dispatch.UpdateUserBioCache)

	return &generated.UpdateUserResponse{
		Msg: &common.ApiResponse{
			Status:  common.ApiResponse_SUCCESS, // 正确：使用常量表示枚举值
			Code:    "202",                      // HTTP 状态码，202 请求已接受，但尚未处理，通常用于异步处理
			Message: "OK",                       // 标识完成
			Details: "UpdateUser success",       // 更详细的成功信息
		},
	}, nil
}
