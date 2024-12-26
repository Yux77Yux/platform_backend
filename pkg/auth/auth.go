package auth

import (
	"fmt"

	"github.com/Yux77Yux/platform_backend/pkg/jwt"
)

var GUEST = []string{
	"get:creation:public", // 查看公开的作品
	"get:comment:public",  // 查看评论
	"get:user:public",     // 查看公开的USER资料
}

func Auth(user_id int64, method, object string, token string) (bool, error) {
	accessToken := token
	owner := "public"
	scope := GUEST

	if accessToken != "" {
		accessClaims, err := jwt.ParseJWT(accessToken)
		if err != nil {
			return false, fmt.Errorf("parseJWT err %w", err)
		}

		if accessClaims.UserID == user_id {
			owner = "own"
		}
		if accessClaims.Role == "ADMIN" {
			owner = "manage"
		}
		if accessClaims.Role == "SUPER_ADMIN" {
			owner = "super"
		}

		scope = accessClaims.Scope
	}

	// 拼接权限
	power := fmt.Sprintf("%s:%s:%s", method, object, owner)

	return includes(scope, power), nil
}

func includes(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}
