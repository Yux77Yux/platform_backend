package internal

import (
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
	jwt "github.com/Yux77Yux/platform_backend/pkg/jwt"
)

func Space(req *generated.SpaceRequest) (*generated.SpaceResponse, error) {
	userId := req.GetUserId()
	accessToken := req.GetAccessToken()
	master := false

	accessClaims, err := jwt.ParseJWT(accessToken.GetValue())
	if err != nil {
		return &generated.SpaceResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "user client failed",
				Details: err.Error(),
			},
		}, fmt.Errorf("error: user client %v", err)
	}

	master = accessClaims.UserID == userId

	user_client, err := client.NewUserClient()
	if err != nil {
		return &generated.SpaceResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "user client failed",
				Details: err.Error(),
			},
		}, fmt.Errorf("error: user client NewUserClient %v", err)
	}

	// 用户服务，拿用户信息
	user_response, err := user_client.GetUser(userId, accessToken)
	if err != nil {
		return &generated.SpaceResponse{
			Msg: user_response.GetMsg(),
		}, fmt.Errorf("error: user client GetUser %v", err)
	}

	// 检查是否有错
	msg := user_response.GetMsg()
	if msg.GetStatus() == common.ApiResponse_ERROR || msg.GetStatus() == common.ApiResponse_FAILED {
		return &generated.SpaceResponse{Msg: msg}, fmt.Errorf("error: error of other cause in user service")
	}

	// 组装返回至前端
	return &generated.SpaceResponse{
		User:   user_response.GetUser(),
		Master: master,
		Block:  user_response.GetBlock(),
		Msg: &common.ApiResponse{
			Status:  common.ApiResponse_SUCCESS,
			Code:    "200",
			Message: "Space Request success",
		},
	}, nil
}
