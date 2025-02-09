package service

import (
	"context"
	"log"

	"github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

func (s *Server) Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {
	log.Println("info: login service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.LoginResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.Login(req)
		if err != nil {
			log.Printf("error: login occur fail %v", err)
			return response, nil
		}

		log.Println("info: login occur success")
		return response, nil
	}
}

func (s *Server) GetUser(ctx context.Context, req *generated.GetUserRequest) (*generated.GetUserResponse, error) {
	log.Println("info: get user service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetUserResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetUser(req)
		if err != nil {
			log.Println("error: get user occur fail: ", err)
			return response, nil
		}

		log.Println("info: get user occur success")
		return response, nil
	}
}

func (s *Server) GetFolloweesByTime(ctx context.Context, req *generated.GetFollowRequest) (*generated.GetFollowResponse, error) {
	log.Println("info: GetFolloweesByTime service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetFollowResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetFolloweesByTime(req)
		if err != nil {
			log.Println("error: get user occur fail: ", err)
			return response, nil
		}

		log.Println("info: get user occur success")
		return response, nil
	}
}

func (s *Server) GetFolloweesByViews(ctx context.Context, req *generated.GetFollowRequest) (*generated.GetFollowResponse, error) {
	log.Println("info: GetFolloweesByViews service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetFollowResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetFolloweesByViews(req)
		if err != nil {
			log.Println("error: get user occur fail: ", err)
			return response, nil
		}

		log.Println("info: get user occur success")
		return response, nil
	}
}

func (s *Server) GetFollowers(ctx context.Context, req *generated.GetFollowRequest) (*generated.GetFollowResponse, error) {
	log.Println("info: GetFollowers service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetFollowResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetFollowers(req)
		if err != nil {
			log.Println("error: get user occur fail: ", err)
			return response, nil
		}

		log.Println("info: get user occur success")
		return response, nil
	}
}
