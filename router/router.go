package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"productBackend/handler"
	"productBackend/middleware"
)

func RouterInit()  *gin.Engine {
	r := gin.Default()
	//默认的cors 策略，解决跨域问题
	r.Use(cors.Default())
	// 只允许 localhost 与的跨域请求
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://localhost"}
	//config.AllowHeaders = []string{
	//	"authorization",
	//}
	//r.Use(cors.New(config))

	//路由
	//r.POST("/user-login", handler.UserLogin)


	//用户登录接口
	r.GET("/login",handler.AppletWeChatLogin)
	authGroup := r.Group("/", middleware.AuthJWT)
	{
		//用户下单-取订单
		authGroup.POST("/takeorder",handler.GetTakeOrder)
		//用户下单-送订单
		authGroup.POST("/deliverorder",handler.GetDeliverOrder)
		//用户下单-买订单
		authGroup.POST("/buyorder",handler.BuyTakeOrder)
		//获取用户所有订单
		authGroup.GET("/allorder",handler.GetAllOrder)
	}
	//获取用户-取 订单-的某一个
	r.GET("/onetakeorder",handler.FindTakeOrder)
	//获取用户-买订单-的某一个
	r.GET("/onebuyorder",handler.FindBuyOrder)
	//获取用户-送订单-的某一个
	r.GET("/onedeliverorder",handler.FindDeliverOrder)
	//获取一个server的信息
	r.GET("/getserver",handler.FindServer)
	return r
}
