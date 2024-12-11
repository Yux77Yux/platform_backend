package service

import (
	"context"
	"log"

	"github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

func (s *Server) UpdateUser(ctx context.Context, req *generated.UpdateUserRequest) (*generated.UpdateUserResponse, error) {
	log.Println("info: update user service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.UpdateUserResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, err
	default:
		response, err := internal.UpdateUser(req)
		if err != nil {
			log.Println("error: update user occur fail")
			return response, err
		}

		log.Println("info: update user occur success")
		return response, nil
	}
}

func (s *Server) UpdateUserName(ctx context.Context, req *generated.UpdateUserNameRequest) (*generated.UpdateUserResponse, error) {
	log.Println("info: update user name service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.UpdateUserResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, err
	default:
		response, err := internal.UpdateUserName(req)
		if err != nil {
			log.Println("error: update user name occur fail")
			return response, err
		}

		log.Println("info: update user name occur success")
		return response, nil
	}
}

func (s *Server) UpdateUserAvator(ctx context.Context, req *generated.UpdateUserAvatorRequest) (*generated.UpdateUserResponse, error) {
	log.Println("info: update user avator service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.UpdateUserResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, err
	default:
		response, err := internal.UpdateUserAvator(req)
		if err != nil {
			log.Println("error: update user avator occur fail")
			return response, err
		}

		log.Println("info: update user avator occur success")
		return response, nil
	}
}

func (s *Server) UpdateUserBio(ctx context.Context, req *generated.UpdateUserBioRequest) (*generated.UpdateUserResponse, error) {
	log.Println("info: update user bio service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.UpdateUserResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, err
	default:
		response, err := internal.UpdateUserBio(req)
		if err != nil {
			log.Println("error: update user bio occur fail")
			return response, err
		}

		log.Println("info: update user bio occur success")
		return response, nil
	}
}

func (s *Server) UpdateUserStatus(ctx context.Context, req *generated.UpdateUserStatusRequest) (*generated.UpdateUserResponse, error) {
	log.Println("info: update user status service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.UpdateUserResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.UpdateUserStatus(req)
		if err != nil {
			log.Println("error: update user status occur fail")
			return response, nil
		}

		log.Println("info: update user status occur success")
		return response, nil
	}
}
