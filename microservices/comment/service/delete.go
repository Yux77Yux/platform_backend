package service

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	internal "github.com/Yux77Yux/platform_backend/microservices/comment/internal"
)

func (s *Server) DeleteComment(ctx context.Context, req *generated.DeleteCommentRequest) (*emptypb.Empty, error) {
	log.Println("info: delete Comment service start")

	err := internal.DeleteComment(req)
	if err != nil {
		log.Println("error: delete Comment occur fail: ", err)
		return nil, err
	}

	log.Println("info: delete Comment occur success")
	return &emptypb.Empty{}, nil
}
