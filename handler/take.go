package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"productBackend/model"
)

//添加订单
func GetTakeOrder(ctx *gin.Context){
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
	takeserver := &model.TakeServer{}
	takeserver.Userid=uid
	takeserver.Status=0
	if err := DB.Create(&takeserver).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	//存数据库
	takeorder := &model.TakeOrder{}
	takeorder.Userid = uid
	takeorder.SendAddress=data.SendAddress
	takeorder.GetAddress=data.GetAddress
	takeorder.Baojia=data.Baojia
	takeorder.Goods=data.Goods
	takeorder.Remarks=data.Remarks
	takeorder.Time=data.Time
	takeorder.Weight=data.Weight
	takeorder.TakeServerID = int(takeserver.ID)

	if err := DB.Create(&takeorder).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	//

	ctx.JSON(http.StatusOK,gin.H{
		"takeorder": takeorder,
		"takeserver": takeserver,
		"error":"",
	})

}
