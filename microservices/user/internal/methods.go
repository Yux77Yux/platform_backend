package internal

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
)

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

	bday, err := ensureTimestampPB(result["user_bday"])
	if err != nil {
		log.Println("error: user_bday ", err)
		return nil
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
		UserAvator:    result["user_avator"].(string),
		UserBio:       result["user_bio"].(string),
		UserEmail:     result["user_email"].(string),
		UserStatus:    status,
		UserGender:    gender,
		UserBday:      bday,
		UserCreatedAt: createdAt,
		UserUpdatedAt: updatedAt,
		UserRole:      role,
	}
}
