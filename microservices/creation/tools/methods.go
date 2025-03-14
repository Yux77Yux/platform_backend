package tools

import (
	"context"
	"encoding/base64"
	"errors"
	"net/url"
	"regexp"
	"strings"

	utils "github.com/Yux77Yux/platform_backend/pkg/utils"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetMetadataValue(ctx context.Context, key string) string {
	return utils.GetMetadataValue(ctx, key)
}

func GetUuid() uuid.UUID {
	return utils.GetUuid()
}

func GetUuidString() string {
	return utils.GetUuidString()
}

func IsValidVideoURL(videoURL string) bool {
	parsedURL, err := url.Parse(videoURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}

	// Supported video file extensions
	allowedExtensions := []string{
		".mp4", ".avi", ".mov", ".mkv", ".flv", ".wmv", ".webm",
		".3gp", ".ogv", ".m4v", ".ts", ".vob", ".rmvb", ".asf", ".mpeg", ".mpg",
	}

	// Check if the URL path ends with one of the allowed extensions
	for _, ext := range allowedExtensions {
		if strings.HasSuffix(strings.ToLower(parsedURL.Path), ext) {
			return true
		}
	}
	return false
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

func LogSuperError(err error) {
	utils.LogSuperError(err)
}
func LogError(traceId, fullName string, err error) {
	utils.LogError(traceId, fullName, err)
}
func LogInfo(traceId, fullName string) {
	utils.LogInfo(traceId, fullName)
}
func LogWarning(traceId, fullName, warning string) {
	utils.LogWarning(traceId, fullName, warning)
}

func GetMainValue(ctx context.Context) string {
	return utils.GetMainValue(ctx)
}

func GetSnowId() int64 {
	return utils.GetId()
}
