package internal

import (
	"context"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
)

func Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {
	response := new(generated.LoginResponse)
	auth_client, err := client.NewAuthClient()
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}
	user_client, err := client.NewUserClient()
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}

	// 登录服务，拿用户信息
	user_response, err := user_client.Login(ctx, req.GetUserCredentials())
	if err != nil {
		var msg *common.ApiResponse
		if user_response != nil {
			msg = user_response.GetMsg()
		} else {
			msg = &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			}
		}
		response.Msg = msg
		return response, err
	}
	if user_response.Msg.GetStatus() != common.ApiResponse_SUCCESS {
		response.Msg = user_response.Msg
		return response, err
	}

	// 传递user_id至Auth Service 生成token并返回
	auth_response, err := auth_client.Login(ctx, user_response.GetUserLogin())
	if err != nil {
		var msg *common.ApiResponse
		if auth_response != nil {
			msg = auth_response.GetMsg()
		} else {
			msg = &common.ApiResponse{
				Code:    "500",
				Status:  common.ApiResponse_ERROR,
				Details: err.Error(),
			}
		}
		response.Msg = msg
		return response, err
	}
	if auth_response.Msg.GetStatus() != common.ApiResponse_SUCCESS {
		response.Msg = auth_response.Msg
		return response, err
	}

	// 组装返回至前端
	return &generated.LoginResponse{
		UserLogin: user_response.GetUserLogin(),
		Tokens:    auth_response.GetTokens(),
		Msg: &common.ApiResponse{
			Status:  common.ApiResponse_SUCCESS,
			Code:    "200",
			Message: "Login success",
		},
	}, nil
}
