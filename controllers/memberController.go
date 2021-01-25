package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

type MemberController struct {
	BaseController
}

func (t *MemberController) List (c *gin.Context) {
	str := "list 111"
	fmt.Println(str)
	t.RenderJson(c, str)
}

func (t *MemberController) Add (c *gin.Context) {
	t.RenderJson(c, "add")
}

func (t *MemberController) Sleep (c *gin.Context) {
	d := c.DefaultQuery("duration", "100")
	fmt.Println("list params:", d)
	n, _ := strconv.ParseInt(d, 10, 64)
	time.Sleep(time.Duration(n) * time.Second)
	t.RenderJson(c, fmt.Sprintf("pid:%d\n", os.Getpid()))
}
