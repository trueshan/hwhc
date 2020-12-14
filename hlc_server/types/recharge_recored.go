package types

type User_recharge struct {
	Tokenname string `json:"tokename"`
	Time      int64  `json:"time"`
	Amount    int64  `json:"amount"`
}

type WalletRechargeRecord struct {
	CoinId int64   `json:"coin_id"`
	Amount float64 `json:"amount"`
	Time   int64   `json:"time"`
	Status int64   `json:"status"`
	Txhash string  `json:"txhash"`
	UserId int64   `json:"user_id"`
	Desc   string  `json:"desc"`
	Fee    float64 `json:"fee"`
}

type WalletRechargeRecord_call struct {
	CoinId int64   `json:"coinid"`
	Amount float64 `json:"amount"`
	Time   int64   `json:"time"`
	Status int64   `json:"status"`
	Txhash string  `json:"tx_hash"`
	UserId int64   `json:"user_id"`
}
