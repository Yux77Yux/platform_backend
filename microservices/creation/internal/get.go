package internal

import (
	// "fmt"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	// cache "github.com/Yux77Yux/platform_backend/microservices/creation/cache"
	// messaging "github.com/Yux77Yux/platform_backend/microservices/creation/messaging"
	// db "github.com/Yux77Yux/platform_backend/microservices/creation/repository"
)

func GetCreation(req *generated.GetCreationRequest) (*generated.GetCreationResponse, error) {
	// creation_id := req.GetcreationId()
	// 用于后来的黑名单,尚未开发
	// accessToken := req.GetAccessToken()
	// block := false

	return &generated.GetCreationResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}
