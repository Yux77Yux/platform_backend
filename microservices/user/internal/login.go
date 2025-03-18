package internal

import (
	"context"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	tools "github.com/Yux77Yux/platform_backend/microservices/user/tools"
	errMap "github.com/Yux77Yux/platform_backend/pkg/error"
)

func Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {
	response := new(generated.LoginResponse)
	user_credentials := req.GetUserCredentials()
	// 检查空值
	if (user_credentials.GetUsername() == "" && user_credentials.GetUserEmail() == "") || user_credentials.GetPassword() == "" {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "400",
			Message: "Username and Password cannot be empty",
		}
		return response, nil
	}

	// 验证密码
	var (
		user_part_info *generated.UserCredentials
		err            error
	)
	user_part_info, err = cache.GetUserCredentials(ctx, user_credentials)
	if err != nil {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		go tools.LogError(traceId, fullName, err)
	}
	// 通过database获取
	if user_part_info == nil {
		user_part_info, err = db.UserVerifyInTranscation(ctx, user_credentials)
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

		go func(user_part_info *generated.UserCredentials, ctx context.Context) {
			traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
			err = messaging.SendMessage(ctx, EXCHANGE_STORE_CREDENTIAL, KEY_STORE_CREDENTIAL, user_part_info)
			if err != nil {
				tools.LogError(traceId, fullName, err)
			}
		}(user_part_info, ctx)
	}

	var (
		user_info *generated.UserLogin
		result    map[string]string
		fields    = []string{"user_name", "user_avatar"}
		user_id   = user_part_info.GetUserId()
	)

	// 先从redis取信息
	result, err = cache.GetUserInfo(ctx, user_id, fields)
	if err != nil {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		go tools.LogError(traceId, fullName, err)
	}

	if len(result) <= 0 {
		user, err := db.UserGetInfoInTransaction(ctx, user_id)
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
		user_info = &generated.UserLogin{
			UserDefault: user.GetUserDefault(),
			UserAvatar:  user.GetUserAvatar(),
			UserRole:    user.GetUserRole(),
		}
		go func(user *generated.User, ctx context.Context) {
			traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
			err := messaging.SendMessage(ctx, EXCHANGE_STORE_USER, KEY_STORE_USER, user)
			if err != nil {
				tools.LogError(traceId, fullName, err)
			}
		}(user, ctx)
	} else {
		user_info = &generated.UserLogin{
			UserDefault: &common.UserDefault{
				UserId:   user_id,
				UserName: result["user_name"],
			},
			UserAvatar: result["user_avatar"],
			UserRole:   user_part_info.GetUserRole(),
		}

	}

	response.UserLogin = user_info
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}
