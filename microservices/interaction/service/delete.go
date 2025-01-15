package service

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	generated "github.com/Yux77Yux/platform_backend/generated/internal"
	internal "github.com/Yux77Yux/platform_backend/microservices/internal/internal"
)

func (s *Server) Deleteinternal(ctx context.Context, req *generated.DeleteinternalRequest) (*emptypb.Empty, error) {
	log.Println("info: delete internal service start")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("error: service exceeded timeout: %v", err)
		return nil, err
	default:
		err := internal.Deleteinternal(req)
		if err != nil {
			log.Println("error: delete internal occur fail: ", err)
			return nil, err
		}

		log.Println("info: delete internal occur success")
		return &emptypb.Empty{}, nil
	}
}
