package config

import (
	"os"

	cache "github.com/Yux77Yux/platform_backend/microservices/comment/cache"
	messaging "github.com/Yux77Yux/platform_backend/microservices/comment/messaging"
	oss "github.com/Yux77Yux/platform_backend/microservices/comment/oss"
	db "github.com/Yux77Yux/platform_backend/microservices/comment/repository"
	service "github.com/Yux77Yux/platform_backend/microservices/comment/service"
)

const (
	MYSQL_WRITER_STR = "yuxyuxx:yuxyuxx@tcp(127.0.0.1:12306)/"
	MYSQL_READER_STR = "yuxyuxx:yuxyuxx@tcp(127.0.0.1:12307)/"

	RABBITMQ_STR = "amqp://yuxyuxx:yuxyuxx@127.0.0.1:5672"

	REDIS_STR = "127.0.0.1:16379"

	GRPC_SERVER_ADDRESS = ":50040"
)

var REDIS_PASSWORD string = os.Getenv("REDIS_PASSWORD")

func init() {
	service.InitStr(GRPC_SERVER_ADDRESS)
	messaging.InitStr(RABBITMQ_STR)
	cache.InitStr(REDIS_STR, REDIS_PASSWORD)
	db.InitStr(MYSQL_READER_STR, MYSQL_WRITER_STR)

	messaging.Init()
	db.Init()
	cache.Init()
	oss.Init()
}
