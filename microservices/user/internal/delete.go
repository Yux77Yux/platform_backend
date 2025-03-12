package internal

import (
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/user"
	db "github.com/Yux77Yux/platform_backend/microservices/user/repository"
)

func CancelFollow(req *generated.FollowRequest) error {
	follow := req.GetFollow()
	err := db.CancelFollow(follow)
	if err != nil {
		log.Printf("error:CancelFollow %v", err)
		return err
	}
	return nil
}
