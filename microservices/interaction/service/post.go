package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	internal "github.com/Yux77Yux/platform_backend/microservices/interaction/internal"
)

func (s *Server) PostInteraction(ctx context.Context, req *generated.PostInteractionRequest) (*generated.PostInteractionResponse, error) {
	log.Println("info: PostInteraction service start")

	response, err := internal.PostInteraction(req)
	if err != nil {
		log.Println("error: PostInteraction fail: ", err)
		return response, nil
	}

	log.Println("info: PostInteraction success")
	return response, nil
}
