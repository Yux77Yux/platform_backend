package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 内存不够了，把主从复制去掉

type MysqlClass struct {
	mainDB *sql.DB
	// replicaDB *sql.DB
}

func (dbs *MysqlClass) InitDb(readStr string, writeStr string) error {
	var err error

	if dbs.mainDB, err = sql.Open("mysql", writeStr); err != nil {
		return err
	}

	if err = dbs.mainDB.Ping(); err != nil {
		return err
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
	dbs.mainDB.SetMaxOpenConns(40)    // 最大打开连接数
	dbs.mainDB.SetMaxIdleConns(10)    // 最大空闲连接数
	dbs.mainDB.SetConnMaxLifetime(20) // 连接的最大生命周期（秒）

	return nil
}

func (dbs *MysqlClass) Close() error {
	err := dbs.mainDB.Close()
	if err != nil {
		return fmt.Errorf("failed to close main database connection baecause %w", err)
	}

	// err = dbs.replicaDB.Close()
	// if err != nil {
	// 	return fmt.Errorf("failed to close replica database connection baecause %w", err)
	// }

	return nil
}

// 执行查询操作
func (dbs *MysqlClass) QueryRow(query string, args ...interface{}) *sql.Row {
	return dbs.mainDB.QueryRow(query, args...)
	// return dbs.replicaDB.QueryRow(query, args...)
}

func (dbs *MysqlClass) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return dbs.mainDB.QueryRowContext(ctx, query, args...)
	// return dbs.replicaDB.QueryRow(query, args...)
}

// 执行查询返回多个结果
func (dbs *MysqlClass) Query(query string, args ...interface{}) (*sql.Rows, error) {
	result, err := dbs.mainDB.Query(query, args...)
	// result, err := dbs.replicaDB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("query failed because %w", err)
	}
	return result, nil
}

func (dbs *MysqlClass) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	result, err := dbs.mainDB.QueryContext(ctx, query, args...)
	// result, err := dbs.replicaDB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("query failed because %w", err)
	}
	return result, nil
}

// 执行插入、更新或删除操作
func (dbs *MysqlClass) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := dbs.mainDB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec failed because %w", err)
	}
	return result, nil
}

func (dbs *MysqlClass) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := dbs.mainDB.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec failed because %w", err)
	}
	return result, nil
}

// 开始一个新的事务
func (dbs *MysqlClass) BeginTransaction() (*sql.Tx, error) {
	tx, err := dbs.mainDB.Begin()
	if err != nil {
		return nil, fmt.Errorf("open transaction failed because %w", err)
	}
	return tx, nil
}

// 提交事务
func (dbs *MysqlClass) CommitTransaction(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("commit transaction failed because %w", err)
	}
	return nil
}

// 回滚事务
func (dbs *MysqlClass) RollbackTransaction(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return fmt.Errorf("rollback transaction failed because %w", err)
	}
	return nil
}
