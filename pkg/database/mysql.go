package database

import (
	"context"
	"database/sql"

	errMap "github.com/Yux77Yux/platform_backend/pkg/error"
	utils "github.com/Yux77Yux/platform_backend/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
)

// 内存不够了，把主从复制去掉

type MysqlClass struct {
	mainDB *sql.DB
	// replicaDB *sql.DB
}

func InitDb(readStr string, writeStr string) (SqlMethods, error) {
	db := &MysqlClass{}
	mainDB, err := sql.Open("mysql", writeStr)
	if err != nil {
		utils.LogSuperError(err)
		return nil, err
	}
	db.mainDB = mainDB

	if err = db.mainDB.Ping(); err != nil {
		utils.LogSuperError(err)
		return nil, err
	}

	// if dbs.replicaDB, err = sql.Open("mysql", readStr); err != nil {
	// 	return err
	// }

	// if err = dbs.replicaDB.Ping(); err != nil {
	// 	return err
	// }

	// dbs.replicaDB.SetMaxOpenConns(20)    // 最大打开连接数
	// dbs.replicaDB.SetMaxIdleConns(10)    // 最大空闲连接数
	// dbs.replicaDB.SetConnMaxLifetime(20) // 连接的最大生命周期（秒）
	db.mainDB.SetMaxOpenConns(40)    // 最大打开连接数
	db.mainDB.SetMaxIdleConns(10)    // 最大空闲连接数
	db.mainDB.SetConnMaxLifetime(20) // 连接的最大生命周期（秒）

	return db, nil
}

func (dbs *MysqlClass) Close() error {
	err := dbs.mainDB.Close()
	if err != nil {
		return errMap.MapMySQLErrorToStatus(err)
	}

	return nil
}

// 执行查询操作
func (dbs *MysqlClass) QueryRow(query string, args ...interface{}) *sql.Row {
	return dbs.mainDB.QueryRow(query, args...)
}

func (dbs *MysqlClass) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return dbs.mainDB.QueryRowContext(ctx, query, args...)
}

// 执行查询返回多个结果
func (dbs *MysqlClass) Query(query string, args ...interface{}) (*sql.Rows, error) {
	result, err := dbs.mainDB.Query(query, args...)
	if err != nil {
		return nil, errMap.MapMySQLErrorToStatus(err)
	}
	return result, nil
}

func (dbs *MysqlClass) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	result, err := dbs.mainDB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errMap.MapMySQLErrorToStatus(err)
	}
	return result, nil
}

// 执行插入、更新或删除操作
func (dbs *MysqlClass) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := dbs.mainDB.Exec(query, args...)
	if err != nil {
		return nil, errMap.MapMySQLErrorToStatus(err)
	}
	return result, nil
}

func (dbs *MysqlClass) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := dbs.mainDB.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, errMap.MapMySQLErrorToStatus(err)
	}
	return result, nil
}

// 开始一个新的事务
func (dbs *MysqlClass) BeginTransaction() (*sql.Tx, error) {
	tx, err := dbs.mainDB.Begin()
	if err != nil {
		return nil, errMap.MapMySQLErrorToStatus(err)
	}
	return tx, nil
}

// 提交事务
func (dbs *MysqlClass) CommitTransaction(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return errMap.MapMySQLErrorToStatus(err)
	}
	return nil
}

// 回滚事务
func (dbs *MysqlClass) RollbackTransaction(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return errMap.MapMySQLErrorToStatus(err)
	}
	return nil
}
