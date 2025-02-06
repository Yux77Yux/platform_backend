package service

import (
	"context"
	"log"

	"github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/interaction"
	internal "github.com/Yux77Yux/platform_backend/microservices/interaction/internal"
)

func (s *Server) CancelCollections(ctx context.Context, req *generated.UpdateInteractionsRequest) (*generated.UpdateInteractionResponse, error) {
	log.Println("info: CancelCollections service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.UpdateInteractionResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.CancelCollections(req)
		if err != nil {
			log.Println("error: CancelCollections fail")
			return response, err
		}

		log.Println("info: CancelCollections success")
		return response, nil
	}
}

func (s *Server) CancelLike(ctx context.Context, req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
	log.Println("info: CancelLike service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.UpdateInteractionResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.CancelLike(req)
		if err != nil {
			log.Println("error: CancelLike fail")
			return response, err
		}

		log.Println("info: CancelLike success")
		return response, nil
	}
}

func (s *Server) ClickCollection(ctx context.Context, req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
	log.Println("info: ClickCollection service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.UpdateInteractionResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.ClickCollection(req)
		if err != nil {
			log.Println("error: ClickCollection fail")
			return response, err
		}

		log.Println("info: ClickCollection success")
		return response, nil
	}
}

func (s *Server) ClickLike(ctx context.Context, req *generated.UpdateInteractionRequest) (*generated.UpdateInteractionResponse, error) {
	log.Println("info: ClickLike service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.UpdateInteractionResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.ClickLike(req)
		if err != nil {
			log.Println("error: ClickLike fail")
			return response, err
		}

		log.Println("info: ClickLike success")
		return response, nil
	}
}

func (s *Server) DelHistories(ctx context.Context, req *generated.UpdateInteractionsRequest) (*generated.UpdateInteractionResponse, error) {
	log.Println("info: DelHistories service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.UpdateInteractionResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.DelHistories(req)
		if err != nil {
			log.Println("error: DelHistories fail")
			return response, err
		}

		log.Println("info: DelHistories success")
		return response, nil
	}
}
