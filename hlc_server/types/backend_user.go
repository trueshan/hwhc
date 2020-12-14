package types

type Transfer struct {
	Id         int64   `json:"id"`
	UserId     int64   `json:"user_id"`
	Amount     float64 `json:"amount"`
	Address    string  `json:"address"`
	Tx_data    string  `json:"tx_data"`
	Tx_hash    string  `json:"tx_hash"`
	Tx_status  int64   `json:"tx_status"`
	Type       int64   `json:"type"`
	Fee        float64 `json:"fee"`
	Memo       string  `json:"memo"`
	UpdateTime string  `json:"update_time"`
	CoinId     int64   `json:"coin_id"`
	CreateTime string  `json:"create_time"`
	TxDesc     string  `json:"tx_desc"`
	IsShop     int64   `json:"is_shop"`
}

type RechargeRecord struct {
	Id       int64   `json:"id"`
	TxHash   string  `json:"tx_hash"`
	UserId   int64   `json:"user_id"`
	Time     int64   `json:"time"`
	Amount   float64 `json:"amount"`
	Coinname string  `json:"coinname"`
	Coinid   int64   `json:"coinid"`
	Status   int64   `json:"status"`
	Desc     string  `json:"desc"`
	Fee      float64 `json:"fee"`
	Num      int64   `json:"num"`
}

type AdminFlashExchangeRecord struct {
	Id             int64   `json:"id"`
	Num            float64 `json:"num"`         //兑换数量
	UsdtTotal      float64 `json:"usdt_total"`  //qc总额
	CoinName       string  `json:"coin_name"`   //兑换币种名称
	CreateDate     string  `json:"create_date"` //创建时间
	Cost           float64 `json:"cost"`        //手续费
	UserId         int64   `json:"user_id"`
	ExchangeId     int64   `json:"exchange_id"`
	PayCoinName    string  `json:"pay_coin_name"`
	ReturnCoinName string  `json:"return_coin_name"`
}
