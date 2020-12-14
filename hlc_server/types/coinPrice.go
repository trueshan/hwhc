package types

type CoinPirce struct {
	SortName string  `json:"sortName"`
	Price    float64 `json:"price"`
	Icon     string  `json:"icon"`
	CoinName string  `json:"coinName"`
	Rate     float64 `json:"rate"`
}

type CoinInfo struct {
	Icon     string `json:"icon"`
	CoinName string `json:"coinName"`
}

type RespForBY struct {
	Code int64   `json:"code"`
	Msg  string  `json:"msg"`
	Time int64   `json:"time"`
	Buy  float64 `json:"buy"`
	Sell float64 `json:"sell"`
	Data struct {
		Buy  float64 `json:"buy"`
		High float64 `json:"high"`
		Last float64 `json:"last"`
		Low  float64 `json:"low"`
		Sell float64 `json:"sell"`
		Vol  float64 `json:"vol"`
	} `json:"data"`
}

type RespHc struct {
	Date   string `json:"date"`
	Ticker struct {
		Buy  string `json:"buy"`
		High string `json:"high"`
		Last string `json:"last"`
		Low  string `json:"low"`
		Sell string `json:"sell"`
		Vol  string `json:"vol"`
	} `json:"ticker"`
}

type RespMOFOK struct {
	Last string `json:"last"`
}

type RespMOF struct {
	Date   string `json:"date"`
	Ticker struct {
		High string `json:"high"`
		Vol  string `json:"vol"`
		Last string `json:"last"`
		Low  string `json:"low"`
		Buy  string `json:"buy"`
		Sell string `json:"sell"`
	} `json:"ticker"`
}

type RespUsdt struct {
	Date   string `json:"date"`
	Ticker struct {
		Buy  string `json:"buy"`
		High string `json:"high"`
		Last string `json:"last"`
		Low  string `json:"low"`
		Sell string `json:"sell"`
		Vol  string `json:"vol"`
	} `json:"ticker"`
}

type RespHuoBi struct {
	/**
		"status":"ok",
	    "ch":"market.btcusdt.detail.merged",
	    "ts":1581435911073,
	    "tick":{
	        "amount":22505.637814070833,
	        "open":9888.76,
	        "close":10149.53,
	        "high":10156,
	        "id":209004695425,
	        "count":258662,
	        "low":9717.24,
	        "version":209004695425,
	        "ask":[
	            10151.02,
	            0.060115
	        ],
	        "vol":221698283.5713691,
	        "bid":[
	            10150.01,
	            1.938716
	        ]
	    }

	*/
	Status string `json:"status"`
	Ch     string `json:"ch"`
	Ts     int64  `json:"ts"`
	Tick   struct {
		Amount  float64   `json:"amount"`
		Open    float64   `json:"open"`
		Close   float64   `json:"close"`
		High    float64   `json:"high"`
		Id      int64     `json:"id"`
		Count   int64     `json:"count"`
		Low     float64   `json:"low"`
		Version int64     `json:"version"`
		Ask     []float64 `json:"ask"`
		Vol     float64   `json:"vol"`
		Bid     []float64 `json:"bid"`
	} `json:"tick"`
}

type RespWbf struct {
	//{"code":"0","msg":"suc","data":{
	//"high":"0.038229","
	//vol":"58575320.8733980299110987",
	//"last":"0.03725",
	//"low":"0.03386002",
	//"buy":0.03702034,
	//"sell":0.03725,
	//"rose":"0.0988200589970501",
	//"time":null}}
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		High string  `json:"high"`
		Vol  string  `json:"vol"`
		Last string  `json:"last"` //最新成交价
		Low  string  `json:"low"`
		Buy  float64 `json:"buy"`
		Sell float64 `json:"sell"`
		Rose string  `json:"rose"` //涨幅
		Time string  `json:"time"`
	} `json:"data"`
}

type Dcoin struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Time  int64    `json:"time"`
		Open  string   `json:"open"`
		Close string   `json:"close"`
		High  string   `json:"high"`
		Low   string   `json:"low"`
		Vol   string   `json:"vol"`
		Buy   []string `json:"buy"`
		Sell  []string `json:"sell"`
	} `json:"data"`
}

type Okex struct {
	BestAsk        string `json:"best_ask"`
	BestBid        string `json:"best_bid"`
	InstrumentId   string `json:"instrument_id"`
	ProductId      string `json:"product_id"`
	Last           string `json:"last"`
	LastQty        string `json:"last_qty"`
	Ask            string `json:"ask"`
	BestAskSize    string `json:"best_ask_size"`
	Bid            string `json:"bid"`
	BestBidSize    string `json:"best_bid_size"`
	Open24h        string `json:"open_24h"`
	High24h        string `json:"high_24h"`
	Low24h         string `json:"low_24h"`
	BaseVolume24h  string `json:"base_volume_24h"`
	Timestamp      string `json:"timestamp"`
	QuoteVolume24h string `json:"quote_volume_24h"`
}

type RedisStock struct {
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	LatestPrice float64 `json:"latest_price"`
	UdPercent   float64 `json:"ud_percent"`
	UdValue     float64 `json:"ud_value"`
	Type        int64   `json:"type"`
}
