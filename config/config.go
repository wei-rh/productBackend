package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var App *AppConf

// 配置类型
type AppConf struct {
	AppName string	`json:"APP_NAME"`
	ServiceAddr string 	`json:"SERVICE_ADDR"`
	DSN string 	`json:"DSN"`
	SigninKey string 	`json:"SIGNIN_KEY"`
	APP_ID string 	`json:"APP_ID"`
	AppSecret string 	`json:"AppSecret"`
}

func InitConf() {
	// 配置文件
	confFile := "./conf/app.json"
	// 读
	content, err := ioutil.ReadFile(confFile)
	if err != nil {
		// 配置错误
		log.Println(err)
		// 使用默认配置
	}
	// 使用 JSON 解析 unmarshal
	conf := &AppConf{}
	if err := json.Unmarshal(content, conf); err != nil {
		// 解析配置文件错误
		log.Println(err)
		// 使用默认配置
	}

	App = conf
}
