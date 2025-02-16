package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/aggregator"
	internal "github.com/Yux77Yux/platform_backend/microservices/aggregator/internal"
)

func (s *Server) Collections(ctx context.Context, req *generated.CollectionsRequest) (*generated.GetCardsResponse, error) {
	log.Println("info: Collections service start")

	response, err := internal.Collections(ctx, req)
	if err != nil {
		log.Printf("error: Collections occur fail: %v", err)
		return response, nil
	}

	log.Println("info: Collections occur success")
	return response, nil
}

func (s *Server) History(ctx context.Context, req *generated.HistoryRequest) (*generated.GetCardsResponse, error) {
	log.Println("info: History service start")

	response, err := internal.History(ctx, req)
	if err != nil {
		log.Printf("error: History occur fail: %v", err)
		return response, nil
	}

	log.Println("info: History occur success")
	return response, nil
}

func (s *Server) HomePage(ctx context.Context, req *generated.HomeRequest) (*generated.GetCardsResponse, error) {
	log.Println("info: HomePage service start")

	response, err := internal.HomePage(ctx, req)
	if err != nil {
		log.Printf("error: HomePage occur fail: %v", err)
		return response, nil
	}

	log.Println("info: HomePage occur success")
	return response, nil
}
