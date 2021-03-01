package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"productBackend/middleware"
	"productBackend/model"
)

// /wechat/applet_login?code=xxx [get]  路由
// 微信小程序登录
func AppletWeChatLogin(ctx *gin.Context) {
	code := ctx.Query("code")     //  获取code
	// 根据code获取 openID 和 session_key
	wxLoginResp,err := model.WXLogin(code)
	if err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
		return
	}
	user := &model.User{}
	DB.Where("openid = ?",wxLoginResp.OpenId).Find(&user)
	if user.ID == 0 {
		//未找到该记录,添加
		user.Openid = wxLoginResp.OpenId
		//插入
		if err := DB.Create(&user).Error; err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	tokenStr,err := middleware.SignWxToken(user.ID,180)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token":tokenStr,
		"error": "",
	})
}
