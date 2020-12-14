package persistence

import (
	"fmt"
	"github.com/go-zhouxun/xutil/xtime"

	"github.com/hwhc/hlc_server/mysql"
)

func RechargeRecord(xmysql *mysql.XMySQL, userId int64, txHash string, amount float64, coinname string, coinid int64, desc string, fee float64) (success bool) {
	sql := "INSERT INTO `recharge_record`(`tx_hash`, `user_id`, `amount`, `time`, `coinname`,coinid, `desc`,`fee`) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := xmysql.Exec(sql, txHash, userId, amount, xtime.Now(), coinname, coinid, desc, fee)
	if err != nil {
		fmt.Print("this record already exist, %s, %v", txHash, err)
		return false
	}
	id, err := result.LastInsertId()
	if err != nil || id <= 0 {
		fmt.Print("this record already exist, %s, %d, %v", txHash, id, err)
		return false
	}
	fmt.Print("添加用户充值记录成功！～" + "\n")
	return true
}

func RechargeRecordStatusAddress(xmysql *mysql.XMySQL, userId int64, txHash string, amount float64, coinname string, coinid int64, desc string, fee float64, status int64, address string) (success bool) {
	sql := "INSERT INTO `recharge_record`(`tx_hash`, `user_id`, `amount`, `time`, `coinname`,coinid, `desc`,`fee`,status,address) VALUES(?, ?, ?, ?, ?, ?, ?, ?,?,?)"
	result, err := xmysql.Exec(sql, txHash, userId, amount, xtime.Now(), coinname, coinid, desc, fee, status, address)
	if err != nil {
		fmt.Print("this record already exist, %s, %v", txHash, err)
		return false
	}
	id, err := result.LastInsertId()
	if err != nil || id <= 0 {
		fmt.Print("this record already exist, %s, %d, %v", txHash, id, err)
		return false
	}
	fmt.Print("添加用户充值记录成功！～" + "\n")
	return true
}

func RechargeRecordStatus(xmysql *mysql.XMySQL, userId int64, txHash string, amount float64, coinname string, coinid int64, desc string, fee float64, status int64) (success bool) {
	sql := "INSERT INTO `recharge_record`(`tx_hash`, `user_id`, `amount`, `time`, `coinname`,coinid, `desc`,`fee`,status) VALUES(?, ?, ?, ?, ?, ?, ?, ?,?)"
	result, err := xmysql.Exec(sql, txHash, userId, amount, xtime.Now(), coinname, coinid, desc, fee, status)
	if err != nil {
		fmt.Print("this record already exist, %s, %v", txHash, err)
		return false
	}
	id, err := result.LastInsertId()
	if err != nil || id <= 0 {
		fmt.Print("this record already exist, %s, %d, %v", txHash, id, err)
		return false
	}
	fmt.Print("添加用户充值记录成功！～" + "\n")
	return true
}
