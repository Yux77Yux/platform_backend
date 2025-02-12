package internal

import (
	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	messaging "github.com/Yux77Yux/platform_backend/microservices/creation/messaging"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func UpdateCreation(req *generated.UpdateCreationRequest) (*generated.UpdateCreationResponse, error) {
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
	UpdateInfo.AuthorId = user_id
	err = messaging.SendMessage(messaging.UpdateCreation, messaging.UpdateCreation, UpdateInfo)
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

func UpdateCreationStatus(req *generated.UpdateCreationStatusRequest) (*generated.UpdateCreationResponse, error) {
	response := new(generated.UpdateCreationResponse)

	pass, isADMIN, user_id, err := auth.AuthRole("update", "creation", req.GetAccessToken().GetValue())
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

	if isADMIN {
		// 如果是则不用加userId，这里已经验证为审核员
		UpdateInfo := req.GetUpdateInfo()
		UpdateInfo.AuthorId = -403
		err = messaging.SendMessage(messaging.UpdateCreationStatus, messaging.UpdateCreationStatus, UpdateInfo)
		if err != nil {
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Details: err.Error(),
			}
			return response, nil
		}
		return response, nil
	}

	// 如果不是审核员，则一定加user_id
	UpdateInfo := req.GetUpdateInfo()
	UpdateInfo.AuthorId = user_id // 此处的user_id为token中，难被更改
	err = messaging.SendMessage(messaging.UpdateCreation, messaging.UpdateCreation, UpdateInfo)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, nil
	}
	return response, nil
}
