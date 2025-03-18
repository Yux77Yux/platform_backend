package tools

import (
	"fmt"
	"strconv"

	"log"

	"google.golang.org/protobuf/types/known/timestamppb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	generated "github.com/Yux77Yux/platform_backend/generated/user"
)

func MapUserByString(result map[string]string) (*generated.User, error) {
	converted := make(map[string]interface{})
	// 将 map[string]string 转换为 map[string]interface{}
	for key, value := range result {
		converted[key] = value
	}
	return MapUser(converted)
}

func MapUser(result map[string]interface{}) (*generated.User, error) {
	statusStr, ok := result["user_status"].(string)
	if !ok {
		return nil, fmt.Errorf("MapUser user_status error")
	}
	genderStr, ok := result["user_gender"].(string)
	if !ok {
		return nil, fmt.Errorf("MapUser user_gender error")
	}

	status := generated.UserStatus(generated.UserStatus_value[statusStr])
	gender := generated.UserGender(generated.UserGender_value[genderStr])

	var bday *timestamppb.Timestamp = nil
	if result["user_bday"] != nil {
		var err error
		bday, err = EnsureTimestampPB(result["user_bday"])
		if err != nil {
			return nil, err
		}
	}

	createdAt, err := EnsureTimestampPB(result["user_created_at"])
	if err != nil {
		return nil, err
	}

	updatedAt, err := EnsureTimestampPB(result["user_updated_at"])
	if err != nil {
		log.Println("error: user_updated_at ", err)
		return nil, err
	}

	followersStr := result["followers"].(string)
	followeesStr := result["followees"].(string)

	followers, err := strconv.ParseInt(followersStr, 10, 64)
	if err != nil {
		return nil, err
	}
	followees, err := strconv.ParseInt(followeesStr, 10, 64)
	if err != nil {
		return nil, err
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
		Followers:     int32(followers),
		Followees:     int32(followees),
	}, nil
}
