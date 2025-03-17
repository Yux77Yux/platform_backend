package receiver

// 由于不同的exchange，需要不同的接收者，事实上需要被调度，统一开关

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/anypb"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	common "github.com/Yux77Yux/platform_backend/generated/common"
)

func JoinCommentProcessor(_ context.Context, msg *anypb.Any) error {
	req := new(generated.Comment)
	err := msg.UnmarshalTo(req)
	if err != nil {
		return err
	}
	// 传递至责任链
	dispatcher.HandleRequest(req, DispatchInsert)
	return nil
}

func DeleteCommentProcessor(ctx context.Context, msg *anypb.Any) error {
	req := new(common.AfterAuth)
	// 反序列化
	err := msg.UnmarshalTo(req)
	if err != nil {
		return fmt.Errorf("deleteCommentProcessor processor error: %w", err)
	}

	creationId, userId, err := db.GetCreationIdInTransaction(ctx, req.GetCommentId())
	if err != nil {
		return err
	}
	if creationId <= 0 {
		return fmt.Errorf("creationId <= 0")
	}
	req.CreationId = creationId

	if req.GetUserId() != -403 && req.GetUserId() != userId {
		return fmt.Errorf("not the master")
	}

	// 发送集中处理
	dispatcher.HandleRequest(req, DispatchDelete)
	return nil
}
