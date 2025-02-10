package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	internal "github.com/Yux77Yux/platform_backend/microservices/creation/internal"
)

func (s *Server) UploadCreation(ctx context.Context, req *generated.UploadCreationRequest) (*generated.UploadCreationResponse, error) {
	log.Println("info: upload creation service start")

	response, err := internal.UploadCreation(req)
	if err != nil {
		log.Println("error: upload creation occur fail")
		return response, nil
	}

	log.Println("info: upload creation occur success")
	return response, nil
}
