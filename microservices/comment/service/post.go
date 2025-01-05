package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	"github.com/Yux77Yux/platform_backend/generated/common"
	internal "github.com/Yux77Yux/platform_backend/microservices/comment/internal"
)

func (s *Server) UploadComment(ctx context.Context, req *generated.PublishCommentRequest) (*generated.PublishCommentResponse, error) {
	log.Println("info: upload comment service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.PublishCommentResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.PublishComment(req)
		if err != nil {
			log.Println("error: upload comment occur fail")
			return response, nil
		}

		log.Println("info: upload comment occur success")
		return response, nil
	}
}
