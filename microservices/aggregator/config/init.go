package config

import (
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
	queueMQ "github.com/Yux77Yux/platform_backend/microservices/aggregator/messaging"
	service "github.com/Yux77Yux/platform_backend/microservices/aggregator/service"
)

const (
	RABBITMQ_STR = "amqp://yuxyuxx:yuxyuxx@127.0.0.1:5672"

	INIT_SERVER_ADDRESS = ":50000" // 聚合层的grpc服务器地址
	SERVER_ADDRESS      = ":8080"  // envoy地址代理
)

func init() {
	service.InitStr(INIT_SERVER_ADDRESS)

	queueMQ.InitStr(RABBITMQ_STR)
	client.InitStr(SERVER_ADDRESS)
}
