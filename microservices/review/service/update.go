package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
	internal "github.com/Yux77Yux/platform_backend/microservices/review/internal"
)

func (s *Server) UpdateReview(ctx context.Context, req *generated.UpdateReviewRequest) (*generated.UpdateReviewResponse, error) {
	log.Printf("info: UpdateReview service start")
	response, err := internal.UpdateReview(req)
	if err != nil {
		log.Printf("info: UpdateReview service fail")
		return response, nil
	}
	log.Printf("info: UpdateReview service success")
	return response, nil
}
