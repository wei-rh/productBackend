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
	//关联查询
	DB.Model(&takeorder).Related(&takeorder.TakeServer)


	ctx.JSON(http.StatusOK,gin.H{
		"takeorder": takeorder,
		"takeserverbyuser": takeserver,
		"error":"",
	})


}

//查询订单,通过id查询一个订单
func FindTakeOrder(ctx *gin.Context)  {
	 id := ctx.Query("id")
	 if id=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"id is null",
		})
		 return
	}
	takeorder := []model.TakeOrder{}
	DB.Where("id = ?", id).Find(&takeorder)
	for i, _ := range takeorder {
		//关联查询
		DB.Model(&takeorder[i]).Related(&takeorder[i].TakeServer)
	}
	ctx.JSON(http.StatusOK,gin.H{
		"order": takeorder,
		"error":"",
	})
}

//获取所有订单

func AllTakeOrder() []model.TakeOrder {
	takeorder := []model.TakeOrder{}
	DB.Find(&takeorder)
	for i, _ := range takeorder {
		//关联查询
		DB.Model(&takeorder[i]).Related(&takeorder[i].TakeServer)
	}
	return takeorder
}


