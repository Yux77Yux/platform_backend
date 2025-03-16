package repository

import (
	"fmt"

	internal "github.com/Yux77Yux/platform_backend/microservices/comment/internal"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/comment/messaging/dispatch"
	receiver "github.com/Yux77Yux/platform_backend/microservices/comment/messaging/receiver"
	tools "github.com/Yux77Yux/platform_backend/microservices/comment/tools"
	pkgDb "github.com/Yux77Yux/platform_backend/pkg/database"
)

var (
	onlyReadStr  string
	readWriteStr string
)

func InitStr(or, wr string) {
	onlyReadStr = or
	readWriteStr = wr
}

func GetDB() (SqlMethod, error) {
	_db, err := pkgDb.InitDb(onlyReadStr, readWriteStr)
	if err != nil {
		return nil, err
	}

	return &SqlMethodStruct{
		db: _db,
	}, nil
}

func Run() func() {
	db, err := GetDB()
	if err != nil {
		tools.LogSuperError(err)
	}

	dispatch.InitDb(db)
	receiver.InitDb(db)
	internal.InitDb(db)

	return func() {
		if err := db.Close(); err != nil {
			tools.LogError("database", "Close", fmt.Errorf("error: database close failed: %w", err))
		}
	}
}
