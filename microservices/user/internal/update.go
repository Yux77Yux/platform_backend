package internal

import (
	// "fmt"
	// "log"

	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	// cache "github.com/Yux77Yux/platform_backend/microservices/user/cache"
	// userMQ "github.com/Yux77Yux/platform_backend/microservices/user/messaging"
	// db "github.com/Yux77Yux/platform_backend/microservices/user/repository"
)

func UpdateUser(req *generated.UpdateUserRequest) (*generated.UpdateUserResponse, error) {
	return &generated.UpdateUserResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

func UpdateUserName(req *generated.UpdateUserNameRequest) (*generated.UpdateUserResponse, error) {
	return &generated.UpdateUserResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

func UpdateUserAvator(req *generated.UpdateUserAvatorRequest) (*generated.UpdateUserResponse, error) {
	return &generated.UpdateUserResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

func UpdateUserBio(req *generated.UpdateUserBioRequest) (*generated.UpdateUserResponse, error) {
	return &generated.UpdateUserResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}

func UpdateUserStatus(req *generated.UpdateUserStatusRequest) (*generated.UpdateUserResponse, error) {
	return &generated.UpdateUserResponse{
		Msg: &common.ApiResponse{
			Status: common.ApiResponse_SUCCESS,
			Code:   "200",
		},
	}, nil
}
