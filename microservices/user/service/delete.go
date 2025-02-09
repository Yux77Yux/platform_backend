package service

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	internal "github.com/Yux77Yux/platform_backend/microservices/user/internal"
)

func (s *Server) CancelFollow(ctx context.Context, req *generated.FollowRequest) (*emptypb.Empty, error) {
	log.Println("info: CancelFollow service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return nil, nil
	default:
		response, err := internal.CancelFollow(req)
		if err != nil {
			log.Printf("error: CancelFollow occur fail %v", err)
			return response, nil
		}

		log.Println("info: CancelFollow occur success")
		return response, nil
	}
}
