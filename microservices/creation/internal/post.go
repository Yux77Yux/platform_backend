package internal

import (
	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	messaging "github.com/Yux77Yux/platform_backend/microservices/creation/messaging"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func UploadCreation(req *generated.UploadCreationRequest) (*generated.UploadCreationResponse, error) {
	pass, author_id, err := auth.Auth("post", "creation", req.GetAccessToken().GetValue())
	if err != nil {
		return &generated.UploadCreationResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_FAILED,
				Code:   "500",
			},
		}, err
	}
	if !pass {
		return &generated.UploadCreationResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_ERROR,
				Code:   "403",
			},
		}, nil
	}
	// 以上为鉴权
	req.BaseInfo.AuthorId = author_id

	// 异步处理
	if req.GetBaseInfo().GetStatus() == generated.CreationStatus_DRAFT {
		messaging.SendMessage(messaging.DraftCreation, messaging.DraftCreation, req.GetBaseInfo())
	} else if req.GetBaseInfo().GetStatus() == generated.CreationStatus_PENDING {
		messaging.SendMessage(messaging.PendingCreation, messaging.PendingCreation, req.GetBaseInfo())
	}

	return &generated.UploadCreationResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_PENDING,
			Code:   "202",
		},
	}, nil
}
