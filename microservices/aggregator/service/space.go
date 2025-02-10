package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	internal "github.com/Yux77Yux/platform_backend/microservices/aggregator/internal"
)

func (s *Server) Space(ctx context.Context, req *generated.SpaceRequest) (*generated.SpaceResponse, error) {
	log.Println("info: space service start")

	response, err := internal.Space(ctx, req)
	if err != nil {
		log.Printf("error: space occur fail: %v", err)
		return response, nil
	}

	log.Println("info: space occur success")
	return response, nil
}
