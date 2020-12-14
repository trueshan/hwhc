package persistence

import (
	"fmt"
	"github.com/go-zhouxun/xutil/xtime"
	"github.com/hwhc/hlc_server/mysql"
)

func GetUserIdByToken(xmysql *mysql.XMySQL, token string) int64 {
	sql := "SELECT `user_id` FROM `token` WHERE `token` = ? AND `status` = 0 AND `create_time` > ? ORDER BY `id` DESC LIMIT 1"
	row := xmysql.QueryRow(sql, token, 0)
	id := int64(-1)
	_ = row.Scan(&id)
	return id
}

func GetUserIdByGameToken(xmysql *mysql.XMySQL, token string) int64 {
	sql := "SELECT `user_id` FROM `game_token` WHERE `token` = ?"
	row := xmysql.QueryRow(sql, token)
	id := int64(-1)
	err := row.Scan(&id)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return id
}

func GetUserIdByOTCToken(xmysql *mysql.XMySQL, token string) int64 {
	sql := "SELECT `user_id` FROM `otc_token` WHERE `token` = ?"
	row := xmysql.QueryRow(sql, token)
	id := int64(-1)
	err := row.Scan(&id)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return id
}

func SaveToken(xmysql *mysql.XMySQL, token string, userId int64) {
	//fmt.Println("userId",userId)
	sql := "UPDATE `token` SET `status` = -1 WHERE `user_id` = ?"
	_, _ = xmysql.Exec(sql, userId)
	sql = "INSERT INTO `token`(`token`, `user_id`, `create_time`) VALUES(?, ?, ?)"
	_, _ = xmysql.Exec(sql, token, userId, xtime.Now())
}

func DeleteToken(mysql *mysql.XMySQL, token string) {
	sql := "UPDATE `token` SET `status` = -1 WHERE `token` = ?"
	_, _ = mysql.Exec(sql, token)
}
