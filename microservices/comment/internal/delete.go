package internal

import (
	"context"
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
	messaging "github.com/Yux77Yux/platform_backend/microservices/comment/messaging"
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
	err = messaging.SendMessage(ctx, messaging.DeleteComment, messaging.DeleteComment, afterAuth)
	if err != nil {
		err = fmt.Errorf("error: SendMessage DeleteComment error %w", err)
		return err
	}

	return nil
}
