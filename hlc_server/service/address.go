package service

import (
	"fmt"
	"github.com/hwhc/hlc_server/log"
	"github.com/hwhc/hlc_server/persistence"
	"strconv"
	"strings"

	"github.com/hwhc/hlc_server/mysql"
)

func GenAddressPool() bool {
	leftCount := persistence.GetLeftAddress(mysql.Get())
	if leftCount < 1000 {
		coinName := "ETH"
		num := "2000"
		b := CreateAddress(coinName,num)
		log.Info(fmt.Sprintf("【地址检测】创建不足创建地址,leftCount:%d,num:%s,b:%v",leftCount,num,b))
		return b
	}else{
		log.Info(fmt.Sprintf("【地址检测】地址充足,leftCount:%d",leftCount))
		return true
	}
}

func GetUserIdByAddress(coinname string, address string) int64 {
	userId := persistence.GetUserIdByAddress(mysql.Get(), coinname, address)
	return userId
}

func CreateAddress(coinname string, num string) bool {
	return persistence.CreateAddress(mysql.Get(), coinname, num)
}

func GetAddressByMianchain(coinname string, userId int64) string {
	main_name := persistence.GetCoinTokenforName(mysql.Get(), coinname)
	fmt.Println("获取地址 币种主链名称", main_name, "传入参数币种名称", coinname)
	address := persistence.GetAddressByUserId(mysql.Get(), main_name, userId)
	//说明这个用户没有地址 需要给这个用户创建地址

	if address == "" {
		persistence.CreateUserAddress(mysql.Get(), userId, main_name)
		address = persistence.GetAddressByUserId(mysql.Get(), main_name, userId)
		fmt.Print(4, " ", address)
	}

	return address
}

func GetAddressByUserId(coinname string, userId int64) string {
	types := persistence.GetCoinTokenforName(mysql.Get(), coinname)
	//fmt.Println("获取地址 币种主链名称", tokenname, "传入参数币种名称", coinname)
	address := ""
	if strings.ToLower(coinname) == strings.ToLower("EOS") || strings.ToLower(coinname) == strings.ToLower("XRP") {
		address = strconv.FormatInt(userId, 10)
	} else {
		address = persistence.GetAddressByUserId(mysql.Get(), types, userId)
	}
	//说明这个用户没有地址 需要给这个用户创建地址
	if address == "" {
		persistence.CreateUserAddress(mysql.Get(), userId, types)
		address = persistence.GetAddressByUserId(mysql.Get(), types, userId)
		fmt.Print(4, " ", address)
	}

	return address
}
