package config

import (
	internal "github.com/Yux77Yux/platform_backend/microservices/auth/internal"
	service "github.com/Yux77Yux/platform_backend/microservices/auth/service"
)

const (
	GRPC_SERVER_ADDRESS = ":50020"
)

var RoleScopeMapping = map[string][]string{
	"GUEST": {
		"read:video", // 仅限观看公开视频
	},
	"USER": {
		"read:video",
		"upload:video",
		"write:comment",
		"read:profile",
		"edit:profile",
		"delete:video", // 删除自己上传的视频
	},
	"ADMIN": {
		"read:video",
		"delete:video:all",
		"moderate:comments",
		"ban:user",
		"read:analytics",
	},
	"SUPER_ADMIN": {
		"*:*", // 全权限
	},
}

func GenerateScope(role string, roleScopeMapping map[string][]string) []string {
	// Check if the role exists in the RoleScopeMapping
	if scopes, exists := roleScopeMapping[role]; exists {
		return scopes
	}
	// Return an empty slice if the role is not found in the mapping
	return []string{}
}

func init() {
	service.InitStr(GRPC_SERVER_ADDRESS)
	internal.InitScope(RoleScopeMapping)
}
