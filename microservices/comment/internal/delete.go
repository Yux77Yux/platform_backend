package internal

import (
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	messaging "github.com/Yux77Yux/platform_backend/microservices/comment/messaging"
)

func DeleteComment(req *generated.DeleteCommentRequest) error {
	// 开始异步删除
	// 异步处理
	err := messaging.SendMessage("DeleteComment", "DeleteComment", req)
	if err != nil {
		err = fmt.Errorf("error: SendMessage DeleteComment error %w", err)
		return err
	}

	return nil
}
