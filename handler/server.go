package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"productBackend/model"
)

//查找服务人员信息
func FindServer(ctx *gin.Context){
	id := ctx.Query("id")
	if id=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"id is null",
		})
	}

	server := &model.Server{}
	if err := DB.Where("id = ?", id).Find(&server).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"server": server,
		"error":"",
	})
}

