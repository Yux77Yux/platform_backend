package config

import (
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
	queueMQ "github.com/Yux77Yux/platform_backend/microservices/aggregator/messaging"
	service "github.com/Yux77Yux/platform_backend/microservices/aggregator/service"
)

const (
	RABBITMQ_STR = "amqp://yuxyuxx:yuxyuxx@127.0.0.1:5672"

	GRPC_SERVER_ADDRESS = ":50000" // 聚合层的grpc服务器地址
	USER_SERVER_ADDRESS = ":50010" // 聚合层的grpc服务器地址
	AUTH_SERVER_ADDRESS = ":50020" // 聚合层的grpc服务器地址
)

func init() {
	service.InitStr(GRPC_SERVER_ADDRESS)
	queueMQ.InitStr(RABBITMQ_STR)
	client.InitStr(USER_SERVER_ADDRESS, AUTH_SERVER_ADDRESS)
}
