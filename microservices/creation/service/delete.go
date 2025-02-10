package service

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	internal "github.com/Yux77Yux/platform_backend/microservices/creation/internal"
)

func (s *Server) DeleteCreation(ctx context.Context, req *generated.DeleteCreationRequest) (*emptypb.Empty, error) {
	log.Println("info: delete creation service start")

	err := internal.DeleteCreation(req)
	if err != nil {
		log.Println("error: delete creation occur fail: ", err)
		return nil, err
	}

	log.Println("info: delete creation occur success")
	return &emptypb.Empty{}, nil
}
