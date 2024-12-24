package tools

import (
	"fmt"
	// "log"
	"os"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	// common "github.com/Yux77Yux/platform_backend/generated/common"
	// generated "github.com/Yux77Yux/platform_backend/generated/creation"
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
