package tools

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	generated "github.com/Yux77Yux/platform_backend/generated/creation"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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

func MapCreationInfoByString(result map[string]string) *generated.CreationInfo {
	converted := make(map[string]interface{})
	// 将 map[string]string 转换为 map[string]interface{}
	for key, value := range result {
		converted[key] = value
	}
	return MapCreationInfo(converted)
}

func MapCreationInfo(result map[string]interface{}) *generated.CreationInfo {
	statusStr := result["status"].(string)

	status := generated.CreationStatus(generated.CreationStatus_value[statusStr])

	var publishTime *timestamppb.Timestamp = nil
	if result["publish_time"] != nil {
		var err error
		publishTime, err = ensureTimestampPB(result["publish_time"])
		if err != nil {
			log.Println("error: publish_time ", err)
			return nil
		}
	}

	uploadTime, err := ensureTimestampPB(result["upload_time"])
	if err != nil {
		log.Println("error: upload_time ", err)
		return nil
	}

	authorId, err := strconv.ParseInt(result["author_id"].(string), 10, 64)
	if err != nil {
		log.Println("authorId is not int64 type")
	}

	duration, err := strconv.Atoi(result["duration"].(string))
	if err != nil {
		log.Println("duration is not int type")
	}

	categoryId, err := strconv.Atoi(result["category_id"].(string))
	if err != nil {
		log.Println("category_id is not int type")
	}

	categoryParent, err := strconv.Atoi(result["category_parent"].(string))
	if err != nil {
		log.Println("category_parent is not int type")
	}

	views, err := strconv.Atoi(result["views"].(string))
	if err != nil {
		log.Println("views is not int type")
	}

	likes, err := strconv.Atoi(result["likes"].(string))
	if err != nil {
		log.Println("likes is not int type")
	}

	saves, err := strconv.Atoi(result["saves"].(string))
	if err != nil {
		log.Println("saves is not int type")
	}

	return &generated.CreationInfo{
		Creation: &generated.Creation{
			BaseInfo: &generated.CreationUpload{
				AuthorId:   authorId,
				Src:        result["src"].(string),
				Thumbnail:  result["thumbnail"].(string),
				Title:      result["title"].(string),
				Bio:        result["bio"].(string),
				Status:     status,
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
			Name:       result["category_name"].(string),
			Parent:     int32(categoryParent),
		},
	}
}
