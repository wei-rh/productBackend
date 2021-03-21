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
		return
	}
	uid := uidVal.(int)
	//接收参数
	data := &model.TakeOrderForm{}
	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"error": err.Error(),
		})
		return
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
	deliverorder.DeliverServer=*deliverserver
	if err := DB.Create(&deliverorder).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	//更新 deliverserver 表
	deliverserver.Deliverid = int(deliverorder.ID)
	if error := DB.Save(&deliverserver).Error;error!=nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"deliverorder": deliverorder,
		"deliverserverbyuser": deliverserver,
		"error":"",
	})
}

//查询订单,通过id查询一个订单
func FindDeliverOrder(ctx *gin.Context)  {
	id := ctx.Query("id")
	if id=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"id is null",
		})
		return
	}
	deliverorder := []model.DeliverOrder{}
	DB.Where("id = ?", id).Find(&deliverorder)
	for i, _ := range deliverorder {
		//关联查询
		DB.Model(&deliverorder[i]).
			Related(&deliverorder[i].DeliverServer)
	}
	ctx.JSON(http.StatusOK,gin.H{
		"order": deliverorder,
		"error":"",
	})
}
