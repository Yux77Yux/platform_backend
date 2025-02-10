package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/auth"
	internal "github.com/Yux77Yux/platform_backend/microservices/auth/internal"
)

func (s *Server) Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {
	log.Println("info: auth login service start")

	response, err := internal.Login(req)
	if err != nil {
		log.Println("error: auth login occur fail")
		return response, nil
	}

	log.Println("info: auth login occur success")
	return response, nil
}
