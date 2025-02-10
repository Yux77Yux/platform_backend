package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	internal "github.com/Yux77Yux/platform_backend/microservices/comment/internal"
)

func (s *Server) PublishComment(ctx context.Context, req *generated.PublishCommentRequest) (*generated.PublishCommentResponse, error) {
	log.Println("info: publish comment service start")

	response, err := internal.PublishComment(req)
	if err != nil {
		log.Println("error: upload comment occur fail")
		return response, nil
	}

	log.Println("info: upload comment occur success")
	return response, nil
}
