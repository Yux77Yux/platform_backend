package auth

import (
	"fmt"
	"log"

	"github.com/Yux77Yux/platform_backend/pkg/jwt"
)

var GUEST = []string{
	"get:creation:public", // 查看公开的作品
	"get:comment:public",  // 查看评论
	"get:user:public",     // 查看公开的USER资料
}

func Auth(method, object string, token string) (bool, int64, error) {
	accessToken := token
	owner := "public"
	scope := GUEST
	var id int64

	if accessToken != "" {
		log.Println(accessToken)
		accessClaims, err := jwt.ParseJWT(accessToken)
		if err != nil {
			return false, -1, fmt.Errorf("parseJWT err %w", err)
		}

		owner = "own"

		if accessClaims.Role == "ADMIN" {
			owner = "manage"
		}
		if accessClaims.Role == "SUPER_ADMIN" {
			owner = "super"
		}

		scope = accessClaims.Scope
		id = accessClaims.UserID
	}

	// 拼接权限
	power := fmt.Sprintf("%s:%s:%s", method, object, owner)

	return includes(scope, power), id, nil
}

// return (pass，isADMIN，userId，error)
func AuthRole(method, object string, token string) (bool, bool, int64, error) {
	accessToken := token
	owner := "public"
	scope := GUEST
	var id int64

	if accessToken != "" {
		accessClaims, err := jwt.ParseJWT(accessToken)
		if err != nil {
			return false, false, -1, fmt.Errorf("parseJWT err %w", err)
		}

		owner = "own"

		if accessClaims.Role == "ADMIN" {
			owner = "manage"
		}
		if accessClaims.Role == "SUPER_ADMIN" {
			owner = "super"
		}

		scope = accessClaims.Scope
		id = accessClaims.UserID
	}

	// 拼接权限
	power := fmt.Sprintf("%s:%s:%s", method, object, owner)

	return includes(scope, power), owner == "manage", id, nil
}

func includes(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}
