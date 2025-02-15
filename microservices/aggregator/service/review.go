package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	internal "github.com/Yux77Yux/platform_backend/microservices/aggregator/internal"
)

func (s *Server) GetUserReviews(ctx context.Context, req *generated.GetReviewsRequest) (*generated.GetUserReviewsResponse, error) {
	log.Println("info: GetUserReviews service start")

	response, err := internal.GetUserReviews(ctx, req)
	if err != nil {
		log.Printf("error: GetUserReviews occur fail: %v", err)
		return response, nil
	}

	log.Println("info: GetUserReviews occur success")
	return response, nil
}

func (s *Server) GetCreationReviews(ctx context.Context, req *generated.GetReviewsRequest) (*generated.GetCreationReviewsResponse, error) {
	log.Println("info: GetCreationReviews service start")

	response, err := internal.GetCreationReviews(ctx, req)
	if err != nil {
		log.Printf("error: GetCreationReviews occur fail: %v", err)
		return response, nil
	}

	log.Println("info: GetCreationReviews occur success")
	return response, nil
}

func (s *Server) GetCommentReviews(ctx context.Context, req *generated.GetReviewsRequest) (*generated.GetCommentReviewsResponse, error) {
	log.Println("info: GetCommentReviews service start")

	response, err := internal.GetCommentReviews(ctx, req)
	if err != nil {
		log.Printf("error: GetCommentReviews occur fail: %v", err)
		return response, nil
	}

	log.Println("info: GetCommentReviews occur success")
	return response, nil
}

func (s *Server) GetNewUserReviews(ctx context.Context, req *generated.GetNewReviewsRequest) (*generated.GetUserReviewsResponse, error) {
	log.Println("info: GetNewUserReviews service start")

	response, err := internal.GetNewUserReviews(ctx, req)
	if err != nil {
		log.Printf("error: GetNewUserReviews occur fail: %v", err)
		return response, nil
	}

	log.Println("info: GetNewUserReviews occur success")
	return response, nil
}

func (s *Server) GetNewCreationReviews(ctx context.Context, req *generated.GetNewReviewsRequest) (*generated.GetCreationReviewsResponse, error) {
	log.Println("info: GetNewCreationReviews service start")

	response, err := internal.GetNewCreationReviews(ctx, req)
	if err != nil {
		log.Printf("error: GetNewCreationReviews occur fail: %v", err)
		return response, nil
	}

	log.Println("info: GetNewCreationReviews occur success")
	return response, nil
}

func (s *Server) GetNewCommentReviews(ctx context.Context, req *generated.GetNewReviewsRequest) (*generated.GetCommentReviewsResponse, error) {
	log.Println("info: GetNewCommentReviews service start")

	response, err := internal.GetNewCommentReviews(ctx, req)
	if err != nil {
		log.Printf("error: GetNewCommentReviews occur fail: %v", err)
		return response, nil
	}

	log.Println("info: GetNewCommentReviews occur success")
	return response, nil
}
