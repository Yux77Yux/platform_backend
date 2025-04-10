package internal

import (
	"context"
	"fmt"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	tools "github.com/Yux77Yux/platform_backend/microservices/user/tools"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
	errMap "github.com/Yux77Yux/platform_backend/pkg/error"
)

func addCredential(ctx context.Context, user_credentials *generated.UserCredentials) error {
	email := user_credentials.GetUserEmail()
	username := user_credentials.GetUsername()
	_err := fmt.Errorf("error: Please fill in the required fields and ensure they meet the requirements")

	err := tools.CheckStringLength(username, USERNAME_MIN_LENGTH, USERNAME_MAX_LENGTH)
	usernameExist := err == nil
	emailExist := tools.IsValidEmail(email)
	if !usernameExist || !emailExist {
		return _err
	}

	password := user_credentials.GetPassword()
	// 检查密码
	if err := tools.CheckStringLength(password, PASSWORD_MIN_LENGTH, PASSWORD_MAX_LENGTH); err != nil {
		return _err
	}

	// 用户名通过正则
	if usernameExist {
		existErr := fmt.Errorf("error: Sorry, that username you've entered is unavailable. Please pick a different one")
		exist, err := cache.ExistsUsername(ctx, username)
		if err != nil {
			traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
			go tools.LogError(traceId, fullName, err)
		}
		if exist {
			return existErr
		} else {
			exist, err = db.Exists(ctx, false, username)
			if err != nil {
				if errMap.IsServerError(err) {
					return err
				}
				return _err
			}
			if exist {
				return existErr
			}
		}
	}

	// 邮箱通过正则
	if emailExist {
		existErr := fmt.Errorf("error: Sorry, that email you've entered is unavailable. Please pick a different one")
		exist, err := cache.ExistsEmail(ctx, email)
		if err != nil {
			traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
			go tools.LogError(traceId, fullName, err)
		}
		if exist {
			return existErr
		} else {
			exist, err = db.Exists(ctx, true, email)
			if err != nil {
				if errMap.IsServerError(err) {
					return err
				}
				return _err
			}
			if exist {
				return existErr
			}
		}
	}

	// 通过则异步发送
	go func(user_credentials *generated.UserCredentials, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err = messaging.SendMessage(ctx, EXCHANGE_REGISTER, KEY_REGISTER, user_credentials)
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(user_credentials, ctx)

	return nil
}

func AddReviewer(ctx context.Context, req *generated.AddReviewerRequest) (*generated.AddReviewerResponse, error) {
	response := new(generated.AddReviewerResponse)
	token := req.GetAccessToken().GetValue()
	pass, _, err := auth.Auth("post", "user_credentials", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Code:   "403",
			Status: common.ApiResponse_ERROR,
		}
		return response, err
	}

	user_credentials := req.GetUserCredentials()
	err = addCredential(ctx, user_credentials)
	if err != nil {
		if errMap.IsServerError(err) {
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    errMap.GrpcCodeToHTTPStatusString(err),
				Details: err.Error(),
			}
			return response, err
		}
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    errMap.GrpcCodeToHTTPStatusString(err),
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

func Register(ctx context.Context, req *generated.RegisterRequest) (*generated.RegisterResponse, error) {
	response := new(generated.RegisterResponse)
	user_credentials := req.GetUserCredentials()
	err := addCredential(ctx, user_credentials)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Details: err.Error(),
		}
		return response, nil
	}

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "202",
	}
	return response, err
}

func Follow(ctx context.Context, req *generated.FollowRequest) (*generated.FollowResponse, error) {
	response := new(generated.FollowResponse)
	follow := req.GetFollow()
	token := req.GetAccessToken().GetValue()
	pass, userId, err := auth.Auth("post", "user", token)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Code:    "500",
			Status:  common.ApiResponse_ERROR,
			Details: err.Error(),
		}
		return response, err
	}
	if !pass {
		response.Msg = &common.ApiResponse{
			Code:   "403",
			Status: common.ApiResponse_ERROR,
		}
		return response, nil
	}
	follow.FollowerId = userId

	go func(follow *generated.Follow, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err = messaging.SendMessage(ctx, EXCHANGE_FOLLOW, KEY_FOLLOW, follow)
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(follow, ctx)

	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_PENDING,
		Code:   "202",
	}
	return response, nil
}
