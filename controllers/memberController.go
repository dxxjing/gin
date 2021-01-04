package controllers

import (
	"github.com/gin-gonic/gin"
)

type MemberController struct {
	BaseController
}

func (t *MemberController) List (c *gin.Context) {
	t.RenderJson(c, "list")
}

func (t *MemberController) Add (c *gin.Context) {
	t.RenderJson(c, "add")
}
