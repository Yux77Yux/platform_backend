package config

import (
	"os"

	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	mq "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"
	oss "github.com/Yux77Yux/platform_backend/microservices/interaction/oss"
	db "github.com/Yux77Yux/platform_backend/microservices/interaction/repository"
	service "github.com/Yux77Yux/platform_backend/microservices/interaction/service"
)

const (
	MYSQL_WRITER_STR = "yuxyuxx:yuxyuxx@tcp(127.0.0.1:14306)/"
	MYSQL_READER_STR = "yuxyuxx:yuxyuxx@tcp(127.0.0.1:14307)/"

	RABBITMQ_STR = "amqp://yuxyuxx:yuxyuxx@127.0.0.1:5672"

	REDIS_STR = "127.0.0.1:16379"

	GRPC_SERVER_ADDRESS = ":50040"
)

var REDIS_PASSWORD string = os.Getenv("REDIS_PASSWORD")

func init() {
	service.InitStr(GRPC_SERVER_ADDRESS)
	mq.InitStr(RABBITMQ_STR)
	cache.InitStr(REDIS_STR, REDIS_PASSWORD)
	db.InitStr(MYSQL_READER_STR, MYSQL_WRITER_STR)

	mq.Init()
	db.Init()
	cache.Init()
	oss.Init()
}
