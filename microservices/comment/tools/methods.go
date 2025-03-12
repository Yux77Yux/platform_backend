package tools

import (
	"context"

	utils "github.com/Yux77Yux/platform_backend/pkg/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CheckStringLength(obj string, min, max int) error {
	return utils.CheckStringLength(obj, min, max)
}

func GetMetadataValue(ctx context.Context, key string) string {
	return utils.GetMetadataValue(ctx, key)
}

func EnsureTimestampPB(input interface{}) (*timestamppb.Timestamp, error) {
	return utils.EnsureTimestampPB(input)
}
