package repository

import (
	"context"
	"fmt"

	pkgDb "github.com/Yux77Yux/platform_backend/pkg/database"
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

func GetDB() (SqlMethods, error) {
	_db, err := pkgDb.InitDb(onlyReadStr, readWriteStr)
	if err != nil {
		return nil, err
	}
	db = _db
	return db, nil
}

func Run(ctx context.Context) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	<-ctx.Done()
	if err := db.Close(); err != nil {
		return fmt.Errorf("error: database close failed: %w", err)
	}
	return nil
}
