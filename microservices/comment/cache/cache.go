package cache

import (
	"context"
	"strconv"
)

func UpdateCommentsCount(ctx context.Context, creationId int64, count int64) error {
	idStr := strconv.FormatInt(creationId, 10)
	return CacheClient.IncrHash(ctx, "CreationInfo", idStr, "comment_count", count)
}
