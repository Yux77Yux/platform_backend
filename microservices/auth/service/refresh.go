package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/auth"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	internal "github.com/Yux77Yux/platform_backend/microservices/auth/internal"
)

func (s *Server) Refresh(ctx context.Context, req *generated.RefreshRequest) (*generated.RefreshResponse, error) {
	log.Println("info: auth refresh service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)

		return &generated.RefreshResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
			},
		}, err
	default:
		response, err := internal.Refresh(req)
		if err != nil {
			log.Println("error: auth refresh occur fail")
			return response, err
		}

		log.Println("info: auth refresh occur success")
		return response, nil
	}
}
