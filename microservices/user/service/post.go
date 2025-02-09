package service

import (
	"context"
	"log"

	"github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

func (s *Server) AddReviewer(ctx context.Context, req *generated.AddReviewerRequest) (*generated.AddReviewerResponse, error) {
	log.Println("info: AddReviewer service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.AddReviewerResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.AddReviewer(req)
		if err != nil {
			log.Println("error: register occur fail")
			return response, nil
		}

		log.Println("info: register success")
		return response, nil
	}
}

func (s *Server) Follow(ctx context.Context, req *generated.FollowRequest) (*generated.FollowResponse, error) {
	log.Println("info: Follow service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.FollowResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.Follow(req)
		if err != nil {
			log.Println("error: Follow occur fail")
			return response, nil
		}

		log.Println("info: Follow success")
		return response, nil
	}
}

func (s *Server) Register(ctx context.Context, req *generated.RegisterRequest) (*generated.RegisterResponse, error) {
	log.Println("info: register service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.RegisterResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.Register(req)
		if err != nil {
			log.Println("error: register occur fail")
			return response, nil
		}

		log.Println("info: register success")
		return response, nil
	}
}
