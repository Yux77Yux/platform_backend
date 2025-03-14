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
	if db != nil {
		return db
	}
	db := &pkgDb.MysqlClass{}
	err := db.InitDb(onlyReadStr, readWriteStr)
	if err != nil {
		log.Printf("error: database init failed: %v", err)
		return nil
	}

	return db
}

func CloseClient() {
	if err := db.Close(); err != nil {
		log.Printf("error: database close failed: %v", err)
	}
}

func Run() {
	db = GetDB()
}
