package cache

import (
	"context"
	"fmt"
	"strconv"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	tools "github.com/Yux77Yux/platform_backend/microservices/creation/tools"
)

// POST
func CreationAddInCache(creationInfo *generated.CreationInfo) error {
	ctx := context.Background()

	creation := creationInfo.GetCreation()

	id := strconv.FormatInt(creation.GetCreationId(), 10)

	resultCh := make(chan error, 1)

	categoryId := creation.GetBaseInfo().GetCategoryId()

	reqFunc := func(CacheClient CacheInterface) {
		err := CacheClient.SetFieldsHash(ctx, "CreationInfo", id,
			"author_id", creation.GetBaseInfo().GetAuthorId(),
			"src", creation.GetBaseInfo().GetSrc(),
			"thumbnail", creation.GetBaseInfo().GetThumbnail(),
			"title", creation.GetBaseInfo().GetTitle(),
			"bio", creation.GetBaseInfo().GetBio(),
			"status", creation.GetBaseInfo().GetStatus().String(),
			"duration", creation.GetBaseInfo().GetDuration(),
			"category_id", categoryId,
			"upload_time", creation.GetUploadTime().AsTime(),

			"views", 0,
			"saves", 0,
			"likes", 0,
			"publish_time", "none",

			"comment_count", 0,

			"category_name", tools.Categories[categoryId].Name,
			"category_parent", tools.Categories[categoryId].Parent,
		)
		resultCh <- err
	}

	cacheRequestChannel <- reqFunc

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

// GET
func GetCreationInfo(creation_id int64, fields []string) (map[string]string, error) {
	ctx := context.Background()

	resultCh := make(chan struct {
		creationInfo map[string]string
		err          error
	}, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		if len(fields) == 0 {
			result, err := CacheClient.GetAllHash(ctx, "CreationInfo", strconv.FormatInt(creation_id, 10))
			resultCh <- struct {
				creationInfo map[string]string
				err          error
			}{
				creationInfo: result,
				err:          err,
			}
		} else {
			values, err := CacheClient.GetAnyHash(ctx, "CreationInfo", strconv.FormatInt(creation_id, 10), fields...)
			// 构造结果 map
			result := make(map[string]string, len(fields))
			for i, field := range fields {
				// 类型断言并检查 nil 值
				if values[i] != nil {
					strValue, ok := values[i].(string)
					if !ok {
						err = fmt.Errorf("unexpected value type for field %s", field)
						break
					}
					result[field] = strValue
				}
			}
			resultCh <- struct {
				creationInfo map[string]string
				err          error
			}{
				creationInfo: result,
				err:          err,
			}
		}
	}

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result.err != nil {
			return nil, result.err
		}

		return result.creationInfo, nil
	}
}

// DEL
func DeleteCreation(creation_id int64) error {
	idStr := strconv.FormatInt(creation_id, 10)
	ctx := context.Background()

	resultCh := make(chan struct {
		err error
	}, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		err := CacheClient.DelKey(ctx, "Hash_CreationInfo", idStr)
		resultCh <- struct {
			err error
		}{
			err: err,
		}
	}

	// 使用 select 来监听超时和结果
	select {
	case <-ctx.Done():
		return fmt.Errorf("timeout: %w", ctx.Err())
	case result := <-resultCh:
		if result.err != nil {
			return result.err
		}

		return nil
	}
}

// UPDATE
func UpdateCreation(creation *generated.CreationUpdated) error {
	var (
		creationId = creation.GetCreationId()
		thumbnail  = creation.GetThumbnail()
		title      = creation.GetTitle()
		bio        = creation.GetBio()
		src        = creation.GetSrc()
		duration   = creation.GetDuration()
	)

	values := make([]any, 0, 5*2)
	if thumbnail != "" {
		values = append(values, "thumbnail", thumbnail)
	}
	if title != "" {
		values = append(values, "title", title)
	}
	if bio != "" {
		values = append(values, "bio", bio)
	}
	if src != "" {
		values = append(values, "src", src)
	}
	if duration != 0 {
		values = append(values, "duration", duration)
	}
	if len(values) <= 0 {
		return nil
	}

	ctx := context.Background()
	err := CacheClient.SetFieldsHash(ctx, "CreationInfo", strconv.FormatInt(creationId, 10), values...)
	return err
}

func UpdateCreationStatus(creation *generated.CreationUpdateStatus) error {
	var (
		creationId = creation.GetCreationId()
		status     = creation.GetStatus()
	)

	values := make([]any, 0, 2)
	values = append(values, "status", status.String())

	ctx := context.Background()
	err := CacheClient.SetFieldsHash(ctx, "CreationInfo", strconv.FormatInt(creationId, 10), values...)
	return err
}
