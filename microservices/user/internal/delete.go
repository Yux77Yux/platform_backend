package internal

import (
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	db "github.com/Yux77Yux/platform_backend/microservices/user/repository"
)

func CancelFollow(req *generated.FollowRequest) (*emptypb.Empty, error) {
	follow := req.GetFollow()
	err := db.CancelFollow(follow)
	if err != nil {
		log.Printf("error:CancelFollow %v", err)
		return nil, err
	}
	return nil, nil
}
