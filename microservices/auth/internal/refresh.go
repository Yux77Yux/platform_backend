package internal

import (
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	generated "github.com/Yux77Yux/platform_backend/generated/auth"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	jwt "github.com/Yux77Yux/platform_backend/pkg/jwt"
)

type ContextKey string

const RefreshTokenKey ContextKey = "refreshToken"

func Refresh(req *generated.RefreshRequest) (*generated.RefreshResponse, error) {
	response := new(generated.RefreshResponse)
	token := req.GetRefreshToken().GetValue()
	// 解密
	refreshToken, err := jwt.VerifyRefreshToken(token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500", // 返回状态码
			Details: err.Error(),
		}
		return response, err
	}
	// 检测refreshToken是否过期或无效
	claims, err := jwt.ParseJWT(refreshToken)
	if err != nil {
		// 检测错误类型
		var statusCode string
		switch {
		case strings.Contains(err.Error(), "unexpected signing method"), strings.Contains(err.Error(), "signature is invalid"):
			statusCode = "403" // 签名方法无效
		case strings.Contains(err.Error(), "token is malformed"):
			statusCode = "400" // 格式错误
		case strings.Contains(err.Error(), "token is expired"):
			statusCode = "401" // Token 已过期
		default:
			statusCode = "500" // 解析失败，未知错误
		}

		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    statusCode, // 返回状态码
			Details: err.Error(),
		}
		if statusCode[0] == '5' {
			return response, err
		}
		return response, nil
	}

	userID := claims.UserID
	role := claims.Role
	if role == "" {
		role = "USER"
	}

	accessToken, err := jwt.GenerateJWT(userID, role, RoleScopeMapping[role])
	if err != nil {
		err = fmt.Errorf("failed to generate access token: %w", err)
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Details: err.Error(),
		}
		return response, err
	}

	response.AccessToken = &common.AccessToken{
		Value:     accessToken,
		ExpiresAt: timestamppb.New(time.Now().Add(30 * time.Minute)),
	}
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}
