package persistence

import (
	"fmt"
	"github.com/hwhc/hlc_server/log"
	"github.com/hwhc/hlc_server/mysql"
	"github.com/hwhc/hlc_server/types"
	"strconv"
)


func GetExchangeCoinid(mysql *mysql.XMySQL, coinid int64) types.FlashExchange {
	sql := "select id, coin_id, status, limit_status, limit_total, price, purchased, cost ,pay_coin_id , return_coin_id,`buy_status`,real_name,`level_status` from flash_exchange  where `status` = 1 and id = ?"
	rows := mysql.QueryRow(sql, coinid)
	var info types.FlashExchange
	err := rows.Scan(&info.Id, &info.CoinId, &info.Status, &info.LimitStatus, &info.LimitTotal, &info.Price, &info.Purchased, &info.Cost, &info.PayCoinId, &info.ReturnCoinId, &info.BuyStatus, &info.RealName, &info.LevelStatus)
	fmt.Println("info.Cost", info.Cost)
	if err != nil {
		fmt.Println("get swipe failed, %v", err)
		return types.FlashExchange{}
	}

	return info

}

func GetExchangeList(mysql *mysql.XMySQL) []types.FlashExchange {
	sql := "select f.id,coin_id,f.status,limit_status,limit_total,purchased,`name` tokenname , url ,pay_coin_id , return_coin_id , `buy_status`,cost from flash_exchange f WHERE `status` = 1"
	rows, err := mysql.Query(sql)
	if err != nil {
		fmt.Println("getExchangeList failed, %v", err)
		return nil
	}
	list := []types.FlashExchange{}
	for rows.Next() {
		var info types.FlashExchange
		err = rows.Scan(&info.Id, &info.CoinId, &info.Status, &info.LimitStatus, &info.LimitTotal, &info.Purchased, &info.TokenName, &info.Url, &info.PayCoinId, &info.ReturnCoinId, &info.BuyStatus, &info.Cost)
		if err != nil {
			fmt.Println("GetExchangeList failed, %v", err)
		}

		pay_price := GetRealTimePrice(mysql, info.PayCoinId)
		return_price := GetRealTimePrice(mysql, info.ReturnCoinId)
		info.Price = pay_price / return_price

		payCoins := GetCoinInfo(mysql, info.PayCoinId)
		info.PayUrl = payCoins.Icon
		info.PayCoinname = payCoins.SortName
		returnCoins := GetCoinInfo(mysql, info.ReturnCoinId)
		info.ReturnUrl = returnCoins.Icon
		info.ReturnCoinname = returnCoins.SortName

		if info.LimitTotal > 0 {
			purch := Purchased(mysql, info.Id) + info.Purchased
			purch, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", purch), 64)
			info.LimitTotal, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", info.LimitTotal), 64)
			purch = (purch / info.LimitTotal)
			info.Percentage = purch
		}
		list = append(list, info)
	}
	return list
}

func SelUserLevel(xmysql *mysql.XMySQL, userId int64) int64 {
	var level int64
	sql := "SELECT level FROM user WHERE id = ? "
	row := xmysql.QueryRow(sql, userId)
	err := row.Scan(&level)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return level
		}
		fmt.Println("查询用户级别", err)
		log.Error("查询用户级别", err)
		return level
	}
	return level
}
