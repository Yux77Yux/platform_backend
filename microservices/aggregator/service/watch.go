package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	internal "github.com/Yux77Yux/platform_backend/microservices/aggregator/internal"
)

func (s *Server) WatchCreation(ctx context.Context, req *generated.WatchCreationRequest) (*generated.WatchCreationResponse, error) {
	log.Println("info: WatchCreation service start")

	response, err := internal.WatchCreation(ctx, req)
	if err != nil {
		log.Printf("error: WatchCreation occur fail: %v", err)
		return response, nil
	}

	log.Println("info: WatchCreation occur success")
	return response, nil
}

func (s *Server) SimilarCreations(ctx context.Context, req *generated.SimilarCreationsRequest) (*generated.GetCardsResponse, error) {
	log.Println("info: SimilarCreations service start")

	response, err := internal.SimilarCreations(ctx, req)
	if err != nil {
		log.Printf("error: SimilarCreations occur fail: %v", err)
		return response, nil
	}

	log.Println("info: SimilarCreations occur success")
	return response, nil
}

func (s *Server) InitialComments(ctx context.Context, req *generated.InitialCommentsRequest) (*generated.InitialCommentsResponse, error) {
	log.Println("info: InitalComments service start")

	response, err := internal.InitialComments(ctx, req)
	if err != nil {
		log.Printf("error: InitalComments occur fail: %v", err)
		return response, nil
	}

	log.Println("info: InitialComments occur success")
	return response, nil
}

func (s *Server) GetTopComments(ctx context.Context, req *generated.GetTopCommentsRequest) (*generated.GetCommentsResponse, error) {
	log.Println("info: GetTopComments service start")

	response, err := internal.GetTopComments(ctx, req)
	if err != nil {
		log.Printf("error: GetTopComments occur fail: %v", err)
		return response, nil
	}

	log.Println("info: GetTopComments occur success")
	return response, nil
}

func (s *Server) GetSecondComments(ctx context.Context, req *generated.GetSecondCommentsRequest) (*generated.GetCommentsResponse, error) {
	log.Println("info: GetSecondComments service start")

	response, err := internal.GetSecondComments(ctx, req)
	if err != nil {
		log.Printf("error: GetSecondComments occur fail: %v", err)
		return response, nil
	}

	log.Println("info: GetSecondComments occur success")
	return response, nil
}
