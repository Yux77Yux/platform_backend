package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
	internal "github.com/Yux77Yux/platform_backend/microservices/review/internal"
)

func (s *Server) GetReviews(ctx context.Context, req *generated.GetReviewsRequest) (*generated.GetReviewsResponse, error) {
	log.Printf("info: GetReviews start")
	response, err := internal.GetReviews(ctx, req)
	if err != nil {
		log.Println("error: GetReviews occur fail")
		return response, nil
	}

	log.Printf("info: GetReviews over")
	return response, nil
}

func (s *Server) GetNewReviews(ctx context.Context, req *generated.GetNewReviewsRequest) (*generated.GetReviewsResponse, error) {
	log.Printf("info: GetNewReviews start")
	response, err := internal.GetNewReviews(ctx, req)
	if err != nil {
		log.Println("error: GetNewReviews occur fail")
		return response, nil
	}

	log.Printf("info: GetNewReviews over")
	return response, nil
}
