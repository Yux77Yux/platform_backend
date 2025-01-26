package config

import (
	internal "github.com/Yux77Yux/platform_backend/microservices/auth/internal"
	service "github.com/Yux77Yux/platform_backend/microservices/auth/service"
)

const (
	GRPC_SERVER_ADDRESS = ":51000"
	HTTP_SERVER_ADDRESS = ":51001"
)

var RoleScopeMapping = map[string][]string{
	"USER": { // 注册用户权限（包括 GUEST 的权限）
		"get:creation:public",
		"get:creation:own",
		"post:creation:own",   // 上传自己的作品
		"update:creation:own", // 更新自己的作品
		"delete:creation:own", // 删除自己的作品

		"get:user:public", // 查看公开的USER资料
		"get:user:own",    // 查看自己的资料
		"update:user:own", // 更新自己的资料
		"delete:user:own", // 更新自己的资料

		"get:comment:public", // 查看评论
		"post:comment:own",   // 发表自己的评论
		"delete:comment:own", // 删除自己的评论

		"post:review:own", // 提交审核信息
	},
	"ADMIN": { // 审核员权限（针对作品和评论的管理）
		"get:creation:manage",
		"update:creation:manage", // 修改作品状态（如通过/拒绝）
		"delete:creation:manage", // 删除作品

		"get:comment:manage",    // 查看评论
		"delete:comment:manage", // 删除评论

		"get:review:manage",    // 查看所有审核信息
		"update:review:manage", // 更新审核状态

		"update:user:manage", // 更新用户信息（如封禁、激活）
	},
	"SUPER_ADMIN": { // 管理员权限（包括公告和用户凭据的管理）
		"post:announcement:super",   // 发布公告
		"update:announcement:super", // 更新公告

		"post:user_credentials:super",   // 创建用户凭据
		"update:user_credentials:super", // 更新用户凭据
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
	service.InitStr(GRPC_SERVER_ADDRESS, HTTP_SERVER_ADDRESS)
	internal.InitScope(RoleScopeMapping)
}
