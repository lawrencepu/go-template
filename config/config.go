package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

var mc map[string]json.RawMessage

const DefaultConfigPath = "../config/app.json"

type Config struct {
	BaseConf BaseConf `json:"app"`
	DBConf   DBConf   `json:"mysql"`
	RdbConf  RdbConf  `json:"redis"`
	EsConf   EsConf   `json:"es"`
}

type (
	BaseConf struct {
		AppName      string `json:"app_name"`       //app名称
		AppSecretKey string `json:"app_secret_key"` //app密钥
		AppMode      string `json:"app_mode"`       //app运行环境
		HttpListen   string `json:"http_listen"`    //http监听端口
		RunMode      string `json:"run_mode"`       //运行模式
		AppVersion   string `json:"version"`        //app版本号
		AppUrl       string `json:"app_url"`        //当前路由
		ParamLog     bool   `json:"param_log"`      //是否开启请求参数和返回参数打印
		LogPath      string `json:"log_path"`       //日志路径"/home/log/app"
		Timezone      string `json:"timezone"`       //时区
	}
	DBConf struct {
		Dsn             string `json:"dsn"`             //dsn
		MaxIdleConn     int    `json:"maxIdleConn"`     //空闲连接数
		MaxOpenConn     int    `json:"maxOpenConn"`     //最大连接数
		ConnMaxLifetime int    `json:"connMaxLifetime"` //连接时长
		Prefix 			string `json:"prefix"` //表前缀
	}
	RdbConf struct {
		DB          int    `json:"db"`          //默认连接库
		PoolSize    int    `json:"poolSize"`    //连接数量
		MaxRetries  int    `json:"maxRetries"`  //最大重试次数
		IdleTimeout int    `json:"idleTimeout"` //空闲链接超时时间(单位：time.Second)
		Addr        string `json:"addr"`        //DSN
		Pwd         string `json:"pwd"`         //密码
	}
	EsConf struct {
		Addr []string `json:"addr"`
		User string   `json:"user"`
		Pwd  string   `json:"pwd"`
	}
)

func InitConfig(cfp string) error {
	if len(cfp) <= 0 {
		cfp = DefaultConfigPath
	}
	r, err := ioutil.ReadFile(cfp)
	if err != nil {
		panic(err)
	}
	mc = make(map[string]json.RawMessage)
	err = json.Unmarshal(r, &mc)
	if err != nil {
		log.Printf("%s file load err,err is %s\n", cfp, err.Error())
		panic(err)
	}
	return nil
}

func GetBaseConf() (*BaseConf, error) {
	var bcf = BaseConf{}
	raw, ok := mc["app"]
	if !ok {
		return nil, errors.New("config not found")
	}
	err := json.Unmarshal(raw, &bcf)
	if err != nil {
		return nil, err
	}
	return &bcf, nil
}

func GetDBConf(k string) (*DBConf, error) {
	var dbc = DBConf{}
	raw, ok := mc[k]
	if !ok {
		return nil, errors.New("config not found")
	}
	err := json.Unmarshal(raw, &dbc)
	if err != nil {
		return nil, err
	}
	return &dbc, nil
}

func GetRdbConf(k string) (*RdbConf, error) {
	var rdbc = RdbConf{}
	raw, ok := mc[k]
	if !ok {
		return nil, errors.New("config not found")
	}
	err := json.Unmarshal(raw, &rdbc)
	if err != nil {
		return nil, err
	}
	return &rdbc, nil
}

func GetEsConf(k string) (*EsConf, error) {
	var esc = EsConf{}
	raw, ok := mc[k]
	if !ok {
		return nil, errors.New("config not found")
	}
	err := json.Unmarshal(raw, &esc)
	if err != nil {
		return nil, err
	}
	return &esc, nil
}

func GetConf(k string) (json.RawMessage, error) {
	raw, ok := mc[k]
	if !ok {
		return nil, errors.New("config not found")
	}
	return raw, nil
}
