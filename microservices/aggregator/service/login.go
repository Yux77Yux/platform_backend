package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	internal "github.com/Yux77Yux/platform_backend/microservices/aggregator/internal"
)

func (s *Server) Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {
	log.Println("info: login service start")

	response, err := internal.Login(ctx, req)
	if err != nil {
		log.Printf("error: login occur fail: %v", err)
		return response, nil
	}

	log.Println("info: login occur success")
	return response, nil
}
