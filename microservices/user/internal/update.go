package internal

import (
	"bytes"
	"context"
	"fmt"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	tools "github.com/Yux77Yux/platform_backend/microservices/creation/tools"
	messaging "github.com/Yux77Yux/platform_backend/microservices/user/messaging"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/user/messaging/dispatch"
	oss "github.com/Yux77Yux/platform_backend/microservices/user/oss"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func DelReviewer(ctx context.Context, req *generated.DelReviewerRequest) (*generated.DelReviewerResponse, error) {
	response := new(generated.DelReviewerResponse)
	err := messaging.SendMessage(ctx, messaging.DelReviewer, messaging.DelReviewer, req)
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
		Details: "DelReviewer processing",
	}
	return response, err
}

func UpdateUserSpace(ctx context.Context, req *generated.UpdateUserSpaceRequest) (*generated.UpdateUserResponse, error) {
	response := new(generated.UpdateUserResponse)
	accessToken := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("update", "user", accessToken)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "403",
		}
		return response, nil
	}

	space := req.GetUserUpdateSpace()
	space.UserDefault.UserId = userId

	err = messaging.SendMessage(ctx, messaging.UpdateUserSpace, messaging.UpdateUserSpace, space)
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
	}
	return response, err
}

func UpdateUserAvatar(ctx context.Context, req *generated.UpdateUserAvatarRequest) (*generated.UpdateUserAvatarResponse, error) {
	response := new(generated.UpdateUserAvatarResponse)
	accessToken := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("update", "user", accessToken)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "403",
		}
		return response, nil
	}

	fileBytesStr := req.GetUserUpdateAvatar().GetUserAvatar()
	if fileBytesStr == "" {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: "avatar not in body",
		}
		return response, nil
	}

	var userAvatar string
	if isUrl := tools.IsValidImageURL(fileBytesStr); !isUrl {
		fileType, fileBytes, err := tools.ParseBase64Image(fileBytesStr)
		if err != nil {
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "400",
				Details: err.Error(),
			}
			return response, nil
		}

		// 操作oss上传
		fileName := fmt.Sprintf("testAvatar/%d.%s", userId, fileType)

		file := bytes.NewReader(fileBytes)
		userAvatar, err = oss.Client.UploadFile(file, fileName)
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
	} else {
		userAvatar = fileBytesStr
	}

	updateAvatar := &generated.UserUpdateAvatar{
		UserId:     userId,
		UserAvatar: userAvatar,
	}

	go dispatch.HandleRequest(updateAvatar, dispatch.UpdateUserAvatar)
	go dispatch.HandleRequest(updateAvatar, dispatch.UpdateUserAvatarCache)

	response.Msg = &common.ApiResponse{
		Status:  common.ApiResponse_SUCCESS,
		Code:    "202",
		Message: "OK",
		Details: "UpdateUser success",
	}
	response.UserUpdateAvatar = updateAvatar
	return response, nil
}

func UpdateUserStatus(ctx context.Context, req *generated.UpdateUserStatusRequest) (*generated.UpdateUserResponse, error) {
	response := new(generated.UpdateUserResponse)
	accessToken := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("update", "user", accessToken)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "403",
		}
		return response, nil
	}

	updateStatus := req.GetUserUpdateStatus()
	updateStatus.UserId = userId

	err = messaging.SendMessage(ctx, messaging.UpdateUserStatus, messaging.UpdateUserStatus, updateStatus)
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
		Details: "UpdateUser success",
	}
	return response, nil
}

func UpdateUserBio(ctx context.Context, req *generated.UpdateUserBioRequest) (*generated.UpdateUserResponse, error) {
	response := new(generated.UpdateUserResponse)
	accessToken := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("update", "user", accessToken)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_ERROR,
			Code:   "403",
		}
		return response, nil
	}

	updateBio := req.GetUserUpdateBio()
	updateBio.UserId = userId

	go dispatch.HandleRequest(updateBio, dispatch.UpdateUserBio)
	go dispatch.HandleRequest(updateBio, dispatch.UpdateUserBioCache)

	response.Msg = &common.ApiResponse{
		Status:  common.ApiResponse_SUCCESS,
		Code:    "202",
		Message: "OK",
		Details: "UpdateUser success",
	}
	return response, nil
}
