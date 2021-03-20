package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"productBackend/model"
	"strconv"
)

//获取用户所有类型的订单
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

//订单支付成功接口
func SetStatusOne(ctx *gin.Context)  {
	id := ctx.Query("id")
	if id=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"id is null",
		})
		return
	}
	ID,_ := strconv.Atoi(id)
	ordertype := ctx.Query("type")
	if ordertype=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"type is null",
		})
		return
	}

	switch ordertype {
	case "take":
		takeServerByUser := &model.TakeServer{}
		takeServerByUser.ID = uint(ID)
		if error := DB.Model(&takeServerByUser).Update("status",1).Error;error!=nil{
			ctx.JSON(http.StatusOK,gin.H{
				"error":error.Error(),
			})
			return
		}
	case "buy":
		buyServerByUser := &model.BuyServer{}
		buyServerByUser.ID = uint(ID)
		if error := DB.Model(&buyServerByUser).Update("status",1).Error;error!=nil{
			ctx.JSON(http.StatusOK,gin.H{
				"error":error.Error(),
			})
			return
		}
	case "deliver":
		deliverServerByUser := &model.DeliverServer{}
		deliverServerByUser.ID = uint(ID)
		if error := DB.Model(&deliverServerByUser).Update("status",1).Error;error!=nil{
			ctx.JSON(http.StatusOK,gin.H{
				"error":error.Error(),
			})
			return
		}
	default:

	}

	ctx.JSON(http.StatusOK,gin.H{
		"error":"",
	})
}

//取消订单
func CancelOrder(ctx *gin.Context){
	id := ctx.Query("id")
	if id=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"id is null",
		})
		return
	}
	ID,_ := strconv.Atoi(id)
	ordertype := ctx.Query("type")
	if ordertype=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"type is null",
		})
		return
	}


	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(id)
	fmt.Println(ordertype)
	switch ordertype {
	case "take":
		fmt.Println("take")

		takeServerByUser := &model.TakeServer{}
		takeServerByUser.ID = uint(ID)
		if error := DB.Model(&takeServerByUser).Update("status",4).Error;error!=nil{
			ctx.JSON(http.StatusOK,gin.H{
				"error":error.Error(),
			})
			return
		}
	case "buy":
		fmt.Println("buy")

		buyServerByUser := &model.BuyServer{}
		buyServerByUser.ID = uint(ID)
		if error := DB.Model(&buyServerByUser).Update("status",4).Error;error!=nil{
			ctx.JSON(http.StatusOK,gin.H{
				"error":error.Error(),
			})
			return
		}
	case "deliver":
		fmt.Println("deliver")

		deliverServerByUser := &model.DeliverServer{}
		deliverServerByUser.ID = uint(ID)
		if error := DB.Model(&deliverServerByUser).Update("status",4).Error;error!=nil{
			ctx.JSON(http.StatusOK,gin.H{
				"error":error.Error(),
			})
			return
		}
	default:

	}
	ctx.JSON(http.StatusOK,gin.H{
		"error":"成功",
	})
}

//删除订单
func DeleteOrder(ctx *gin.Context) {
	id := ctx.Query("id")
	if id=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"id is null",
		})
	}
	ID,_ := strconv.Atoi(id)
	ordertype := ctx.Query("type")
	if ordertype=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"type is null",
		})
	}

	switch ordertype {
	case "take":
		takeorder := &model.TakeOrder{}
		takeorder.ID = uint(ID)

		if error := DB.Delete(takeorder).Error;error!=nil{
			ctx.JSON(http.StatusOK,gin.H{
				"error":error.Error(),
			})
			return
		}
	case "buy":
		buyorder := &model.BuyOrder{}
		buyorder.ID = uint(ID)


		if error := DB.Delete(buyorder).Error;error!=nil{
			ctx.JSON(http.StatusOK,gin.H{
				"error":error.Error(),
			})
			return
		}
	case "deliver":
		deliverorder := &model.DeliverOrder{}
		deliverorder.ID = uint(ID)
		if error := DB.Delete(deliverorder).Error;error!=nil{
			ctx.JSON(http.StatusOK,gin.H{
				"error":error.Error(),
			})
			return
		}
	default:

	}
	ctx.JSON(http.StatusOK,gin.H{
		"error":"成功",
	})

}

