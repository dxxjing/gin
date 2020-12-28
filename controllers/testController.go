package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestController struct {

}

func (t *TestController) Ping (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "pong",
	})
}


