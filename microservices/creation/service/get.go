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

func (s *Server) GetSimilarCreationList(ctx context.Context, req *generated.GetPublicCreationListRequest) (*generated.GetCreationListResponse, error) {
	log.Println("info: GetSimilarCreationList service start")

	response, err := internal.GetSimilarCreationList(ctx, req)
	if err != nil {
		log.Println("error: get creation occur fail: ", err)
		return response, nil
	}

	log.Println("info: get creation occur success")
	return response, nil
}

func (s *Server) GetSpaceCreationList(ctx context.Context, req *generated.GetPublicCreationListRequest) (*generated.GetCreationListResponse, error) {
	log.Println("info: GetSpaceCreationList service start")

	response, err := internal.GetSpaceCreationList(ctx, req)
	if err != nil {
		log.Println("error: get creation occur fail: ", err)
		return response, nil
	}

	log.Println("info: get creation occur success")
	return response, nil
}

func (s *Server) GetCollectionCreationList(ctx context.Context, req *generated.GetSpecificCreationListRequest) (*generated.GetCreationListResponse, error) {
	log.Println("info: GetCollectionCreationList service start")

	response, err := internal.GetCollectionCreationList(ctx, req)
	if err != nil {
		log.Println("error: get creation occur fail: ", err)
		return response, nil
	}

	log.Println("info: get creation occur success")
	return response, nil
}

func (s *Server) GetHomeCreationList(ctx context.Context, req *generated.GetSpecificCreationListRequest) (*generated.GetCreationListResponse, error) {
	log.Println("info: GetHomeCreationList service start")

	response, err := internal.GetHomeCreationList(ctx, req)
	if err != nil {
		log.Println("error: get creation occur fail: ", err)
		return response, nil
	}

	log.Println("info: get creation occur success")
	return response, nil
}
