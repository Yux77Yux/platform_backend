package service

import (
	"context"
	"log"

	"github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	internal "github.com/Yux77Yux/platform_backend/microservices/interaction/internal"
)

func (s *Server) Uploadinteraction(ctx context.Context, req *generated.UploadinteractionRequest) (*generated.UploadinteractionResponse, error) {
	log.Println("info: upload interaction service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.UploadinteractionResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.Uploadinteraction(req)
		if err != nil {
			log.Println("error: upload interaction occur fail")
			return response, nil
		}

		log.Println("info: upload interaction occur success")
		return response, nil
	}
}
