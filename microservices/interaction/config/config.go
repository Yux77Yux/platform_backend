package config

import (
	"os"

	cache "github.com/Yux77Yux/platform_backend/microservices/interaction/cache"
	messaging "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/interaction/repository"
	service "github.com/Yux77Yux/platform_backend/microservices/interaction/service"
)

const (
	MYSQL_WRITER_STR = "yuxyuxx:yuxyuxx@tcp(127.0.0.1:14306)/information_schema?charset=utf8mb4&parseTime=True"
	MYSQL_READER_STR = "yuxyuxx:yuxyuxx@tcp(127.0.0.1:14307)/information_schema?charset=utf8mb4&parseTime=True"

	RABBITMQ_STR = "amqp://yuxyuxx:yuxyuxx@127.0.0.1:5672"

	REDIS_STR = "127.0.0.1:16379"

	GRPC_SERVER_ADDRESS = ":50040"
)

var REDIS_PASSWORD string = os.Getenv("REDIS_PASSWORD")

func init() {
	service.InitStr(GRPC_SERVER_ADDRESS)
	cache.InitStr(REDIS_STR, REDIS_PASSWORD)
	db.InitStr(MYSQL_READER_STR, MYSQL_WRITER_STR)
	messaging.InitStr(RABBITMQ_STR)
}
