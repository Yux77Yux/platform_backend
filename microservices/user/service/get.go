package service

import (
	"context"
	"log"

	"github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

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
		}, err
	default:
		response, err := internal.GetUser(req)
		if err != nil {
			log.Println("error: get user occur fail: ", err)
			return response, err
		}

		log.Println("info: get user occur success")
		return response, nil
	}
}
