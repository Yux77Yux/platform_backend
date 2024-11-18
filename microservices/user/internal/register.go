package internal

import (
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	generatedUser "github.com/Yux77Yux/platform_backend/generated/user"
	userCache "github.com/Yux77Yux/platform_backend/microservices/user/cache"
	userMQ "github.com/Yux77Yux/platform_backend/microservices/user/messaging"
)

func Register(user_credentials *generatedUser.UserCredentials) (bool, error) {
	//检查空值
	if user_credentials.GetUsername() == "" || user_credentials.GetPassword() == "" {
		err := status.Errorf(codes.InvalidArgument, "username and password cannot be empty")
		log.Printf("warning: %v", err)
		return false, err
	}

	//redis查询账号是否唯一
	exist, err := userCache.ExistsUsername(user_credentials.GetUsername())
	if err != nil {
		log.Printf("error: failed to use redis client: %v", err)
	}
	if exist {
		log.Printf("info: username already exists")
		return false, fmt.Errorf("info: username already exists")
	}

	password, err := encryptWithTimestamp(user_credentials.GetPassword())
	if err != nil {
		return false, err
	}
	credentials := &generatedUser.UserCredentials{
		Username: user_credentials.GetUsername(),
		Password: password,
	}
	select {
	//闭包传递
	case userRequestChannel <- func(req_id string) error {
		log.Printf("info: handling request with ID: %s\n", req_id)

		err = userMQ.SendMessage("register_exchange", "register", credentials)
		if err != nil {
			return fmt.Errorf("err: request_id: %s ,message: %w", req_id, err)
		}
		return nil
	}:
		log.Println("info: handler completely sent to UserRequestChan")
	default:
		log.Println("warning: userChan is full or closed, registration may not complete")
		return false, status.Errorf(codes.Unavailable, "user request queue is full, please try again later")
	}

	return true, nil
}
