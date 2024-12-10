package main

import (
	_ "github.com/Yux77Yux/platform_backend/microservices/auth/config"
	service "github.com/Yux77Yux/platform_backend/microservices/auth/service"
)

func main() {
	service.HttpServerRun()
}
