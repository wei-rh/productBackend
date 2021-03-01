package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"productBackend/config"
)

type WXLoginResp struct {
	OpenId string			`json:"openid"`
	SessionKey string		`json:"session_key"`
	UnionId string			`json:"unionid"`
	ErrCode int				`json:"errcode"`
	ErrMsg string 			`json:"errmsg"`

}
// 这个函数以 code 作为输入, 返回调用微信接口得到的对象指针和异常情况
func WXLogin(code string) (*WXLoginResp, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	fmt.Println(code)
	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, config.App.APP_ID, config.App.AppSecret, code)
	// 创建http get请求
	resp,err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}
	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}
	return &wxResp, nil
}
