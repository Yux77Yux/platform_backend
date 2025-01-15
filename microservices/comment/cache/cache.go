package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
)

func DelTemporaryComments(creationId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	idStr := strconv.FormatInt(creationId, 10)
	return CacheClient.DelKey(ctx, "Hash_TemporaryComments", idStr)
}

func ExistTemporaryComments(creationId int64) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	idStr := strconv.FormatInt(creationId, 10)

	result, err := CacheClient.ExistsHash(ctx, "TemporaryComments", idStr)
	if err != nil {
		return false, err
	}

	return result, nil
}

func RefreshTemporaryComment(creationId int64) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	idStr := strconv.FormatInt(creationId, 10)

	return CacheClient.DelHash(ctx, "TemporaryComments", idStr, "queryComment", "queryCommentContent", "count")
}

func SetTemporaryComments(comment *generated.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	creationTdStr := strconv.FormatInt(comment.GetCreationId(), 10)
	userIdStr := strconv.FormatInt(comment.GetUserId(), 10)

	newComment := fmt.Sprintf("(%d, %d, %d, %s, %s)",
		comment.GetRoot(), comment.GetParent(), comment.GetDialog(), creationTdStr, userIdStr)

	newCommentContent := fmt.Sprintf("(?, %s,%s)",
		comment.GetContent(), comment.GetMedia())

	var (
		queryComment        string
		queryCommentContent string
		count               int = 0
	)

	values, err := GetTemporaryComments(comment.GetCreationId())
	if err != nil {
		return fmt.Errorf("error in GetTemporaryComments :%w", err)
	}
	if values == nil || len(values) <= 0 {
		queryComment = `
				INSERT INTO db_comment_1.comment (
					root,
					parent,
					dialog,
					creation_id,
					user_id)
				VALUES` + newComment

		queryCommentContent = `
				INSERT INTO db_comment_1.comment (
					comment_id,
					content,
					media)
				VALUES` + newCommentContent
	} else {
		queryComment = values["queryComment"] + "," + newComment
		queryCommentContent = values["queryCommentContent"] + "," + newCommentContent
		count, err = strconv.Atoi(values["count"])
		if err != nil {
			return fmt.Errorf("count in redis is not int %w", err)
		}
	}

	resultCh := make(chan error, 1)
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		err := CacheClient.SetFieldsHash(ctx, "TemporaryComments", creationTdStr,
			"queryComment", queryComment,
			"queryCommentContent", queryCommentContent,
			"count", strconv.Itoa(count+1),
		)
		resultCh <- err
	}

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result != nil {
			return result
		}
		return nil
	}
}

func GetTemporaryComments(creationId int64) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	idStr := strconv.FormatInt(creationId, 10)

	result, err := CacheClient.GetAllHash(ctx, "TemporaryComments", idStr)
	if err != nil {
		return nil, err
	}
	return result, nil
}
