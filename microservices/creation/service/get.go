package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	internal "github.com/Yux77Yux/platform_backend/microservices/creation/internal"
)

func (s *Server) GetCreation(ctx context.Context, req *generated.GetCreationRequest) (*generated.GetCreationResponse, error) {
	log.Println("info: get creation service start")

	response, err := internal.GetCreation(ctx, req)
	if err != nil {
		log.Println("error: get creation occur fail: ", err)
		return response, nil
	}

	log.Println("info: get creation occur success")
	return response, nil
}

func (s *Server) GetCreationList(ctx context.Context, req *generated.GetCreationListRequest) (*generated.GetCreationListResponse, error) {
	log.Println("info: GetCreationList service start")

	response, err := internal.GetCreationList(ctx, req)
	if err != nil {
		log.Println("error: get creation occur fail: ", err)
		return response, nil
	}

	log.Println("info: get creation occur success")
	return response, nil
}

func (s *Server) GetPublicCreationList(ctx context.Context, req *generated.GetCreationListRequest) (*generated.GetCreationListResponse, error) {
	log.Println("info: GetPublicCreationList service start")

	response, err := internal.GetPublicCreationList(ctx, req)
	if err != nil {
		log.Println("error: get creation occur fail: ", err)
		return response, nil
	}

	log.Println("info: get creation occur success")
	return response, nil
}

func (s *Server) GetSpaceCreations(ctx context.Context, req *generated.GetSpaceCreationsRequest) (*generated.GetCreationListResponse, error) {
	log.Println("info: GetSpaceCreations service start")

	response, err := internal.GetSpaceCreations(ctx, req)
	if err != nil {
		log.Println("error: get creation occur fail: ", err)
		return response, nil
	}

	log.Println("info: get creation occur success")
	return response, nil
}
