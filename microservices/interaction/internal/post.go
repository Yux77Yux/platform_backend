package internal

import (
	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	messaging "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func Uploadinteraction(req *generated.UploadinteractionRequest) (*generated.UploadinteractionResponse, error) {
	pass, user_id, err := auth.Auth("post", "interaction", req.GetAccessToken().GetValue())
	if err != nil {
		return &generated.UploadinteractionResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_FAILED,
				Code:   "500",
			},
		}, err
	}
	if !pass {
		return &generated.UploadinteractionResponse{
			Msg: &common.ApiResponse{
				Status: common.ApiResponse_ERROR,
				Code:   "403",
			},
		}, nil
	}
	// 以上为鉴权
	req.Interaction.UserId = user_id

	// 异步处理
	messaging.SendMessage("pendinginteraction", "pendinginteraction", req.GetBaseInfo())

	return &generated.UploadinteractionResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_PENDING,
			Code:   "202",
		},
	}, nil
}
