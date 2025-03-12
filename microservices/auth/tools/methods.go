package tools

import (
	"context"

	utils "github.com/Yux77Yux/platform_backend/pkg/utils"
)

func GetMetadataValue(ctx context.Context, key string) string {
	return utils.GetMetadataValue(ctx, key)
}
