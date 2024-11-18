package repository

import (
	"log"

	db "github.com/Yux77Yux/platform_backend/pkg/db"
)

var (
	onlyReadStr  string
	readWriteStr string
)

func InitStr(or, wr string) {
	onlyReadStr = or
	readWriteStr = wr
}

func GetDB() SqlMethods {
	dbs := &db.MysqlClass{}
	err := dbs.InitDb(onlyReadStr, readWriteStr)
	if err != nil {
		log.Printf("error: database init failed: %v", err)
		return nil
	}

	return dbs
}
