package cache

import (
	"context"
	"strconv"
)

type CacheMethodStruct struct {
	CacheClient CacheInterface
}

func (c *CacheMethodStruct) UpdateCommentsCount(ctx context.Context, creationId int64, count int64) error {
	idStr := strconv.FormatInt(creationId, 10)
	return c.CacheClient.IncrHash(ctx, "CreationInfo", idStr, "comment_count", count)
}
