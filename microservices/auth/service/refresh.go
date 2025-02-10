package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/auth"
	internal "github.com/Yux77Yux/platform_backend/microservices/auth/internal"
)

func (s *Server) Refresh(ctx context.Context, req *generated.RefreshRequest) (*generated.RefreshResponse, error) {
	log.Println("info: auth refresh service start")

	response, err := internal.Refresh(req)
	if err != nil {
		log.Printf("error: auth refresh occur fail: %v", err)
		return response, nil
	}

	log.Println("info: auth refresh occur success")
	return response, nil
}
