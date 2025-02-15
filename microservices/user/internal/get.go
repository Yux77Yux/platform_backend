package internal

import (
	"context"
	"fmt"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	cache "github.com/Yux77Yux/platform_backend/microservices/user/cache"
	userMQ "github.com/Yux77Yux/platform_backend/microservices/user/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/user/repository"
	tools "github.com/Yux77Yux/platform_backend/microservices/user/tools"
	jwt "github.com/Yux77Yux/platform_backend/pkg/jwt"
)

func GetUser(ctx context.Context, req *generated.GetUserRequest) (*generated.GetUserResponse, error) {
	user_id := req.GetUserId()

	// 判断redis有无存有
	exist, err := cache.ExistsUserInfo(ctx, user_id)
	if err != nil {
		return &generated.GetUserResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "cache client error occur",
				Details: err.Error(),
			},
		}, fmt.Errorf("fail to get user info in redis: %w", err)
	}

	var user_info *generated.User
	if exist {
		// 先从redis取信息
		result, err := cache.GetUserInfo(ctx, user_id, nil)
		if err != nil {
			return &generated.GetUserResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "500",
					Message: "redis client error occur",
					Details: err.Error(),
				},
			}, fmt.Errorf("fail to get user info in redis: %w", err)
		}

		// 调用函数，传递转换后的 map
		user_info = tools.MapUserByString(result)
	} else {
		// redis未存有，则从数据库取信息
		result, err := db.UserGetInfoInTransaction(ctx, user_id, nil)
		if err != nil {
			return &generated.GetUserResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "500",
					Message: "mysql client error occur",
					Details: err.Error(),
				},
			}, fmt.Errorf("fail to get user info in db: %w", err)
		}

		if result == nil {
			return &generated.GetUserResponse{
				Msg: &common.ApiResponse{
					Status:  common.ApiResponse_ERROR,
					Code:    "404",
					Message: "user not found",
					Details: "No user found with the given ID",
				},
			}, nil
		}

		user_info = tools.MapUser(result)
		go userMQ.SendMessage("storeUser", "storeUser", user_info)
	}

	user_info.UserDefault.UserId = user_id

	return &generated.GetUserResponse{
		User: user_info,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

func GetFolloweesByTime(ctx context.Context, req *generated.GetFollowRequest) (*generated.GetFollowResponse, error) {
	response := new(generated.GetFollowResponse)

	master := false
	userId := req.GetUserId()
	page := req.GetPage()
	accessToken := req.GetAccessToken()
	if accessToken.Value != "none" {
		accessClaims, err := jwt.ParseJWT(accessToken.GetValue())
		if err != nil {
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "access token error",
				Details: err.Error(),
			}
			return response, fmt.Errorf("error: %v", err)
		}

		// 是否为登录用户
		master = accessClaims.UserID == userId
	} else {
		page = 1
	}

	if !master {
		page = 1
	}

	cards, err := db.GetFolloweesByTime(ctx, userId, page)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Message: "database error",
			Details: err.Error(),
		}
		return response, err
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
	accessToken := req.GetAccessToken()
	if accessToken.Value != "none" {
		accessClaims, err := jwt.ParseJWT(accessToken.GetValue())
		if err != nil {
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "access token error",
				Details: err.Error(),
			}
			return response, fmt.Errorf("error: %v", err)
		}

		// 是否为登录用户
		master = accessClaims.UserID == userId
	} else {
		page = 1
	}

	if !master {
		page = 1
	}

	cards, err := db.GetFolloweesByViews(ctx, userId, page)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Message: "database error",
			Details: err.Error(),
		}
		return response, err
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
	accessToken := req.GetAccessToken()
	if accessToken.Value != "none" {
		accessClaims, err := jwt.ParseJWT(accessToken.GetValue())
		if err != nil {
			response.Msg = &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Message: "access token error",
				Details: err.Error(),
			}
			return response, fmt.Errorf("error: %v", err)
		}

		// 是否为登录用户
		master = accessClaims.UserID == userId
	} else {
		page = 1
	}

	if !master {
		page = 1
	}

	cards, err := db.GetFolloweers(ctx, userId, page)
	if err != nil {
		response.Msg = &common.ApiResponse{
			Status:  common.ApiResponse_ERROR,
			Code:    "500",
			Message: "database error",
			Details: err.Error(),
		}
		return response, err
	}

	response.Master = master
	response.Users = cards
	response.Msg = &common.ApiResponse{
		Status: common.ApiResponse_SUCCESS,
		Code:   "200",
	}
	return response, nil
}

func GetUsers(ctx context.Context, req *generated.GetUsersRequest) (*generated.GetUsersResponse, error) {
	ids := req.GetIds()

	users, err := db.GetUsers(ctx, ids)
	if err != nil {
		return &generated.GetUsersResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_ERROR,
				Code:    "500",
				Details: err.Error(),
			},
		}, err
	}

	return &generated.GetUsersResponse{
		Users: users,
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}
