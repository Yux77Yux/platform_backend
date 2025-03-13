package tools

import (
	"context"
	"encoding/base64"
	"errors"
	"regexp"
	"strings"

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
	const urlPattern = `^(https?|ftp)://[^\s]+\.(jpg|jpeg|png|gif|bmp|svg|webp|avif)$`
	return utils.CheckString(url, urlPattern)
}

func CheckStringLength(obj string, min, max int) error {
	return utils.CheckStringLength(obj, min, max)
}

func EnsureTimestampPB(input interface{}) (*timestamppb.Timestamp, error) {
	return utils.EnsureTimestampPB(input)
}

func ParseBase64Image(dataURL string) (fileType string, fileBytes []byte, err error) {
	re := regexp.MustCompile(`^data:image/(.+?);base64,`)
	matches := re.FindStringSubmatch(dataURL)

	if len(matches) != 2 {
		return "", nil, errors.New("invalid image data URL format")
	}

	fileType = matches[1]
	base64Data := strings.Split(dataURL, ",")[1]

	fileBytes, err = base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", nil, errors.New("error decoding Base64 string")
	}

	return fileType, fileBytes, nil
}
