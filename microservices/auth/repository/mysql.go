package repository

import (
	"database/sql"
	"fmt"
	"log"

	authConfig "github.com/Yux77Yux/platform_backend/microservices/auth/config"
	pkgDb "github.com/Yux77Yux/platform_backend/pkg/db"
)

var (
	MysqlReaderConnection *sql.DB
	MysqlWriterConnection *sql.DB
	err                   error
)

func init() {
	MysqlReaderConnection, err = pkgDb.OpenMysql(authConfig.MYSQL_READER_STR)

	if err != nil {
		wrappedErr := fmt.Errorf("error: could not connect to MySQL reader: %w", err)
		log.Fatalf("%v", wrappedErr)
	}

	MysqlWriterConnection, err = pkgDb.OpenMysql(authConfig.MYSQL_WRITER_STR)

	if err != nil {
		wrappedErr := fmt.Errorf("error: could not connect to MySQL writer: %w", err)
		log.Fatalf("%v", wrappedErr)
	}
}
