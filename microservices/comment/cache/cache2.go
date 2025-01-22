package cache

import (
	"context"
	// "fmt"
	"strconv"
	// "google.golang.org/protobuf/proto"
	// generated "github.com/Yux77Yux/platform_backend/generated/comment"
)

func UpdateCommentsCount(creationId int64, count int64) error {
	ctx := context.Background()

	idStr := strconv.FormatInt(creationId, 10)
	resultCh := make(chan error, 1)

	cacheRequestChannel <- func(CacheClient CacheInterface) {
		err := CacheClient.IncrHash(ctx, "CreationInfo", idStr, "comment_count", count)
		resultCh <- err
	}
	return <-resultCh
}
