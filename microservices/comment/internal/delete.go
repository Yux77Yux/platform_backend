package internal

import (
	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
	// messaging "github.com/Yux77Yux/platform_backend/microservices/comment/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
)

func DeleteComment(req *generated.DeleteCommentRequest) error {
	accessToken := req.GetAccessToken().GetValue()
	// 取作品信息，鉴权
	pass, user_id, err := auth.Auth("delete", "creation", accessToken)
	if err != nil {
		return err
	}
	if !pass {
		return fmt.Errorf("no pass")
	}
	// 以上为鉴权

	comment_id := req.GetCommentId()

	// 取发布者id
	var author_id int64 = -1
	author_id, err = db.GetPublisherIdInTransaction(comment_id)
	if err != nil {
		return err
	}
	if author_id != user_id {
		return fmt.Errorf("error: author %v not match the token", author_id)
	}

	// 开始删除

	// 删除数据库中作品
	err = db.DeleteCommentInTransaction(creationId)
	if err != nil {
		return fmt.Errorf("error: db error %w", err)
	}

	return nil
}
