package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

func (s *Server) AddReviewer(ctx context.Context, req *generated.AddReviewerRequest) (*generated.AddReviewerResponse, error) {
	log.Println("info: AddReviewer service start")

	response, err := internal.AddReviewer(req)
	if err != nil {
		log.Println("error: register occur fail")
		return response, nil
	}

	log.Println("info: register success")
	return response, nil
}

func (s *Server) Follow(ctx context.Context, req *generated.FollowRequest) (*generated.FollowResponse, error) {
	log.Println("info: Follow service start")

	response, err := internal.Follow(req)
	if err != nil {
		log.Println("error: Follow occur fail")
		return response, nil
	}

	log.Println("info: Follow success")
	return response, nil
}

func (s *Server) Register(ctx context.Context, req *generated.RegisterRequest) (*generated.RegisterResponse, error) {
	log.Println("info: register service start")

	response, err := internal.Register(req)
	if err != nil {
		log.Println("error: register occur fail")
		return response, nil
	}

	log.Println("info: register success")
	return response, nil
}
