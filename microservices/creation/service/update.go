package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	internal "github.com/Yux77Yux/platform_backend/microservices/creation/internal"
)

func (s *Server) UpdateCreation(ctx context.Context, req *generated.UpdateCreationRequest) (*generated.UpdateCreationResponse, error) {
	log.Println("info: UpdateCreation service start")

	response, err := internal.UpdateCreation(req)
	if err != nil {
		log.Println("error: UpdateCreation occur fail: ", err)
		return response, nil
	}

	log.Println("info: UpdateCreation occur success")
	return response, nil
}

func (s *Server) UpdateCreationStatus(ctx context.Context, req *generated.UpdateCreationStatusRequest) (*generated.UpdateCreationResponse, error) {
	log.Println("info: UpdateCreationStatus service start")

	response, err := internal.UpdateCreationStatus(req)
	if err != nil {
		log.Println("error: UpdateCreationStatus occur fail: ", err)
		return response, nil
	}

	log.Println("info: UpdateCreationStatus occur success")
	return response, nil
}
