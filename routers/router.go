package routers

import (
	"gin-test/controllers"
	"gin-test/middlewares"
	"github.com/gin-gonic/gin"
)

//路由
func InitRouter () *gin.Engine {
	r := gin.Default()
	//测试
	test := new(controllers.TestController)
	r.GET("ping", test.Ping)

	//路由分组 并使用自定义中间件
	member := r.Group("/api/member", middlewares.UseTime())
	memberController := new(controllers.MemberController)
	{
		member.GET("/list", memberController.List)
		member.POST("/add", memberController.Add)
	}


	return r
}
