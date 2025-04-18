package config

import (
	"os"

	messaging "github.com/Yux77Yux/platform_backend/microservices/review/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/review/repository"
	service "github.com/Yux77Yux/platform_backend/microservices/review/service"
)

const (
	MYSQL_WRITER_STR = "yuxyuxx:yuxyuxx@tcp(127.0.0.1:15306)/information_schema?charset=utf8mb4&parseTime=True"
	MYSQL_READER_STR = "yuxyuxx:yuxyuxx@tcp(127.0.0.1:15307)/information_schema?charset=utf8mb4&parseTime=True"

	RABBITMQ_STR = "amqp://yuxyuxx:yuxyuxx@127.0.0.1:5672"

	GRPC_SERVER_ADDRESS = ":50050"
)

var REDIS_PASSWORD string = os.Getenv("REDIS_PASSWORD")

func init() {
	service.InitStr(GRPC_SERVER_ADDRESS)
	db.InitStr(MYSQL_READER_STR, MYSQL_WRITER_STR)
	messaging.InitStr(RABBITMQ_STR)
}
