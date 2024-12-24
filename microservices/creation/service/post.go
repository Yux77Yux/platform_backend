package service

import (
	"context"
	"log"

	"github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	internal "github.com/Yux77Yux/platform_backend/microservices/creation/internal"
)

func (s *Server) UploadCreation(ctx context.Context, req *generated.UploadCreationRequest) (*generated.UploadCreationResponse, error) {
	log.Println("info: upload creation service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.UploadCreationResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.UploadCreation(req)
		if err != nil {
			log.Println("error: upload creation occur fail")
			return response, err
		}

		log.Println("info: upload creation occur success")
		return response, nil
	}
}
