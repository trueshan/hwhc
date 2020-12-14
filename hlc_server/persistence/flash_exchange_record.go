package persistence

import (
	"fmt"
	"github.com/go-zhouxun/xutil/xtime"
	"github.com/hwhc/hlc_server/mysql"
	"github.com/hwhc/hlc_server/types"
)

//闪兑
func GetExchangeRecordList(mysql *mysql.XMySQL, userid int64) (int, []types.FlashExchangeRecord) {
	sql := "select `coin_id`,`status`,`num`,`usdt_total`,`price`,`coin_name`,`create_date`,user_id,`cost`,pay_coin_name,return_coin_name from flash_exchange_record  where `user_id` = ? order by id desc "
	rows, err := mysql.Query(sql, userid)
	if err != nil {
		fmt.Println("getExchangeList failed, %v", err)
		return -1002, nil
	}
	list := []types.FlashExchangeRecord{}
	for rows.Next() {
		var info types.FlashExchangeRecord
		err = rows.Scan(&info.CoinId, &info.Status, &info.Num, &info.UsdtTotal, &info.Price, &info.CoinName, &info.CreateDate, &info.UserId, &info.Cost, &info.PayCoinname, &info.ReturnCoinname)
		if err != nil {
			fmt.Println("get swipe failed, %v", err)
		}
		list = append(list, info)
	}
	return 200, list
}

func AddExchangeRecord(mysql *mysql.XMySQL, coinid int64, num float64, qcTotal float64, price float64, coinname string, user_id int64, cost float64, payCoidName string, returnCoinName string, exchangeId int64) bool {
	sql := "insert into `flash_exchange_record` (`coin_id`,`status`,`num`,`usdt_total`,`price`,`coin_name`,`create_date`,user_id,date,cost,pay_coin_name,return_coin_name,exchange_id) values (?,2,?,?,?,?,?,?,?,?,?,?,?);"
	result, err := mysql.Exec(sql, coinid, num, qcTotal, price, coinname, xtime.TodayDateTimeStr(), user_id, xtime.TodayDateStr(), cost, payCoidName, returnCoinName, exchangeId)

	if err != nil {

		fmt.Print("addExchangeRecord failed, %v, %d", err)
	}
	if id, err := result.LastInsertId(); err != nil || id <= 0 {
		fmt.Print("addExchangeRecord failed, %v, %d", err, id)

		return false
	}
	return true

}

func Purchased(mysql *mysql.XMySQL, coinid int64) float64 {
	var purch float64
	sql := "SELECT IFNULL(SUM(num),0) from flash_exchange_record where exchange_id = ? AND `date` = ?"
	rows := mysql.QueryRow(sql, coinid, xtime.TodayDateStr())
	err := rows.Scan(&purch)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return purch
		}
		fmt.Println("Purchased failed, %v", err)
	}
	return purch
}

func PurchasedUserID(mysql *mysql.XMySQL, coinid int64, userId int64) float64 {
	var purch float64
	sql := "SELECT IFNULL(SUM(num),0) from flash_exchange_record where exchange_id = ? AND user_id = ?"
	rows := mysql.QueryRow(sql, coinid, userId)
	err := rows.Scan(&purch)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return purch
		}
		fmt.Println("Purchased failed, %v", err)
	}
	return purch
}

func GetTotalUserID(mysql *mysql.XMySQL, coinid int64, userId int64) float64 {
	var purch float64
	sql := "SELECT IFNULL(SUM(usdt_total),0) from flash_exchange_record where exchange_id = ? AND user_id = ?"
	rows := mysql.QueryRow(sql, coinid, userId)
	err := rows.Scan(&purch)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return purch
		}
		fmt.Println("Purchased failed, %v", err)
	}
	return purch
}
