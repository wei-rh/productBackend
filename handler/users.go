package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"productBackend/model"
)

func GetAllOrder(ctx *gin.Context){
	uidVal, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(http.StatusOK,gin.H{
			"error": ok,
		})
	}
	uid := uidVal.(int)

	takeorder := []model.TakeOrder{}
	DB.Where("userid = ?", uid).Find(&takeorder)
	for i, _ := range takeorder {
		//关联查询
		DB.Model(&takeorder[i]).Related(&takeorder[i].TakeServer)
	}
	deliverorder := []model.DeliverOrder{}
	DB.Where("userid = ?", uid).Find(&deliverorder)
	for i, _ := range deliverorder {
		//关联查询
		DB.Model(&deliverorder[i]).
			Related(&deliverorder[i].DeliverServer)
	}
	buyorder := []model.BuyOrder{}
	DB.Where("userid = ?", uid).Find(&buyorder)
	for i, _ := range buyorder {
		//关联查询
		DB.Model(&buyorder[i]).
			Related(&buyorder[i].BuyServer)
	}
	ctx.JSON(http.StatusOK,gin.H{
		"takeorder": takeorder,
		"deliverorder": deliverorder,
		"buyorder": buyorder,
		"error":"",
	})

}
