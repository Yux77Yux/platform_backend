package main

import (
	_ "github.com/Yux77Yux/platform_backend/microservices/creation/config"
	service "github.com/Yux77Yux/platform_backend/microservices/creation/service"
)

func main() {
	service.HttpServerRun()
}
