package http

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/EducationEKT/xserver/x_err"
	"github.com/EducationEKT/xserver/x_http/x_req"
	"github.com/EducationEKT/xserver/x_http/x_resp"
	"github.com/EducationEKT/xserver/x_http/x_router"
	"github.com/hwhc/hlc_server/hoo"
	"github.com/hwhc/hlc_server/log"
	"github.com/hwhc/hlc_server/service"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

func init() {
	x_router.All("/hoo/callback/tx", txCallback)
	x_router.All("/hoo/callback/pushtx", pushtx)
	x_router.All("/hoo/callback/address", addressCallback)
}

type strlist []string

func (list strlist) Len() int {
	return len(list)
}

func (list strlist) Less(i, j int) bool {
	return list[i] < list[j]
}

func (list strlist) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

type TxCallback struct {
	Sign            string `json:"sign"`
	ChainName       string `json:"chain_name"`
	CoinName        string `json:"coin_name"`
	Alias           string `json:"alias"` //币种别名，当币种名称中含有“-”时，此参数值为“-”前的字符，其他情况和币种名称一致，如coin_name=USDT-ERC20时，alias=USDT
	TradType        string `json:"trad_type"`
	BlockHeight     string `json:"block_height"`
	TransactionId   string `json:"transaction_id"`
	TrxN            string `json:"trx_n"`
	Confirmations   string `json:"confirmations"`
	FromAddress     string `json:"from_address"`
	ToAddress       string `json:"to_address"`
	Memo            string `json:"memo"`
	Amount          string `json:"amount"`
	Fee             string `json:"fee"`
	ContractAddress string `json:"contract_address"`
	OuterOrderNo    string `json:"outer_order_no"`
	ConfirmTime     string `json:"confirm_time"`
	Message         string `json:"message"` //消息提示，success为成功，confirming为确认中，其他为失败
}

func (cb TxCallback) check() bool {
	fmt.Println("确认数：%s cb.Message：%s cb.TradType：%s ", cb.Confirmations, cb.Message, cb.TradType)
	//log.Info("确认数：%s", cb.Confirmations)
	confirmations, _ := strconv.ParseInt(cb.Confirmations, 10, 64)
	tradType, _ := strconv.ParseInt(cb.TradType, 10, 64)
	if strings.ToUpper(cb.CoinName) == "BTC" {
		if confirmations < 1 || cb.Message != "success" || tradType != 2 {
			return false
		}
	} else {
		if confirmations < 3 || cb.Message != "success" || tradType != 2 {
			return false
		}
	}

	return true
}

var fail = []byte(`{"code": -1}`)

func txCallback(req *x_req.XReq) (*x_resp.XRespContainer, *x_err.XErr) {
	log.Info("body-hoo %s", string(req.Body))
	//if strings.Split(req.R.RemoteAddr, ":")[0] != "3.113.98.223" {
	//	log.Error("IP地址错误，", req.R.RemoteAddr)
	//	return &x_resp.XRespContainer{
	//		HttpCode: 200,
	//		Body:     fail,
	//	}, nil
	//}

	m, _ := url.ParseQuery(string(req.Body))

	var list strlist = make([]string, 0)
	for key, _ := range m {
		list = append(list, key)
	}
	sort.Sort(list)
	str := ""
	for _, xxx := range list {
		if xxx == "sign" {
			continue
		}

		str = fmt.Sprintf("%s&%s=%v", str, xxx, m[xxx][0])

	}
	str = str[1:]
	fmt.Println("str==", str)
	sign := req.MustGetString("sign")
	fmt.Println("sign==", sign)
	correctSign := hoo.GetHmacCode(str)
	if correctSign != sign {
		log.Error("签名错误： %s = %s", correctSign, m.Get("sign"))
		fmt.Println("签名错误： %s = %s", correctSign, m.Get("sign"))
		return &x_resp.XRespContainer{
			HttpCode: 200,
			Body:     fail,
		}, nil
	}

	fmt.Println(correctSign, "=", sign)

	cb_json, _ := json.Marshal(m) //转json
	param := strings.Replace(string(cb_json), "[", "", -1)
	param = strings.Replace(param, "]", "", -1)

	var cb TxCallback
	d := json.NewDecoder(bytes.NewReader([]byte(param)))
	d.UseNumber()
	_ = d.Decode(&cb)

	if cb.check() {
		fmt.Println("确认数够了。 ")
		//log.Info("确认数够了。 ")
		amount, _ := strconv.ParseFloat(cb.Amount, 64)
		fmt.Println("ToAddress %s TransactionID %s cointype %s", cb.ToAddress, cb.TransactionId, cb.CoinName)
		str = str[1:]
		str = cb.TransactionId + "_" + cb.TrxN
		address := cb.ToAddress
		if strings.ToLower(cb.CoinName) == strings.ToLower("EOS") || strings.ToLower(cb.CoinName) == strings.ToLower("XRP") {
			address = cb.Memo
		}
		service.Recharge(address, amount, str, cb.CoinName)
	} else {
		return &x_resp.XRespContainer{
			HttpCode: 200,
			Body:     fail,
		}, nil
	}
	body := []byte(`{"code": 0}`)
	resp := x_resp.XRespContainer{
		HttpCode: 200,
		Body:     body,
	}
	return &resp, nil

}

func addressCallback(req *x_req.XReq) (*x_resp.XRespContainer, *x_err.XErr) {
	fmt.Println(string(req.Body))
	body := []byte(`{"code": 0}`)
	resp := x_resp.XRespContainer{
		HttpCode: 200,
		Body:     body,
	}
	return &resp, nil
}

func pushtx(req *x_req.XReq) (*x_resp.XRespContainer, *x_err.XErr) {
	fmt.Println(string(req.Body))

	log.Info("提币地址回调：%s", req.Body)

	//log.Info("提币地址回调：%s", req.Body)

	body := []byte(`{"code": 0}`)
	resp := x_resp.XRespContainer{
		HttpCode: 200,
		Body:     body,
	}
	return &resp, nil
}

func computeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sha))
}
