package db

import (
	"database/sql"
)

func OpenMysql(conn_str string) (*sql.DB, error) {
	var (
		mysql_connection *sql.DB
		err              error
	)
	if mysql_connection, err = sql.Open("mysql", conn_str); err != nil {
		return nil, err
	}

	if err = mysql_connection.Ping(); err != nil {
		return nil, err
	}

	mysql_connection.SetMaxOpenConns(10)    // 最大打开连接数
	mysql_connection.SetMaxIdleConns(5)     // 最大空闲连接数
	mysql_connection.SetConnMaxLifetime(20) // 连接的最大生命周期（秒）

	return mysql_connection, nil
}
