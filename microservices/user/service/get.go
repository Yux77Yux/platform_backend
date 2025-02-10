package service

import (
	"context"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

func (s *Server) Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {
	log.Println("info: login service start")

	response, err := internal.Login(ctx, req)
	if err != nil {
		log.Printf("error: login occur fail %v", err)
		return response, nil
	}

	log.Println("info: login occur success")
	return response, nil
}

func (s *Server) GetUser(ctx context.Context, req *generated.GetUserRequest) (*generated.GetUserResponse, error) {
	log.Println("info: get user service start")

	response, err := internal.GetUser(ctx, req)
	if err != nil {
		log.Println("error: get user occur fail: ", err)
		return response, nil
	}

	log.Println("info: get user occur success")
	return response, nil
}

func (s *Server) GetFolloweesByTime(ctx context.Context, req *generated.GetFollowRequest) (*generated.GetFollowResponse, error) {
	log.Println("info: GetFolloweesByTime service start")

	response, err := internal.GetFolloweesByTime(ctx, req)
	if err != nil {
		log.Println("error: get user occur fail: ", err)
		return response, nil
	}

	log.Println("info: get user occur success")
	return response, nil
}

func (s *Server) GetFolloweesByViews(ctx context.Context, req *generated.GetFollowRequest) (*generated.GetFollowResponse, error) {
	log.Println("info: GetFolloweesByViews service start")

	response, err := internal.GetFolloweesByViews(ctx, req)
	if err != nil {
		log.Println("error: get user occur fail: ", err)
		return response, nil
	}

	log.Println("info: get user occur success")
	return response, nil
}

func (s *Server) GetFollowers(ctx context.Context, req *generated.GetFollowRequest) (*generated.GetFollowResponse, error) {
	log.Println("info: GetFollowers service start")

	response, err := internal.GetFollowers(ctx, req)
	if err != nil {
		log.Println("error: get user occur fail: ", err)
		return response, nil
	}

	log.Println("info: get user occur success")
	return response, nil
}
