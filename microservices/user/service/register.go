package service

import (
	"context"
	"log"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generatedUser "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

func (s *Server) Register(ctx context.Context, req *generatedUser.RegisterRequest) (*generatedUser.RegisterResponse, error) {
	log.Println("info: register service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generatedUser.RegisterResponse{
			Msg: &common.ApiResponse{},
		}, err
	default:
		response, err := internal.Register(req)
		if err != nil {
			log.Println("error: register occur fail")
			return response, err
		}

		log.Println("info: register occur success")
		return response, nil
	}
}
