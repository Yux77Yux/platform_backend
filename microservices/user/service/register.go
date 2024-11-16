package service

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	generatedUser "github.com/Yux77Yux/platform_backend/generated/user"
	userCache "github.com/Yux77Yux/platform_backend/microservices/user/cache"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
	userMQ "github.com/Yux77Yux/platform_backend/microservices/user/messaging"
)

func (s *Server) Register(ctx context.Context, req *generatedUser.RegisterRequest) (*generatedUser.RegisterResponse, error) {
	log.Println("info: register service start")
	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout%v", err)
		return &generatedUser.RegisterResponse{
			Success: false,
		}, err
	default:
		user_credentials := req.GetUserCredentials()

		//检查空值
		if user_credentials.GetUsername() == "" || user_credentials.GetPassword() == "" {
			err := status.Errorf(codes.InvalidArgument, "username and password cannot be empty")
			log.Printf("warning: %v", err)
			return &generatedUser.RegisterResponse{
				Success: false,
			}, err
		}

		//redis查询账号是否唯一
		pos, err := userCache.ExistsUsername(user_credentials.GetUsername())
		if err != nil {
			log.Printf("error: failed to use redis client: %v", err)
		}
		if pos > -1 {
			return &generatedUser.RegisterResponse{
				Success: false,
			}, fmt.Errorf("username already exists")
		}

		select {
		//闭包传递
		case internal.UserRequestChan <- internal.RequestHandlerFunc(func(req_id string) error {
			log.Printf("info: handling request with ID: %s\n", req_id)
			err := userMQ.SendMessage("register_exchange", "register", user_credentials)
			if err != nil {
				return fmt.Errorf("err: request_id: %s ,message: %w", req_id, err)
			}
			return nil
		}):
			log.Println("info: handler completely sent to UserRequestChan")
		default:
			log.Println("warning: userChan is full or closed, registration may not complete")
			return &generatedUser.RegisterResponse{
				Success: false,
			}, status.Errorf(codes.Unavailable, "user request queue is full, please try again later")
		}

		return &generatedUser.RegisterResponse{
			Success: true,
		}, nil
	}
}
