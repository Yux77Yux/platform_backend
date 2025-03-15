package internal

import (
	"context"
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	messaging "github.com/Yux77Yux/platform_backend/microservices/comment/messaging"
	tools "github.com/Yux77Yux/platform_backend/microservices/comment/tools"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func DeleteComment(ctx context.Context, req *generated.DeleteCommentRequest) error {
	accessToken := req.GetAccessToken().GetValue()
	pass, user_id, err := auth.Auth("delete", "comment", accessToken)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	if !pass {
		return fmt.Errorf("error: no pass")
	}

	// 第一次过滤，发到消息队列
	afterAuth := &common.AfterAuth{
		UserId:     user_id,
		CommentId:  req.GetCommentId(),
		CreationId: req.GetCreationId(),
	}
	go func(afterAuth *common.AfterAuth, ctx context.Context) {
		traceId, fullName := tools.GetMetadataValue(ctx, "trace-id"), tools.GetMetadataValue(ctx, "full-name")
		err = messaging.SendMessage(ctx, messaging.DeleteComment, messaging.DeleteComment, afterAuth)
		if err != nil {
			tools.LogError(traceId, fullName, err)
		}
	}(afterAuth, ctx)

	return nil
}
