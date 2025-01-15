package internal

import (
	// "fmt"

	"fmt"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	cache "github.com/Yux77Yux/platform_backend/microservices/creation/cache"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
	// messaging "github.com/Yux77Yux/platform_backend/microservices/creation/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/creation/repository"
)

func DeleteCreation(req *generated.DeleteCreationRequest) error {
	accessToken := req.GetAccessToken().GetValue()
	if accessToken == "" {
		// 无token直接返回
		return fmt.Errorf("no token")
	}

	// 取作品信息，鉴权
	pass, user_id, err := auth.Auth("delete", "creation", accessToken)
	if err != nil {
		return err
	}
	if !pass {
		return fmt.Errorf("no pass")
	}
	// 以上为鉴权

	creationId := req.GetCreationId()

	// 取作品作者id
	var author_id int64 = -1
	str, err := cache.GetCreationInfo(creationId, []string{"author_id"})
	if err != nil {
		return fmt.Errorf("error: get author in cache error")
	}

	if str["author_id"] == "" {
		author_id, err = db.GetAuthorIdInTransaction(creationId)
		if err != nil {
			return err
		}
	}
	if author_id != user_id {
		return fmt.Errorf("error: author %v not match the token", author_id)
	}

	// 开始删除

	// 删除数据库中作品
	err = db.DeleteCreationInTransaction(creationId)
	if err != nil {
		return fmt.Errorf("error: db error %w", err)
	}

	// 删除缓存中作品
	err = cache.DeleteCreation(creationId)
	if err != nil {
		return fmt.Errorf("error: cache error %w", err)
	}

	return nil
}
