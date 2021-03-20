package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"productBackend/model"
)

//查找服务人员信息，通过id查询server表中的数据
func FindServer(ctx *gin.Context){
	id := ctx.Query("id")
	if id=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"id is null",
		})
		return
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

//查找服务人员信息，通过uid查询server表中的数据,通过用户查看有没有跑腿权限
func ByFindServer(ctx *gin.Context){
	uidVal, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(http.StatusOK,gin.H{
			"error": ok,
		})
		return
	}
	uid := uidVal.(int)
	server := &model.Server{}
	if err := DB.Where("userid = ?", uid).Find(&server).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"server":"",
			"error":"",
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"server": server,
		"error":"",
	})
}


//骑手注册
func Register(ctx *gin.Context)  {
	uidVal, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(http.StatusOK,gin.H{
			"error": ok,
		})
		return
	}
	uid := uidVal.(int)

	name := ctx.Query("name")
	if name=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"name is null",
		})
		return
	}
	phone := ctx.Query("phone")
	if phone=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"phone is null",
		})
		return
	}
	number := ctx.Query("number")
	if number=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"number is null",
		})
		return
	}
	if number != "1234"{
		ctx.JSON(http.StatusOK,gin.H{
			"error":"number is error",
		})
		return
	}


	server := &model.Server{}
	server.Userid = uid
	server.Name= name
	server.Tel = phone

	if error := DB.Create(&server).Error;error!=nil{
		ctx.JSON(http.StatusOK,gin.H{
			"error":error,
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"error":"",
		"server":server,
	})


}

//骑手获取所有可接订单
func GetAllStatusOne(ctx *gin.Context){

	take := []model.TakeOrder{}
	t := AllTakeOrder()
	for i,_ := range t{
		if t[i].TakeServer.Status==1 {
			take = append(take, t[i])
		}
	}
	buy := []model.BuyOrder{}
	b := AllBuyOrder()
	for i,_ := range b{
		if b[i].BuyServer.Status==1 {
			buy = append(buy, b[i])
		}
	}
	deliver := []model.DeliverOrder{}
	d := AllDeliverOrder()
	for i,_ := range d{
		if d[i].DeliverServer.Status==1 {
			deliver = append(deliver, d[i])
		}
	}
	fmt.Println(t[0].TakeServer)
	ctx.JSON(http.StatusOK,gin.H{
		"take":take,
		"buy":buy,
		"deliver":deliver,
	})
}



//骑手获取已接订单，派送中。。。
func GetServerStatusTwo()  {

}

//骑手获取已接订单，已完成。。。
func GetServerStatusThree()  {

}
