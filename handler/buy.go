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
		return
	}
	uid := uidVal.(int)
	//接收参数
	data := &model.BuyOrderForm{}
	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"error": err.Error(),
		})
		return
	}
	//
	buyserver := model.BuyServer{}
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
	buyorder.BuyServer = buyserver
	if err := DB.Create(&buyorder).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	//更新 buyserver 表
	buyserver.Buyid = int(buyorder.ID)
	if error := DB.Save(&buyserver).Error;error!=nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": error.Error(),
		})
		return
	}



	ctx.JSON(http.StatusOK,gin.H{
		"buyorder": buyorder,
		"buyserverbyuser": buyserver,
		"error":"",
	})
}

//查询订单,通过id查询一个订单
func FindBuyOrder(ctx *gin.Context)  {
	id := ctx.Query("id")
	if id=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"id is null",
		})
		return
	}
	buyorder := []model.BuyOrder{}
	DB.Where("id = ?", id).Find(&buyorder)
	for i, _ := range buyorder {
		//关联查询
		DB.Model(&buyorder[i]).
			Related(&buyorder[i].BuyServer)
	}
	ctx.JSON(http.StatusOK,gin.H{
		"order": buyorder,
		"error":"",
	})
}
