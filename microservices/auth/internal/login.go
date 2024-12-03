package internal

import (
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	generated "github.com/Yux77Yux/platform_backend/generated/auth"
	common "github.com/Yux77Yux/platform_backend/generated/common"
)

func Login(req *generated.LoginRequest) (*generated.LoginResponse, error) {
	
	return &generated.LoginResponse{
		Tokens: &generated.Tokens{
			RefreshToken: &generated.RefreshToken{},
			AccessToken: &generated.AccessToken{},
		},
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}
