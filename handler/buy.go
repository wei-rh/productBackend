package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"productBackend/model"
)

//添加订单
func BuyTakeOrder(ctx *gin.Context){
	uidVal, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(http.StatusOK,gin.H{
			"error": ok,
		})
	}
	uid := uidVal.(int)
	//接收参数
	data := &model.BuyOrderForm{}
	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"error": err.Error(),
		})
	}
	//
	buyserver := &model.BuyServer{}
	buyserver.Userid=uid
	buyserver.Status=0
	if err := DB.Create(&buyserver).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	//存数据库
	buyorder := &model.BuyOrder{}
	buyorder.Userid = uid
	buyorder.SendAddress=data.SendAddress
	buyorder.GetAddress=data.GetAddress
	buyorder.Baojia=data.Baojia
	buyorder.Context=data.Context
	buyorder.Remarks=data.Remarks
	buyorder.Time=data.Time
	buyorder.BuyServerID= int(buyserver.ID)
	if err := DB.Create(&buyorder).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"buyorder": buyorder,
		"buyserver": buyserver,
		"error":"",
	})
}

