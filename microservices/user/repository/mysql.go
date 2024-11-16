package repository

import (
	"database/sql"
	"fmt"
	"log"

	config "github.com/Yux77Yux/platform_backend/microservices/auth/config"
	db "github.com/Yux77Yux/platform_backend/pkg/db"
)

var (
	MysqlReaderConnection *sql.DB
	MysqlWriterConnection *sql.DB
	err                   error
)

func init() {
	MysqlReaderConnection, err = db.OpenMysql(config.MYSQL_READER_STR)
	wired_err := fmt.Errorf("failed to connect the MySQL reader: %w", err)
	log.Printf("error: %v", wired_err)

	MysqlWriterConnection, err = db.OpenMysql(config.MYSQL_WRITER_STR)
	wired_err = fmt.Errorf("failed to connect the MySQL writer: %w", err)
	log.Printf("error: %v", wired_err)
}
