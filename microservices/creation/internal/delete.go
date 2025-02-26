package internal

import (
	"fmt"
	"log"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	messaging "github.com/Yux77Yux/platform_backend/microservices/creation/messaging"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func DeleteCreation(req *generated.DeleteCreationRequest) error {
	accessToken := req.GetAccessToken().GetValue()
	if accessToken == "" {
		// 无token直接返回
		return fmt.Errorf("no token")
	}

	pass, user_id, err := auth.Auth("delete", "creation", req.GetAccessToken().GetValue())
	if err != nil {
		return fmt.Errorf("405")
	}
	if !pass {
		return fmt.Errorf("403")
	}
	// 以上为鉴权

	creationId := req.GetCreationId()

	deleteInfo := &generated.CreationUpdateStatus{
		CreationId: creationId,
		Status:     generated.CreationStatus_DELETE,
		AuthorId:   user_id,
	}

	err = messaging.SendMessage(messaging.DeleteCreation, messaging.DeleteCreation, deleteInfo)
	if err != nil {
		log.Printf("error: publish failed because %v", err)
		return err
	}

	return nil
}
