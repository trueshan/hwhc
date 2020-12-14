package types

type FlashExchangeRecord struct {
	Id         int64   `json:"id"`
	CoinId     int64   `json:"coin_id"`
	Status     int64   `json:"status"`
	Num        float64 `json:"num"`         //兑换数量
	UsdtTotal  float64 `json:"usdt_total"`  //qc总额
	Price      float64 `json:"price"`       //闪购价格
	CoinName   string  `json:"coin_name"`   // 兑换币种名称
	CreateDate string  `json:"create_date"` //创建时间
	Cost       float64 `json:"cost"`

	PayCoinname    string `json:"pay_coinname"`
	ReturnCoinname string `json:"return_coinname"`

	UserId int64 `json:"user_id"`
}
