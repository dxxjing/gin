package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {

}

func (b *BaseController) RenderJson(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code" : 0,
		"msg" : "success",
		"data" : data,
	})
}

func (b *BaseController) RenderError(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : msg,
		"data" : "",
	})
}
