package service

import (
	"context"
	"log"

	"github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	internal "github.com/Yux77Yux/platform_backend/microservices/creation/internal"
)

func (s *Server) GetCreation(ctx context.Context, req *generated.GetCreationRequest) (*generated.GetCreationResponse, error) {
	log.Println("info: get creation service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetCreationResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetCreation(req)
		if err != nil {
			log.Println("error: get creation occur fail: ", err)
			return response, nil
		}

		log.Println("info: get creation occur success")
		return response, nil
	}
}

func (s *Server) GetSimilarCreationList(ctx context.Context, req *generated.GetPublicCreationListRequest) (*generated.GetCreationListResponse, error) {
	log.Println("info: GetSimilarCreationList service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetCreationListResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetSimilarCreationList(req)
		if err != nil {
			log.Println("error: get creation occur fail: ", err)
			return response, nil
		}

		log.Println("info: get creation occur success")
		return response, nil
	}
}

func (s *Server) GetSpaceCreationList(ctx context.Context, req *generated.GetPublicCreationListRequest) (*generated.GetCreationListResponse, error) {
	log.Println("info: GetSpaceCreationList service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetCreationListResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetSpaceCreationList(req)
		if err != nil {
			log.Println("error: get creation occur fail: ", err)
			return response, nil
		}

		log.Println("info: get creation occur success")
		return response, nil
	}
}

func (s *Server) GetCollectionCreationList(ctx context.Context, req *generated.GetSpecificCreationListRequest) (*generated.GetCreationListResponse, error) {
	log.Println("info: GetCollectionCreationList service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetCreationListResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetCollectionCreationList(req)
		if err != nil {
			log.Println("error: get creation occur fail: ", err)
			return response, nil
		}

		log.Println("info: get creation occur success")
		return response, nil
	}
}

func (s *Server) GetHomeCreationList(ctx context.Context, req *generated.GetSpecificCreationListRequest) (*generated.GetCreationListResponse, error) {
	log.Println("info: GetHomeCreationList service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetCreationListResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetHomeCreationList(req)
		if err != nil {
			log.Println("error: get creation occur fail: ", err)
			return response, nil
		}

		log.Println("info: get creation occur success")
		return response, nil
	}
}
