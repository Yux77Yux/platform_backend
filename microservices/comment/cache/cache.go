package cache

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/protobuf/proto"

	generated "github.com/Yux77Yux/platform_backend/generated/comment"
)

func ExistTemporaryComments(creationId int64) (bool, error) {
	ctx := context.Background()

	idStr := strconv.FormatInt(creationId, 10)

	result, err := CacheClient.ExistsHash(ctx, "TemporaryComments", idStr)
	if err != nil {
		return false, err
	}

	return result, nil
}

// 插入评论
func PushTemporaryComment(comment *generated.Comment) error {
	ctx := context.Background()

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
	ctx := context.Background()

	idStr := strconv.FormatInt(creationId, 10)

	result, err := CacheClient.LRangeList(ctx, "TemporaryComments", idStr, 0, 49)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func RefreshTemporaryComments(creationId int64, count int64) error {
	ctx := context.Background()

	idStr := strconv.FormatInt(creationId, 10)
	return CacheClient.LTrimList(ctx, "TemporaryComments", idStr, 0, count-1)
}

func PushChangingTemporaryComments(creationId int64, comments []string) error {
	ctx := context.Background()

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

func RefreshChangingTemporaryComments(creationId int64, count int64) error {
	ctx := context.Background()

	idStr := strconv.FormatInt(creationId, 10)
	return CacheClient.LTrimList(ctx, "ChangingTemporaryComments", idStr, 0, count-1)
}

// 将评论改成待删除状态
func PushDeleteStatusComment(comment *generated.AfterAuth) error {
	ctx := context.Background()

	data, err := proto.Marshal(comment)
	if err != nil {
		return fmt.Errorf("proto Marshal error%w", err)
	}

	idStr := strconv.FormatInt(comment.GetCreationId(), 10)

	resultCh := make(chan error, 1)
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		// List尾部 推入
		err := CacheClient.RPushList(ctx, "DeleteStatusComments", idStr, data)
		if err != nil {
			err = fmt.Errorf("error RPushList error %w", err)
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

func GetDeleteStatusComments(creationId int64) ([]string, error) {
	ctx := context.Background()

	idStr := strconv.FormatInt(creationId, 10)

	result, err := CacheClient.LRangeList(ctx, "DeleteStatusComments", idStr, 0, 49)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func RefreshDeleteStatusComments(creationId int64, count int64) error {
	ctx := context.Background()

	idStr := strconv.FormatInt(creationId, 10)
	return CacheClient.LTrimList(ctx, "DeleteStatusComments", idStr, 0, count-1)
}

// 永久删除评论
// 将永久删除评论待入
func PushPreDeleteComments(comments []string) error {
	ctx := context.Background()

	count := len(comments)
	values := make([]interface{}, 0, count)
	for _, value := range comments {
		values = append(values, value)
	}

	resultCh := make(chan error, 1)
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		// List尾部 推入
		err := CacheClient.RPushList(ctx, "PreDeleteComments", "Clear", values...)
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

// 取永久删除评论
func GetPreDeleteComments() ([]string, error) {
	ctx := context.Background()

	result, err := CacheClient.LRangeList(ctx, "PreDeleteComments", "Clear", 0, 49)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ClearDeleteComments(count int64) error {
	ctx := context.Background()

	return CacheClient.LTrimList(ctx, "PreDeleteComments", "Clear", 0, count-1)
}

// 查询评论
func PushSelectComment(comment []byte) error {
	ctx := context.Background()

	resultCh := make(chan error, 1)
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		// List尾部 推入
		err := CacheClient.RPushList(ctx, "SelectComments", "", comment)
		if err != nil {
			err = fmt.Errorf("error RPushList error %w", err)
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

func GetSelectComments() ([]string, error) {
	ctx := context.Background()

	result, err := CacheClient.LRangeList(ctx, "SelectComments", "", 0, 99)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func RefreshSelectComments(count int64) error {
	ctx := context.Background()

	return CacheClient.LTrimList(ctx, "SelectComments", "", 0, count-1)
}

func PushChangingSelectComments(comments []string) error {
	ctx := context.Background()

	count := len(comments)
	values := make([]interface{}, 0, count)
	for _, value := range comments {
		values = append(values, value)
	}

	resultCh := make(chan error, 1)
	cacheRequestChannel <- func(CacheClient CacheInterface) {
		// List尾部 推入
		err := CacheClient.RPushList(ctx, "ChangingSelectComments", "", values...)
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

func RefreshChangingSelectComments(count int64) error {
	ctx := context.Background()

	return CacheClient.LTrimList(ctx, "SelectComments", "", 0, count-1)
}
