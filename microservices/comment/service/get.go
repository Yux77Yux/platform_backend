package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	internal "github.com/Yux77Yux/platform_backend/microservices/comment/internal"
)

func (s *Server) GetComments(ctx context.Context, req *generated.GetCommentsRequest) (*generated.GetCommentsResponse, error) {
	log.Println("info: GetComments service start")

	response, err := internal.GetComments(ctx, req)
	if err != nil {
		log.Println("error: GetComments occur fail: ", err)
		return response, nil
	}

	log.Println("info: GetComments occur success")
	return response, nil
}

func (s *Server) InitialComments(ctx context.Context, req *generated.InitialCommentsRequest) (*generated.InitialCommentsResponse, error) {
	log.Println("info: InitalComments service start")

	response, err := internal.InitialComments(ctx, req)
	if err != nil {
		log.Println("error: InitalComments occur fail: ", err)
		return response, nil
	}

	log.Println("info: InitalComments occur success")
	return response, nil
}

func (s *Server) InitialSecondComments(ctx context.Context, req *generated.InitialSecondCommentsRequest) (*generated.InitialSecondCommentsResponse, error) {
	log.Println("info: InitialSecondComments service start")

	response, err := internal.InitialSecondComments(ctx, req)
	if err != nil {
		log.Println("error: InitialSecondComments occur fail: ", err)
		return response, nil
	}

	log.Println("info: InitialSecondComments occur success")
	return response, nil
}

func (s *Server) GetTopComments(ctx context.Context, req *generated.GetTopCommentsRequest) (*generated.GetCommentsResponse, error) {
	log.Println("info: GetTopComment service start")

	response, err := internal.GetTopComments(ctx, req)
	if err != nil {
		log.Println("error: get TopComment occur fail: ", err)
		return response, nil
	}

	log.Println("info: get TopComment occur success")
	return response, nil
}

func (s *Server) GetSecondComments(ctx context.Context, req *generated.GetSecondCommentsRequest) (*generated.GetCommentsResponse, error) {
	log.Println("info: GetSecondComment service start")

	response, err := internal.GetSecondComments(ctx, req)
	if err != nil {
		log.Println("error: GetSecondComment occur fail: ", err)
		return response, nil
	}

	log.Println("info: GetSecondComment occur success")
	return response, nil
}

func (s *Server) GetReplyComments(ctx context.Context, req *generated.GetReplyCommentsRequest) (*generated.GetCommentsResponse, error) {
	log.Println("info: GetReplyComment service start")

	response, err := internal.GetReplyComments(ctx, req)
	if err != nil {
		log.Println("error: GetReplyComment occur fail: ", err)
		return response, nil
	}

	log.Println("info: GetReplyComment occur success")
	return response, nil
}
