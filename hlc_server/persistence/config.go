package persistence

import (
	"github.com/hwhc/hlc_server/mysql"
)

const (
	AuditCoinKey  =  "auditCoin"
)

//获取config
func GetConfig(xmysql *mysql.XMySQL, key string) string {
	sqlStr := "select value from config where `key` = ?"
	row := xmysql.QueryRow(sqlStr, key)
	if row == nil {
		return  ""
	}
	var value string
	_ = row.Scan(&value)
	return value
}
