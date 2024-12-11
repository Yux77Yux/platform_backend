package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	"github.com/Yux77Yux/platform_backend/generated/common"
	internal "github.com/Yux77Yux/platform_backend/microservices/aggregator/internal"
)

func (s *Server) Space(ctx context.Context, req *generated.SpaceRequest) (*generated.SpaceResponse, error) {
	log.Println("info: space service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.SpaceResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.Space(req)
		if err != nil {
			log.Printf("error: space occur fail: %v", err)
			return response, nil
		}

		log.Println("info: space occur success")
		return response, nil
	}
}
