package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	"github.com/Yux77Yux/platform_backend/generated/common"
	internal "github.com/Yux77Yux/platform_backend/microservices/aggregator/internal"
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
			},
		}, err
	default:
		response, err := internal.Login(req)
		if err != nil {
			log.Printf("error: login occur fail: %v", err)
			return response, err
		}

		log.Println("info: login occur success")
		return response, nil
	}
}
