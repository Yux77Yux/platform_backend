package service

import (
	"context"
	"log"

	"google.golang.org/protobuf/proto"

	generatedUser "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

func (s *Server) Register(ctx context.Context, req *generatedUser.RegisterRequest) (*generatedUser.RegisterResponse, error) {
	log.Println("info: register service start")
	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return &generatedUser.RegisterResponse{
			Success: false,
			Error:   proto.String(err.Error()),
		}, err
	default:
		success, err := internal.Register(req.GetUserCredentials())
		return &generatedUser.RegisterResponse{
			Success: success,
		}, err
	}
}
