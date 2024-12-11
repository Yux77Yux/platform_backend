package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/auth"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	internal "github.com/Yux77Yux/platform_backend/microservices/auth/internal"
)

func (s *Server) Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {
	log.Println("info: auth login service start")

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
			log.Println("error: auth login occur fail")
			return response, nil
		}

		log.Println("info: auth login occur success")
		return response, nil
	}
}
