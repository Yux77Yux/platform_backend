package internal

import (
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	cache "github.com/Yux77Yux/platform_backend/microservices/creation/cache"
	messaging "github.com/Yux77Yux/platform_backend/microservices/creation/messaging"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func DeleteCreation(req *generated.DeleteCreationRequest) error {
	accessToken := req.GetAccessToken().GetValue()
	if accessToken == "" {
		// 无token直接返回
		return fmt.Errorf("no token")
	}

	pass, user_id, err := auth.Auth("update", "creation", req.GetAccessToken().GetValue())
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

	// 删除缓存中作品
	err = cache.UpdateCreationStatus(deleteInfo)
	if err != nil {
		return fmt.Errorf("error: cache error %w", err)
	}

	// 将删除信息发到消息队列
	err = messaging.SendMessage(messaging.UpdateCreationStatus, messaging.UpdateCreationStatus, deleteInfo)
	if err != nil {
		return err
	}

	return nil
}
