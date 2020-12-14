package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-zhouxun/xmysql"

	_ "github.com/go-sql-driver/mysql"
)

var mysql *MySQL

type MySQL struct {
	database *sql.DB
}

func getInst() *MySQL {
	return mysql
}

func Get() *XMySQL {
	return getInst().NewXMySQL(false)
}

func Begin() *XMySQL {
	return getInst().NewXMySQL(true)
}

func newMySQl(config xmysql.XMySQLConfig) *MySQL {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", config.User, config.Password, config.Address, config.Port, config.DBName)
	db, err := sql.Open("mysql", url)
	db.SetMaxIdleConns(config.MaxIdle)
	db.SetMaxOpenConns(config.MaxConn)
	db.SetConnMaxLifetime(time.Minute * 10)
	if err != nil {
		return nil
	}
	return &MySQL{database: db}
}

func InitDB(config xmysql.XMySQLConfig) {
	mysql = newMySQl(config)
}
