package repository

import (
	"log"

	pkgDb "github.com/Yux77Yux/platform_backend/pkg/db"
)

var (
	onlyReadStr  string
	readWriteStr string
	db           SqlMethods
)

func InitStr(or, wr string) {
	onlyReadStr = or
	readWriteStr = wr
}

func GetDB() SqlMethods {
	dbs := &pkgDb.MysqlClass{}
	err := dbs.InitDb(onlyReadStr, readWriteStr)
	if err != nil {
		log.Printf("error: database init failed: %v", err)
		return nil
	}

	return dbs
}

func Init() {
	db = GetDB()
}
