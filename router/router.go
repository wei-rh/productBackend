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
	r.GET("/login",handler.AppletWeChatLogin)
	authGroup := r.Group("/", middleware.AuthJWT)
	{
		authGroup.POST("/takeorder",handler.GetTakeOrder)
		authGroup.POST("/deliverorder",handler.GetDeliverOrder)
		authGroup.POST("/buyorder",handler.BuyTakeOrder)
		authGroup.GET("/allorder",handler.GetAllOrder)
	}
	return r
}
