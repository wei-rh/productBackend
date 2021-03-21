package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"productBackend/model"
	"strconv"
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
	takeorder := model.TakeOrder{}
	takeserver := []model.TakeServer{}
	//获取订单状态为1的take订单
	if err := DB.Where("status = ?",1).Find(&takeserver).Error;err!=nil{
		ctx.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}
	//获取take订单
	for _,v:=range takeserver{
		//DB.Where("id = ?",v.Takeid).Find(&takeorder)

		takeorder.ID = uint(v.Takeid)
		DB.Find(&takeorder)
		takeorder.TakeServer = v
		take = append(take,takeorder)
	}

	buy := []model.BuyOrder{}
	buyorder := model.BuyOrder{}
	buyserver := []model.BuyServer{}
	DB.Where("status=?",1).Find(&buyserver)
	for _,v := range buyserver{
		buyorder.ID = uint(v.Buyid)
		DB.Find(&buyorder)
		buyorder.BuyServer = v
		buy = append(buy,buyorder)
	}

	deliver := []model.DeliverOrder{}
	deliverorder := model.DeliverOrder{}
	deliverserver := []model.DeliverServer{}
	DB.Where("status=?",1).Find(&deliverserver)
	for _,v := range deliverserver {
		deliverorder.ID = uint(v.Deliverid)
		DB.Find(&deliverorder)
		deliverorder.DeliverServer = v
		deliver = append(deliver,deliverorder)
	}

	ctx.JSON(http.StatusOK,gin.H{
		"take":take,
		"buy":buy,
		"deliver":deliver,
		"error":"",
	})
}



//骑手获取已接订单，派送中。。。
func GetServerStatusTwo(ctx *gin.Context)  {
	uidVal, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(http.StatusOK,gin.H{
			"error": ok,
		})
		return
	}
	uid := uidVal.(int)

	server := model.Server{}

	if error := DB.Where("userid=?",uid).Find(&server).Error;error!=nil{
		ctx.JSON(http.StatusOK,gin.H{
			"error": error.Error(),
		})
		return
	}

	take := []model.TakeOrder{}
	takeorder := model.TakeOrder{}
	takeserver := []model.TakeServer{}
	//获取订单状态为1的take订单
	if err := DB.Where("status = ? AND serverid=?",2,server.ID).Find(&takeserver).Error;err!=nil{
		ctx.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}
	//获取take订单
	for _,v:=range takeserver{
		takeorder.ID = uint(v.Takeid)
		DB.Find(&takeorder)
		takeorder.TakeServer = v
		take = append(take,takeorder)
	}

	buy := []model.BuyOrder{}
	buyorder := model.BuyOrder{}
	buyserver := []model.BuyServer{}
	DB.Where("status = ? AND serverid=?",2,server.ID).Find(&buyserver)
	for _,v := range buyserver{
		buyorder.ID = uint(v.Buyid)
		DB.Find(&buyorder)
		buyorder.BuyServer = v
		buy = append(buy,buyorder)
	}

	deliver := []model.DeliverOrder{}
	deliverorder := model.DeliverOrder{}
	deliverserver := []model.DeliverServer{}
	DB.Where("status = ? AND serverid=?",2,server.ID).Find(&deliverserver)
	for _,v := range deliverserver {
		deliverorder.ID = uint(v.Deliverid)
		DB.Find(&deliverorder)
		deliverorder.DeliverServer = v
		deliver = append(deliver,deliverorder)
	}
	ctx.JSON(http.StatusOK,gin.H{
		"take":take,
		"buy":buy,
		"deliver":deliver,
		"error":"",
	})






}

//骑手获取已接订单，已完成。。。
func GetServerStatusThree(ctx *gin.Context)  {
	uidVal, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(http.StatusOK,gin.H{
			"error": ok,
		})
		return
	}
	uid := uidVal.(int)

	server := model.Server{}

	if error := DB.Where("userid=?",uid).Find(&server).Error;error!=nil{
		ctx.JSON(http.StatusOK,gin.H{
			"error": error.Error(),
		})
		return
	}

	take := []model.TakeOrder{}
	takeorder := model.TakeOrder{}
	takeserver := []model.TakeServer{}
	//获取订单状态为1的take订单
	if err := DB.Where("status = ? AND serverid=?",3,server.ID).Find(&takeserver).Error;err!=nil{
		ctx.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
	}
	//获取take订单
	for _,v:=range takeserver{
		takeorder.ID = uint(v.Takeid)
		DB.Find(&takeorder)
		takeorder.TakeServer = v
		take = append(take,takeorder)
	}

	buy := []model.BuyOrder{}
	buyorder := model.BuyOrder{}
	buyserver := []model.BuyServer{}
	DB.Where("status = ? AND serverid=?",3,server.ID).Find(&buyserver)
	for _,v := range buyserver{
		buyorder.ID = uint(v.Buyid)
		DB.Find(&buyorder)
		buyorder.BuyServer = v
		buy = append(buy,buyorder)
	}

	deliver := []model.DeliverOrder{}
	deliverorder := model.DeliverOrder{}
	deliverserver := []model.DeliverServer{}
	DB.Where("status = ? AND serverid=?",3,server.ID).Find(&deliverserver)
	for _,v := range deliverserver {
		deliverorder.ID = uint(v.Deliverid)
		DB.Find(&deliverorder)
		deliverorder.DeliverServer = v
		deliver = append(deliver,deliverorder)
	}
	ctx.JSON(http.StatusOK,gin.H{
		"take":take,
		"buy":buy,
		"deliver":deliver,
		"error":"",
	})
}

//骑手接单或者点击确定
func OrderReceiving(ctx *gin.Context){
	uidVal, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(http.StatusOK,gin.H{
			"error": ok,
		})
		return
	}
	uid := uidVal.(int)

	ordertype := ctx.Query("type")
	if ordertype=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"ordertype is null",
		})
		return
	}
	id := ctx.Query("id")
	if id=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"id is null",
		})
		return
	}
	orderid,_ := strconv.Atoi(id)
	//类型，接单或者确认订单  接单 0， 确认订单1
	current2 := ctx.Query("current2")
	if current2=="" {
		ctx.JSON(http.StatusOK,gin.H{
			"error":"current2 is null",
		})
		return
	}
	if current2 =="2" {
		return
	}
	//获取server信息，拿到serverid
	server := model.Server{}
	DB.Where("userid=?",uid).Find(&server)


	var cc interface{}

	switch ordertype {
	case "take":
		takeserver := model.TakeServer{}
		DB.Where("takeid=?",orderid).Find(&takeserver)
		if current2=="0" {
			//接单
			takeserver.Serverid = int(server.ID)
			takeserver.Status = 2
		}else {
			//确认
			takeserver.Status = 3
		}
		DB.Save(&takeserver)
		cc = takeserver
		break
	case "buy":
		buyserver := model.BuyServer{}
		DB.Where("buyid=?",orderid).Find(&buyserver)
		if current2=="0" {
			//接单
			buyserver.Serverid = int(server.ID)
			buyserver.Status = 2
		}else {
			buyserver.Status = 3
		}
		DB.Save(&buyserver)
		cc = buyserver

		break
	case "deliver":
		deliverserver := model.DeliverServer{}
		DB.Where("deliverid=?",orderid).Find(&deliverserver)
		if current2=="0" {
			//接单
			deliverserver.Serverid = int(server.ID)
			deliverserver.Status = 2
		}else {
			deliverserver.Status = 3
		}
		DB.Save(&deliverserver)
		cc=deliverserver
		break
	}
	ctx.JSON(http.StatusOK,gin.H{
		"error":"",
		"cc":cc,

	})

}