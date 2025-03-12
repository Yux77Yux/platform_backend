package tools

import (
	"context"

	utils "github.com/Yux77Yux/platform_backend/pkg/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetMetadataValue(ctx context.Context, key string) string {
	return utils.GetMetadataValue(ctx, key)
}

func IsValidVideoURL(url string) bool {
	const urlPattern = `^(https?|ftp)://[^\s]+\.(mp4|avi|mov|mkv|flv|wmv|webm)$`
	return utils.CheckString(url, urlPattern)
}

func IsValidImageURL(url string) bool {
	const urlPattern = `^(https?|ftp)://[^\s]+\.(jpg|jpeg|png|gif|bmp|svg|webp)$`
	return utils.CheckString(url, urlPattern)
}

func CheckStringLength(obj string, min, max int) error {
	return utils.CheckStringLength(obj, min, max)
}

func EnsureTimestampPB(input interface{}) (*timestamppb.Timestamp, error) {
	return utils.EnsureTimestampPB(input)
}
