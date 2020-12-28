package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemberController struct {

}

func (t *MemberController) List (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "list",
	})
}

func (t *MemberController) Add (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "add",
	})
}
