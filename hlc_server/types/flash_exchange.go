package types

type FlashExchange struct {
	Id           int64   `json:"id"`
	CoinId       int64   `json:"coin_id"`
	Status       int64   `json:"status"`
	LimitStatus  int64   `json:"limit_status"`   //限购状态 是否限购1 限购 0 不限购
	LimitTotal   float64 `json:"limit_total"`    //限购总额
	Price        float64 `json:"price"`          //闪购价格
	Purchased    float64 `json:"purchased"`      //已购买数量  用于调整
	Cost         float64 `json:"cost"`           //手续费
	PayCoinId    int64   `json:"pay_coin_id"`    //支付币种id
	ReturnCoinId int64   `json:"return_coin_id"` //购买币种id
	RealName     int64   `json:"real_name"`
	TokenName    string  `json:"token_name"`

	Url       string `json:"url"`
	Name      string `json:"name"`
	PayUrl    string `json:"pay_url"`
	ReturnUrl string `json:"return_url"`

	PayCoinname    string  `json:"pay_coinname"`
	ReturnCoinname string  `json:"return_coinname"`
	Percentage     float64 `json:"percentage"` //百分比
	BuyStatus      int64   `json:"buy_status"`
	LevelStatus    int64   `json:"level_status"`
}
