package types

import (
	"encoding/json"
)

const (
	RECHARGE                = 1  //充值
	TRANSFER                = 2  //提现
	IN_TRANSFER             = 3  //内部转账
	FLASH_EXCHANGE          = 4  //闪兑
	SHOP                    = 5  //购物消费
	UPGRADE                 = 6  //晋升会员
	WELFARE                 = 7  //每日福利
	Fee                     = 8  //手续费
	MP                      = 9  //新增加的业务类型  忘记了
	AP                      = 10 //新增加的业务类型  忘记了
	a                       = 11
	BUY_ORDER               = 12 //推荐会员消费
	REMD_SHOP_SELL          = 13 //推荐商家销售
	REMD_SHOP_SELL_CITY     = 14 //本市商家销售
	REMD_SHOP_SELL_PROVINCE = 15 //本省商家销售
	COMUNITY_PROFIT         = 16 //推荐社区收益
	PINJGJI_COMUNITY        = 17 //推荐平级社区收益
	AMOUNT                  = "AMOUNT"
	FROZEN_AMOUNT           = "FROZEN_AMOUNT"

	SYSTEM_SCRIPT_FREE_RECHARGE        = 22 //系统脚本充值小人币种
	OTC_TRANSFER                       = 38 //otc提现
	REDUCE_SYSTEM_SCRIPT_FREE_RECHARGE = 50 //一定时间未使用，扣除系统脚本充值小人币种

)

type TXData struct {
	Method string
	To     string
	Amount int64
}

type ETHTransaction struct {
	From        string `json:"from"`
	Gas         string `json: "gas"`
	GasPrice    string `json: "gasPrice"`
	Hash        string `json: "hash"`
	Input       string `json: "input"`
	Nonce       string `json: "nonce"`
	To          string `json: "to"`
	Value       string `json: "value"`
	ParsedValue TXData `json:"-"`
	R           string `json: "r"`
	S           string `json: "s"`
	V           string `json: "v"`
}

func (tx ETHTransaction) Bytes() []byte {
	data, _ := json.Marshal(tx)
	return data
}
