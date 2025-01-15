package service

import (
	"context"
	"log"

	"github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	internal "github.com/Yux77Yux/platform_backend/microservices/interaction/internal"
)

func (s *Server) Getinteraction(ctx context.Context, req *generated.GetinteractionRequest) (*generated.GetinteractionResponse, error) {
	log.Println("info: get interaction service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetinteractionResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.Getinteraction(req)
		if err != nil {
			log.Println("error: get interaction occur fail: ", err)
			return response, nil
		}

		log.Println("info: get interaction occur success")
		return response, nil
	}
}

func (s *Server) GetSimilarinteractionList(ctx context.Context, req *generated.GetSpecificinteractionListRequest) (*generated.GetinteractionListResponse, error) {
	log.Println("info: GetSimilarinteractionList service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetinteractionListResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetSimilarinteractionList(req)
		if err != nil {
			log.Println("error: get interaction occur fail: ", err)
			return response, nil
		}

		log.Println("info: get interaction occur success")
		return response, nil
	}
}

func (s *Server) GetCollectioninteractionList(ctx context.Context, req *generated.GetSpecificinteractionListRequest) (*generated.GetinteractionListResponse, error) {
	log.Println("info: GetCollectioninteractionList service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetinteractionListResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetCollectioninteractionList(req)
		if err != nil {
			log.Println("error: get interaction occur fail: ", err)
			return response, nil
		}

		log.Println("info: get interaction occur success")
		return response, nil
	}
}

func (s *Server) GetSpaceinteractionList(ctx context.Context, req *generated.GetSpecificinteractionListRequest) (*generated.GetinteractionListResponse, error) {
	log.Println("info: GetSpaceinteractionList service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetinteractionListResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetCollectioninteractionList(req)
		if err != nil {
			log.Println("error: get interaction occur fail: ", err)
			return response, nil
		}

		log.Println("info: get interaction occur success")
		return response, nil
	}
}

func (s *Server) GetHomeinteractionList(ctx context.Context, req *generated.GetSpecificinteractionListRequest) (*generated.GetinteractionListResponse, error) {
	log.Println("info: GetHomeinteractionList service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetinteractionListResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetHomeinteractionList(req)
		if err != nil {
			log.Println("error: get interaction occur fail: ", err)
			return response, nil
		}

		log.Println("info: get interaction occur success")
		return response, nil
	}
}
