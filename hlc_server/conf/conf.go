package conf

import (
	"github.com/go-zhouxun/xmysql"

	"github.com/EducationEKT/EKT/core/types"
)

type Config struct {
	Env            string              `json:"env"`
	Port           int                 `json:"port"`
	Log            string              `json:"log"`
	GethUrl        string              `json:"geth"`
	EthMainAccount types.HexBytes      `json:"eth_main_account"`
	ETHPrivateKey  types.HexBytes      `json:"ethPrivateKey"`
	MysqlConfig    xmysql.XMySQLConfig `json:"mysql"`
	MyredisConfig  RedisConfig         `json:"redis"`
	TemplatePath   string              `json:"templatePath"`
	OtcBaseUrl     string              `json:"otc_base_url"`
	AliEndpoint    string              `json:"ali_endpoint"`
	AliKeyID       string              `json:"ali_key_id"`
	AliKeySecret   string              `json:"ali_key_secret"`
	AliBucket      string              `json:"ali_bucket"`
}

type RedisConfig struct {
	Address     string `json:"address"`
	DBNum       int    `json:"dbNum"`
	Password    string `json:"password"`
	PoolSize    int    `json:"PoolSize"`
	MaxRetries  int    `json:"MaxRetries"`
	IdleTimeout int    `json:"IdleTimeout"`
}

var config Config

func SetConfig(_config Config) {
	config = _config
}

func GetConfig() Config {
	return config
}

//星球收益倍数
var StartBallEarnings = float64(3)
