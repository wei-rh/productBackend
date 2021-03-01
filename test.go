package main

import (
	"productBackend/model"
)

func main() {
	//url := "https://api.weixin.qq.com/sns/jscode2session?appid=wx8af0d5c69a1b6179&secret=cf65449b3616ccc19f1aba150eb0bffb&js_code=071r5a0003hfeL127B000SlWQC3r5a0X&grant_type=authorization_code"
	//code := "071r5a0003hfeL127B000SlWQC3r5a0X"
	//合成url, 这里的appId和secret是在微信公众平台上获取的
	//url = fmt.Sprintf(url, config.App.APP_ID, config.App.AppSecret, code)

	//创建http get请求
	//resp,_ := http.Get(url)
	//defer resp.Body.Close()
	//
	//wxResp := model.WXLoginResp{}
	//decoder := json.NewDecoder(resp.Body)
	//if err := decoder.Decode(&wxResp); err != nil {
	//	return
	//}
	//fmt.Println(wxResp)
	openid := "oaNGa5W3Lpd5Y33lIiNJgRW67tWc"
	user := &model.User{}
	model.DB.Where("openid = ?",openid).Find(&user)
	//fmt.Println(user)
	//配置读取
	//config.InitConf()
	//tokenStr, _ := middleware.SignWxToken(8,1800)
	//token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
	//	return []byte(config.App.SigninKey), nil
	//})
	//token2, _ := jwt.ParseWithClaims(tokenStr, &middleware.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	//	return []byte(config.App.SigninKey), nil
	//})
	//
	//claims, ok := token2.Claims.(*middleware.MyCustomClaims)
	//if !ok {
	//	fmt.Println("错误")
	//}
	//fmt.Printf("%T",claims)
	//fmt.Println()
	//fmt.Printf("%T",token.Claims)
	//fmt.Println(claims)
	//fmt.Println(token.Claims)

}
