package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"productBackend/model"
)

//添加订单
func GetDeliverOrder(ctx *gin.Context){
	uidVal, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(http.StatusOK,gin.H{
			"error": ok,
		})
	}
	uid := uidVal.(int)
	//接收参数
	data := &model.TakeOrderForm{}
	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"error": err.Error(),
		})
	}
	//
	deliverserver := &model.DeliverServer{}
	deliverserver.Userid=uid
	deliverserver.Status=0
	if err := DB.Create(&deliverserver).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	//存数据库
	deliverorder := &model.DeliverOrder{}
	deliverorder.Userid = uid
	deliverorder.SendAddress=data.SendAddress
	deliverorder.GetAddress=data.GetAddress
	deliverorder.Baojia=data.Baojia
	deliverorder.Goods=data.Goods
	deliverorder.Remarks=data.Remarks
	deliverorder.Time=data.Time
	deliverorder.Weight=data.Weight
	deliverorder.DeliverServerID= int(deliverserver.ID)
	if err := DB.Create(&deliverorder).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"deliverorder": deliverorder,
		"deliverserver": deliverserver,
		"error":"",
	})
}
