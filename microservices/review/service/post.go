package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/review"
	internal "github.com/Yux77Yux/platform_backend/microservices/review/internal"
)

func (s *Server) NewReview(ctx context.Context, req *generated.NewReviewRequest) (*generated.NewReviewResponse, error) {
	log.Printf("info: NewReview start")
	response, err := internal.NewReview(ctx, req)
	if err != nil {
		log.Println("error: NewReview occur fail")
		return response, nil
	}

	log.Printf("info: NewReview over")
	return response, nil
}
