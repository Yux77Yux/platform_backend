package config

import (
	"os"

	cache "github.com/Yux77Yux/platform_backend/microservices/aggregator/cache"
	client "github.com/Yux77Yux/platform_backend/microservices/aggregator/client"
	receiver "github.com/Yux77Yux/platform_backend/microservices/aggregator/messaging/receiver"
	service "github.com/Yux77Yux/platform_backend/microservices/aggregator/service"
)

const (
	RABBITMQ_STR = "amqp://yuxyuxx:yuxyuxx@127.0.0.1:5672"

	REDIS_STR = "127.0.0.1:16379"

	INIT_SERVER_ADDRESS = ":50000" // 聚合层的grpc服务器地址
	SERVER_ADDRESS      = ":8080"  // envoy地址代理
)

var REDIS_PASSWORD string = os.Getenv("REDIS_PASSWORD")

func init() {
	service.InitStr(INIT_SERVER_ADDRESS)
	cache.InitStr(REDIS_STR, REDIS_PASSWORD)

	cache.Init()

	client.InitStr(SERVER_ADDRESS)

	receiver.Init(RABBITMQ_STR)
}
