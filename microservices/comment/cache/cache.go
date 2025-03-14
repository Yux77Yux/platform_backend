package cache

import (
	"context"
	"strconv"
)

func UpdateCommentsCount(creationId int64, count int64) error {
	ctx := context.Background()
	idStr := strconv.FormatInt(creationId, 10)
	return CacheClient.IncrHash(ctx, "CreationInfo", idStr, "comment_count", count)
}
