package config

import (
	"os"

	cache "github.com/Yux77Yux/platform_backend/microservices/creation/cache"
	messaging "github.com/Yux77Yux/platform_backend/microservices/creation/messaging"
	db "github.com/Yux77Yux/platform_backend/microservices/creation/repository"
	search "github.com/Yux77Yux/platform_backend/microservices/creation/search"
	service "github.com/Yux77Yux/platform_backend/microservices/creation/service"
)

const (
	MYSQL_WRITER_STR = "yuxyuxx:yuxyuxx@tcp(127.0.0.1:13306)/db_creation_1?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True"
	MYSQL_READER_STR = "yuxyuxx:yuxyuxx@tcp(127.0.0.1:13307)/db_creation_1?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True"

	RABBITMQ_STR = "amqp://yuxyuxx:yuxyuxx@127.0.0.1:5672"

	REDIS_STR = "127.0.0.1:16379"

	GRPC_SERVER_ADDRESS = ":50030"
	HTTP_SERVER_ADDRESS = ":50031"

	SEARCH_HOST = "http://localhost:7700"
	API_KEY     = "yuxyuxx"
)

var REDIS_PASSWORD string = os.Getenv("REDIS_PASSWORD")

func init() {
	search.InitStr(SEARCH_HOST, API_KEY)
	service.InitStr(GRPC_SERVER_ADDRESS, HTTP_SERVER_ADDRESS)
	cache.InitStr(REDIS_STR, REDIS_PASSWORD)
	db.InitStr(MYSQL_READER_STR, MYSQL_WRITER_STR)
	messaging.InitStr(RABBITMQ_STR)
}
