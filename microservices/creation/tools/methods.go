package tools

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	utils "github.com/Yux77Yux/platform_backend/pkg/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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

func GetSpaceCreationCountType(byWhat generated.ByCount) string {
	typeStr := ""
	switch byWhat {
	case generated.ByCount_VIEWS:
		typeStr = "ByViews"
	case generated.ByCount_LIKES:
		typeStr = "ByLikes"
	case generated.ByCount_COLLECTIONS:
		typeStr = "ByCollections"
	default:
		typeStr = "ByPublished_Time"
	}
	return typeStr
}

func GetUserCreationsCountType(byWhat generated.ByCount) string {
	typeStr := ""
	switch byWhat {
	case generated.ByCount_VIEWS:
		typeStr = "views"
	case generated.ByCount_LIKES:
		typeStr = "likes"
	case generated.ByCount_COLLECTIONS:
		typeStr = "saves"
	default:
		typeStr = "publish_time"
	}
	return typeStr
}

func SaveImage(fileBytes []byte, fileName string) error {
	// 指定保存路径和文件名
	filePath := fmt.Sprintf("./%s.png", fileName) // 保存为 PNG 格式
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// 写入字节到文件
	_, err = file.Write(fileBytes)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	fmt.Printf("Image saved successfully at %s\n", filePath)
	return nil
}

func ParseTimestamp(field string) (*timestamppb.Timestamp, error) {
	if field == "none" {
		return nil, nil
	}

	result, err := time.Parse(time.RFC3339, field)
	if err != nil {
		return nil, fmt.Errorf("invalid format: %v", err)
	}
	return timestamppb.New(result), nil
}

func EnsureTimestampPB(input interface{}) (*timestamppb.Timestamp, error) {
	switch v := input.(type) {
	case string:
		if v == "none" {
			return nil, nil
		}
		// 尝试解析字符串为 time.Time
		parsedTime, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return nil, fmt.Errorf("failed to parse string as timestamp: %v", err)
		}
		return timestamppb.New(parsedTime), nil
	case *timestamppb.Timestamp:
		// 如果已经是 *timestamppb.Timestamp 类型，直接返回
		return v, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", v)
	}
}

func ensureTimestampPB(input interface{}) (*timestamppb.Timestamp, error) {
	switch v := input.(type) {
	case string:
		if v == "none" {
			return nil, nil
		}
		// 尝试解析字符串为 time.Time
		parsedTime, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return nil, fmt.Errorf("failed to parse string as timestamp: %v", err)
		}
		return timestamppb.New(parsedTime), nil
	case *timestamppb.Timestamp:
		// 如果已经是 *timestamppb.Timestamp 类型，直接返回
		return v, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", v)
	}
}

func MapCreationInfoByString(result map[string]string) (*generated.CreationInfo, error) {
	converted := make(map[string]interface{})
	// 将 map[string]string 转换为 map[string]interface{}
	for key, value := range result {
		converted[key] = value
	}
	return MapCreationInfo(converted)
}

func MapCreationInfo(result map[string]interface{}) (*generated.CreationInfo, error) {
	requiredKeys := []string{
		"author_id", "src", "thumbnail", "title", "bio", "duration", "upload_time", "status",
		"views", "saves", "likes", "publish_time",
		"category_id", "category_name", "category_parent",
	}
	for _, key := range requiredKeys {
		if val, ok := result[key]; !ok || val == "" {
			return nil, nil
		}
	}

	statusStr, ok := result["status"].(string)
	if !ok {
		return nil, errors.New("missing or invalid 'status'")
	}

	status, exists := generated.CreationStatus_value[statusStr]
	if !exists {
		return nil, errors.New("invalid 'status' value")
	}

	var publishTime *timestamppb.Timestamp
	if v, exists := result["publish_time"]; exists && v != nil {
		var err error
		publishTime, err = ensureTimestampPB(v)
		if err != nil {
			return nil, fmt.Errorf("error parsing 'publish_time': %w", err)
		}
	}

	uploadTime, err := ensureTimestampPB(result["upload_time"])
	if err != nil {
		return nil, fmt.Errorf("error parsing 'upload_time': %w", err)
	}

	authorId, err := parseInt64(result["author_id"], "author_id")
	if err != nil {
		return nil, err
	}

	duration, err := parseInt(result["duration"], "duration")
	if err != nil {
		return nil, err
	}

	categoryId, err := parseInt(result["category_id"], "category_id")
	if err != nil {
		return nil, err
	}

	categoryParent, err := parseInt(result["category_parent"], "category_parent")
	if err != nil {
		return nil, err
	}

	views, err := parseInt(result["views"], "views")
	if err != nil {
		return nil, err
	}

	likes, err := parseInt(result["likes"], "likes")
	if err != nil {
		return nil, err
	}

	saves, err := parseInt(result["saves"], "saves")
	if err != nil {
		return nil, err
	}

	return &generated.CreationInfo{
		Creation: &generated.Creation{
			BaseInfo: &generated.CreationUpload{
				AuthorId:   authorId,
				Src:        safeString(result, "src"),
				Thumbnail:  safeString(result, "thumbnail"),
				Title:      safeString(result, "title"),
				Bio:        safeString(result, "bio"),
				Status:     generated.CreationStatus(status),
				Duration:   int32(duration),
				CategoryId: int32(categoryId),
			},
			UploadTime: uploadTime,
		},
		CreationEngagement: &generated.CreationEngagement{
			Views:       int32(views),
			Likes:       int32(likes),
			Saves:       int32(saves),
			PublishTime: publishTime,
		},
		Category: &generated.Category{
			CategoryId: int32(categoryId),
			Name:       safeString(result, "category_name"),
			Parent:     int32(categoryParent),
		},
	}, nil
}

func parseInt(value interface{}, fieldName string) (int, error) {
	str, ok := value.(string)
	if !ok {
		return 0, fmt.Errorf("missing or invalid '%s'", fieldName)
	}
	return strconv.Atoi(str)
}

func parseInt64(value interface{}, fieldName string) (int64, error) {
	str, ok := value.(string)
	if !ok {
		return 0, fmt.Errorf("missing or invalid '%s'", fieldName)
	}
	return strconv.ParseInt(str, 10, 64)
}

func safeString(result map[string]interface{}, key string) string {
	if val, ok := result[key].(string); ok {
		return val
	}
	return ""
}
