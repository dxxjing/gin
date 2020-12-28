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
	return
}

func (t *TestController) List (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "list",
	})
	return
}

func (t *TestController) Add (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "add",
	})
	return
}


