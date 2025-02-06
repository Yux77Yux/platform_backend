package service

import (
	"context"
	"log"

	"github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	internal "github.com/Yux77Yux/platform_backend/microservices/interaction/internal"
)

func (s *Server) GetActionTag(ctx context.Context, req *generated.GetCreationInteractionRequest) (*generated.GetCreationInteractionResponse, error) {
	log.Println("info: GetActionTag service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetCreationInteractionResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetActionTag(req)
		if err != nil {
			log.Println("error: GetActionTag fail: ", err)
			return response, nil
		}

		log.Println("info: GetActionTag success")
		return response, nil
	}
}

func (s *Server) GetCollections(ctx context.Context, req *generated.GetCollectionsRequest) (*generated.GetInteractionsResponse, error) {
	log.Println("info: GetCollections service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetInteractionsResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetCollections(req)
		if err != nil {
			log.Println("error: GetCollections fail: ", err)
			return response, nil
		}

		log.Println("info: GetCollections success")
		return response, nil
	}
}

func (s *Server) GetHistories(ctx context.Context, req *generated.GetHistoriesRequest) (*generated.GetInteractionsResponse, error) {
	log.Println("info: GetHistories service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetInteractionsResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetHistories(req)
		if err != nil {
			log.Println("error: GetHistories fail: ", err)
			return response, nil
		}

		log.Println("info: GetHistories success")
		return response, nil
	}
}
func (s *Server) GetRecommend(ctx context.Context, req *generated.GetRecommendRequest) (*generated.GetRecommendResponse, error) {
	log.Println("info: GetRecommend service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetRecommendResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetRecommend(req)
		if err != nil {
			log.Println("error: GetHistories fail: ", err)
			return response, nil
		}

		log.Println("info: GetHistories success")
		return response, nil
	}
}
