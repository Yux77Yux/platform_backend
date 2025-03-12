package internal

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	generated "github.com/Yux77Yux/platform_backend/generated/auth"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	jwt "github.com/Yux77Yux/platform_backend/pkg/jwt"
)

func Login(req *generated.LoginRequest) (*generated.LoginResponse, error) {
	response := new(generated.LoginResponse)
	userAuth := req.GetUserAuth()
	role := userAuth.GetUserRole()
	userID := userAuth.GetUserId()

	// 生成刷新令牌 Refresh Token（刷新令牌可以没有角色和Scope信息，仅用于刷新）
	refreshToken, err := jwt.GenerateJWT(userID, role.String(), nil)
	if err != nil {
		err = fmt.Errorf("failed to generate refresh token: %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	// 生成访问令牌 Access Token
	accessToken, err := jwt.GenerateJWT(userID, role.String(), RoleScopeMapping[role.String()])
	if err != nil {
		err = fmt.Errorf("failed to generate access token: %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	response.Tokens = &generated.Tokens{
		RefreshToken: &generated.RefreshToken{
			Value:     refreshToken,
			ExpiresAt: timestamppb.New(time.Now().Add(7 * 24 * time.Hour)),
		},
		AccessToken: &common.AccessToken{
			Value:     accessToken,
			ExpiresAt: timestamppb.New(time.Now().Add(30 * time.Minute)),
		},
	}
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}
