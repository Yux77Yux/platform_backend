package service

import (
	//generated_common "github.com/Yux77Yux/platform_backend/generated/common"
	generated_user "github.com/Yux77Yux/platform_backend/generated/user"
	"github.com/Yux77Yux/platform_backend/microservices/user/model"
)

func ToModelUserCredentials(user_credentials *generated_user.UserCredentials) *model.UserCredentials {
	return &model.UserCredentials{
		Username: user_credentials.GetUsername(),
		Password: user_credentials.GetPassword(),
	}
}

// func ToProtoUser(user *model.User) *generated_user.User {
// 	return &generated_user.User{
// 		UserDefault: &generated_common.UserDefault{
// 			UserUuid:   user.UserUUID,
// 			UserName:   user.UserName,
// 			UserAvator: user.UserAvatar,
// 		},
// 		UserBio:       user.UserBio,
// 		UserStatus:    user.UserStatus,
// 		UserGender:    user.UserGender,
// 		UserBday:      user.UserBday,
// 		UserCreatedAt: user.CreatedAt,
// 		UserUpdatedAt: user.UpdatedAt,
// 	}
// }
