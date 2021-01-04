package controllers

import (
	"github.com/gin-gonic/gin"
)

type TestController struct {
	BaseController
}

func (t *TestController) Ping (c *gin.Context) {
	t.RenderJson(c,"pong")
}


