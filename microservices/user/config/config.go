package config

import "os"

const (
	MYSQL_READER_STR = "yuxyuxx:yuxyuxx(127.0.0.1:23306)/Auth?parseTime=true"
	MYSQL_WRITER_STR = "yuxyuxx:yuxyuxx(127.0.0.1:23307)/Auth?parseTime=true"

	RABBITMQ_STR = "amqp://yuxyuxx:yuxyuxx@127.0.0.1:5672/"

	REDIS_STR = "redis://127.0.0.1:16379"
)

var REDIS_PASSWORD string = os.Getenv("REDIS_PASSWORD")
