package config

import (
	"os"

	userCache "github.com/Yux77Yux/platform_backend/microservices/user/cache"
	userMQ "github.com/Yux77Yux/platform_backend/microservices/user/messaging"
	userDB "github.com/Yux77Yux/platform_backend/microservices/user/repository"
	service "github.com/Yux77Yux/platform_backend/microservices/user/service"
)

const (
	MYSQL_WRITER_STR = "yuxyuxx:yuxyuxx@tcp(127.0.0.1:13306)/"
	MYSQL_READER_STR = "yuxyuxx:yuxyuxx@tcp(127.0.0.1:13307)/"

	RABBITMQ_STR = "amqp://yuxyuxx:yuxyuxx@127.0.0.1:5672"

	REDIS_STR = "127.0.0.1:16379"

	GRPC_SERVER_ADDRESS = ":50010"
)

var REDIS_PASSWORD string = os.Getenv("REDIS_PASSWORD")

func init() {
	service.InitStr(GRPC_SERVER_ADDRESS)
	userMQ.InitStr(RABBITMQ_STR)
	userCache.InitStr(REDIS_STR, REDIS_PASSWORD)
	userDB.InitStr(MYSQL_READER_STR, MYSQL_WRITER_STR)

	userMQ.Init()
	userDB.Init()
	userCache.Init()
}
