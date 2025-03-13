package internal

import (
	"context"
	"fmt"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	messaging "github.com/Yux77Yux/platform_backend/microservices/creation/messaging"
	tools "github.com/Yux77Yux/platform_backend/microservices/creation/tools"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func UpdateCreation(ctx context.Context, req *generated.UpdateCreationRequest) (*generated.UpdateCreationResponse, error) {
	response := new(generated.UpdateCreationResponse)

	pass, user_id, err := auth.Auth("update", "creation", req.GetAccessToken().GetValue())
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_FAILED,
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

	UpdateInfo := req.GetUpdateInfo()

	src := UpdateInfo.GetSrc()
	if !tools.IsValidVideoURL(src) {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: "Video source URL is invalid",
		}
		return response, err
	}
	thumbnail := UpdateInfo.GetThumbnail()
	if !tools.IsValidImageURL(thumbnail) {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: "Image URL is invalid",
		}
		return response, err
	}

	bio := UpdateInfo.GetBio()
	if err := tools.CheckStringLength(bio, BIO_MIN_LENGTH, BIO_MAX_LENGTH); err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: err.Error(),
		}
		return response, err
	}

	title := UpdateInfo.GetTitle()
	if err := tools.CheckStringLength(title, TITLE_MIN_LENGTH, TITLE_MAX_LENGTH); err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: err.Error(),
		}
		return response, err
	}

	UpdateInfo.AuthorId = user_id
	err = messaging.SendMessage(ctx, messaging.UpdateDbCreation, messaging.UpdateDbCreation, UpdateInfo)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}
	return response, nil
}

// 将草稿发布
func UpdateCreationStatus(ctx context.Context, req *generated.UpdateCreationStatusRequest) (*generated.UpdateCreationResponse, error) {
	response := new(generated.UpdateCreationResponse)

	pass, user_id, err := auth.Auth("update", "creation", req.GetAccessToken().GetValue())
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_FAILED,
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

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}

	updateInfo := req.GetUpdateInfo()
	if updateInfo == nil {
		err := fmt.Errorf("error: not entail the request")
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "403",
			Details: err.Error(),
		}
		return response, err
	}

	updateInfo.AuthorId = user_id
	err = messaging.SendMessage(ctx, messaging.UpdateCreationStatus, messaging.UpdateCreationStatus, updateInfo)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
	}
	return response, nil
}
