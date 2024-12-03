package config

import (
	service "github.com/Yux77Yux/platform_backend/microservices/auth/service"
)

const (
	GRPC_SERVER_ADDRESS = ":50020"
)

func init() {
	service.InitStr(GRPC_SERVER_ADDRESS)
}
