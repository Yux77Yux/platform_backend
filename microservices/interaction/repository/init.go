package repository

import (
	internal "github.com/Yux77Yux/platform_backend/microservices/interaction/internal"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging/dispatch"
	receiver "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging/receiver"
	tools "github.com/Yux77Yux/platform_backend/microservices/interaction/tools"
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

func GetDB() (SqlInterface, error) {
	_db, err := pkgDb.InitDb(onlyReadStr, readWriteStr)
	if err != nil {
		return nil, err
	}
	return _db, nil
}

func Run() func() {
	db, err := GetDB()
	if err != nil {
		tools.LogSuperError(err)
	}
	methods := &SqlMethodStruct{
		db: db,
	}

	internal.InitDb(methods)
	receiver.InitDb(methods)
	dispatch.InitDb(methods)

	return func() {
		if err := db.Close(); err != nil {
			tools.LogError("Db", "Close", err)
		}
	}
}
