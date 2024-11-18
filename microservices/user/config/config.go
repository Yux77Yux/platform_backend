package config

import (
	userCache "github.com/Yux77Yux/platform_backend/microservices/user/cache"
	userMQ "github.com/Yux77Yux/platform_backend/microservices/user/messaging"
	userDB "github.com/Yux77Yux/platform_backend/microservices/user/repository"
	"os"
)

const (
	MYSQL_READER_STR = "yuxyuxx:yuxyuxx(127.0.0.1:23306)/User?parseTime=true"
	MYSQL_WRITER_STR = "yuxyuxx:yuxyuxx(127.0.0.1:23307)/User?parseTime=true"

	RABBITMQ_STR = "amqp://yuxyuxx:yuxyuxx@127.0.0.1:5672/"

	REDIS_STR = "redis://127.0.0.1:16379"
)

var REDIS_PASSWORD string = os.Getenv("REDIS_PASSWORD")

func init() {
	userMQ.InitConnStr(RABBITMQ_STR)
	userCache.InitStr(REDIS_STR, REDIS_PASSWORD)
	userDB.InitStr(MYSQL_READER_STR, MYSQL_WRITER_STR)
}
