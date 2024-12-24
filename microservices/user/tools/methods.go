package tools

import (
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
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

func MapUserByString(result map[string]string) *generated.User {
	converted := make(map[string]interface{})
	// 将 map[string]string 转换为 map[string]interface{}
	for key, value := range result {
		converted[key] = value
	}
	return MapUser(converted)
}

func MapUser(result map[string]interface{}) *generated.User {
	statusStr := result["user_status"].(string)
	genderStr := result["user_gender"].(string)
	roleStr := result["user_role"].(string)

	status := generated.UserStatus(generated.UserStatus_value[statusStr])
	gender := generated.UserGender(generated.UserGender_value[genderStr])
	role := generated.UserRole(generated.UserRole_value[roleStr])

	var bday *timestamppb.Timestamp = nil
	if result["user_bday"] != nil {
		var err error
		bday, err = ensureTimestampPB(result["user_bday"])
		if err != nil {
			log.Println("error: user_bday ", err)
			return nil
		}
	}

	createdAt, err := ensureTimestampPB(result["user_created_at"])
	if err != nil {
		log.Println("error: user_created_at ", err)
		return nil
	}

	updatedAt, err := ensureTimestampPB(result["user_updated_at"])
	if err != nil {
		log.Println("error: user_updated_at ", err)
		return nil
	}

	return &generated.User{
		UserDefault: &common.UserDefault{
			UserName: result["user_name"].(string),
		},
		UserAvatar:    result["user_avatar"].(string),
		UserBio:       result["user_bio"].(string),
		UserStatus:    status,
		UserGender:    gender,
		UserBday:      bday,
		UserCreatedAt: createdAt,
		UserUpdatedAt: updatedAt,
		UserRole:      role,
	}
}
