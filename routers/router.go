package routers

import (
	"gin-test/controllers"
	"gin-test/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

//路由
func InitRouter () *gin.Engine {
	r := gin.Default()
	//启用 zap log中间件
	r.Use(middlewares.GinLogger())
	//测试
	test := new(controllers.TestController)
	r.GET("ping", test.Ping)

	//路由分组 并使用自定义中间件
	member := r.Group("/api/member", middlewares.UseTime())
	memberController := new(controllers.MemberController)
	{
		member.GET("/list", memberController.List)
		member.GET("/add", memberController.Add)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code" : http.StatusNotFound,
			"msg" : "no route",
		})
	})

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code" : http.StatusNotFound,
			"msg" : "no method",
		})
	})

	return r
}
