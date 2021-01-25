package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		zap.L().WithOptions(zap.Fields(zap.Int("uid", 3688)))
		//虽然能实现动态增加 固定字段 但是要每次请求都要new
		//zap.ReplaceGlobals(zap.New(zap.L().Core()).WithOptions(zap.Fields(zap.Int("uid", 3688))))
		//zap.L().With()
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Namespace("data"),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
