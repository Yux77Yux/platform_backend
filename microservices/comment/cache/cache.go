package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
)

func DelChangingTemporaryComments(creationId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	idStr := strconv.FormatInt(creationId, 10)
	return CacheClient.DelKey(ctx, "List_ChangingTemporaryComments", idStr)
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

func RefreshTemporaryComment(creationId int64, count int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	idStr := strconv.FormatInt(creationId, 10)
	return CacheClient.LTrimList(ctx, "TemporaryComments", idStr, 0, count-1)
}

func PushTemporaryComments(comment *generated.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	creationIdStr := strconv.FormatInt(comment.GetCreationId(), 10)

	data, err := proto.Marshal(comment)
	if err != nil {
		return fmt.Errorf("proto Marshal error%w", err)
	}

	resultCh := make(chan error, 1)
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		// List尾部 推入
		errRPushList := CacheClient.RPushList(ctx, "TemporaryComments", creationIdStr, data)
		errIncrHash := CacheClient.IncrHash(ctx, "CreationInfo", creationIdStr, "comment_count", 1)
		if errIncrHash != nil || errRPushList != nil {
			err = fmt.Errorf("error RPushList %w and error IncrHash %w", errRPushList, errIncrHash)
		} else {
			err = nil
		}
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

func GetTemporaryComments(creationId int64) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	idStr := strconv.FormatInt(creationId, 10)

	result, err := CacheClient.LRangeList(ctx, "TemporaryComments", idStr, 0, 199)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ChangingTemporaryComments(creationId int64, comments []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	count := len(comments)
	values := make([]interface{}, 0, count)
	for _, value := range comments {
		values = append(values, value)
	}

	creationIdStr := strconv.FormatInt(creationId, 10)

	resultCh := make(chan error, 1)
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		// List尾部 推入
		err := CacheClient.RPushList(ctx, "ChangingTemporaryComments", creationIdStr, values...)
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
