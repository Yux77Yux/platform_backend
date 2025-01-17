package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	internal "github.com/Yux77Yux/platform_backend/microservices/comment/internal"
)

func (s *Server) InitalComments(ctx context.Context, req *generated.InitalCommentsRequest) (*generated.InitalCommentsResponse, error) {
	log.Println("info: InitalComments service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.InitalCommentsResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.InitalComments(req)
		if err != nil {
			log.Println("error: InitalComments occur fail: ", err)
			return response, nil
		}

		log.Println("info: InitalComments occur success")
		return response, nil
	}
}

func (s *Server) GetTopComments(ctx context.Context, req *generated.GetTopCommentsRequest) (*generated.GetCommentsResponse, error) {
	log.Println("info: GetTopComment service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetCommentsResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetTopComments(req)
		if err != nil {
			log.Println("error: get TopComment occur fail: ", err)
			return response, nil
		}

		log.Println("info: get TopComment occur success")
		return response, nil
	}
}

func (s *Server) GetSecondComments(ctx context.Context, req *generated.GetSecondCommentsRequest) (*generated.GetCommentsResponse, error) {
	log.Println("info: GetSecondComment service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetCommentsResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetSecondComments(req)
		if err != nil {
			log.Println("error: GetSecondComment occur fail: ", err)
			return response, nil
		}

		log.Println("info: GetSecondComment occur success")
		return response, nil
	}
}

func (s *Server) GetReplyComments(ctx context.Context, req *generated.GetReplyCommentsRequest) (*generated.GetCommentsResponse, error) {
	log.Println("info: GetReplyComment service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generated.GetCommentsResponse{
			Msg: &common.ApiResponse{
				Status:  common.ApiResponse_FAILED,
				Code:    "408",
				Message: "Time out",
				Details: err.Error(),
			},
		}, nil
	default:
		response, err := internal.GetReplyComments(req)
		if err != nil {
			log.Println("error: GetReplyComment occur fail: ", err)
			return response, nil
		}

		log.Println("info: GetReplyComment occur success")
		return response, nil
	}
}
