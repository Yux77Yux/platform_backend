package internal

import (
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
)

func Login(req *generated.LoginRequest) (*generated.LoginResponse, error) {
	auth_client, err := client.NewAuthClient()
	if err != nil {
		return nil, fmt.Errorf("error: auth client %v", err)
	}
	user_client, err := client.NewUserClient()
	if err != nil {
		return nil, fmt.Errorf("error: user client %v", err)
	}

	user_response, err := user_client.Login(req.GetUserCredentials())
	if err != nil {
		return nil, fmt.Errorf("error: user client %v", err)
	}

	msg := user_response.GetMsg()
	if msg.GetStatus() == common.ApiResponse_ERROR || msg.GetStatus() == common.ApiResponse_FAILED {
		return &generated.LoginResponse{
			Msg: msg,
		}, fmt.Errorf("error: error of other cause in user service")
	}

	auth_response, err := auth_client.Login(user_response.GetUserLogin().GetUserDefault().GetUserUuid())
	if err != nil {
		return nil, fmt.Errorf("error: auth client %v", err)
	}

	msg = auth_response.GetMsg()
	if msg.GetStatus() == common.ApiResponse_ERROR || msg.GetStatus() == common.ApiResponse_FAILED {
		return &generated.LoginResponse{
			Msg: msg,
		}, fmt.Errorf("error: error of other cause in auth service")
	}

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
