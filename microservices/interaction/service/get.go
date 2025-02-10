package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	internal "github.com/Yux77Yux/platform_backend/microservices/interaction/internal"
)

func (s *Server) GetActionTag(ctx context.Context, req *generated.GetCreationInteractionRequest) (*generated.GetCreationInteractionResponse, error) {
	log.Println("info: GetActionTag service start")

	response, err := internal.GetActionTag(ctx, req)
	if err != nil {
		log.Println("error: GetActionTag fail: ", err)
		return response, nil
	}

	log.Println("info: GetActionTag success")
	return response, nil
}

func (s *Server) GetCollections(ctx context.Context, req *generated.GetCollectionsRequest) (*generated.GetInteractionsResponse, error) {
	log.Println("info: GetCollections service start")

	response, err := internal.GetCollections(ctx, req)
	if err != nil {
		log.Println("error: GetCollections fail: ", err)
		return response, nil
	}

	log.Println("info: GetCollections success")
	return response, nil
}

func (s *Server) GetHistories(ctx context.Context, req *generated.GetHistoriesRequest) (*generated.GetInteractionsResponse, error) {
	log.Println("info: GetHistories service start")

	response, err := internal.GetHistories(ctx, req)
	if err != nil {
		log.Println("error: GetHistories fail: ", err)
		return response, nil
	}

	log.Println("info: GetHistories success")
	return response, nil
}
func (s *Server) GetRecommend(ctx context.Context, req *generated.GetRecommendRequest) (*generated.GetRecommendResponse, error) {
	log.Println("info: GetRecommend service start")

	response, err := internal.GetRecommend(ctx, req)
	if err != nil {
		log.Println("error: GetHistories fail: ", err)
		return response, nil
	}

	log.Println("info: GetHistories success")
	return response, nil
}
