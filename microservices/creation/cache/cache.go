package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
)

func CreationAddInCache(creation *generated.Creation) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	id := strconv.FormatInt(creation.GetCreationId(), 10)

	resultCh := make(chan error, 1)

	reqFunc := func(CacheClient CacheInterface) {
		err := CacheClient.SetFieldsHash(ctx, "Creation", id,
			"author_id", creation.GetBaseInfo().GetAuthorId(),
			"arc", creation.GetBaseInfo().GetSrc(),
			"thumbnail", creation.GetBaseInfo().GetThumbnail(),
			"title", creation.GetBaseInfo().GetTitle(),
			"bio", creation.GetBaseInfo().GetBio(),
			"status", creation.GetBaseInfo().GetStatus().String(),
			"duration", creation.GetBaseInfo().GetDuration(),
			"category_id", creation.GetBaseInfo().GetCategoryId(),
			"upload_time", creation.GetUploadTime().AsTime(),
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
