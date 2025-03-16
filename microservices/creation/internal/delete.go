package internal

import (
	"context"
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	tools "github.com/Yux77Yux/platform_backend/microservices/creation/tools"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func DeleteCreation(ctx context.Context, req *generated.DeleteCreationRequest) error {
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
	go func(deleteInfo *generated.CreationUpdateStatus, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err = messaging.SendMessage(ctx, EXCHANGE_DELETE_CREATION, KEY_DELETE_CREATION, deleteInfo)
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(deleteInfo, ctx)

	return nil
}
