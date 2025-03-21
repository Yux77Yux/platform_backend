package internal

import (
	"context"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	tools "github.com/Yux77Yux/platform_backend/microservices/user/tools"
	errMap "github.com/Yux77Yux/platform_backend/pkg/error"
)

func GetUser(ctx context.Context, req *generated.GetUserRequest) (*generated.GetUserResponse, error) {
	response := new(generated.GetUserResponse)
	user_id := req.GetUserId()

	result, err := cache.GetUserInfo(ctx, user_id, nil)
	if err != nil {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		go tools.LogError(traceId, fullName, err)
	}
	if len(result) > 0 {
		user_info, err := tools.MapUserByString(result)
		if err != nil {
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Details: err.Error(),
			}
			return response, err
		}
		user_info.UserDefault.UserId = user_id
		response.User = user_info
		response.Msg = &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		}
		return response, nil
	}

	user_info, err := db.UserGetInfoInTransaction(ctx, user_id)
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
	go func(user_info *generated.User, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err := messaging.SendMessage(ctx, EXCHANGE_STORE_USER, KEY_STORE_USER, user_info)
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(user_info, ctx)

	user_info.UserDefault.UserId = user_id

	response.User = user_info
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetFolloweesByTime(ctx context.Context, req *generated.GetFollowRequest) (*generated.GetFollowResponse, error) {
	response := new(generated.GetFollowResponse)

	master := false
	userId := req.GetUserId()
	page := req.GetPage()

	cards, err := db.GetFolloweesByTime(ctx, userId, page)
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

	response.Master = master
	response.Users = cards
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetFolloweesByViews(ctx context.Context, req *generated.GetFollowRequest) (*generated.GetFollowResponse, error) {
	response := new(generated.GetFollowResponse)

	master := false
	userId := req.GetUserId()
	page := req.GetPage()

	cards, err := db.GetFolloweesByViews(ctx, userId, page)
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

	response.Master = master
	response.Users = cards
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetFollowers(ctx context.Context, req *generated.GetFollowRequest) (*generated.GetFollowResponse, error) {
	response := new(generated.GetFollowResponse)

	master := false
	userId := req.GetUserId()
	page := req.GetPage()

	cards, err := db.GetFolloweers(ctx, userId, page)
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

	response.Master = master
	response.Users = cards
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func ExistFollowee(ctx context.Context, req *generated.ExistFolloweeRequest) (*generated.ExistFolloweeResponse, error) {
	response := new(generated.ExistFolloweeResponse)
	follow := req.GetFollow()
	exist, err := db.ExistsFollowee(ctx, follow.GetFollowerId(), follow.GetFolloweeId())
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

	response.Exist = exist
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetUsers(ctx context.Context, req *generated.GetUsersRequest) (*generated.GetUsersResponse, error) {
	response := new(generated.GetUsersResponse)
	ids := req.GetIds()

	users, err := db.GetUsers(ctx, ids)
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

	return &generated.GetUsersResponse{
		Users: users,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}
