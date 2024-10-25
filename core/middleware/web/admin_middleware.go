package web

import (
	"github.com/gin-gonic/gin"
	"github.com/trancecho/mundo-be-template/core/libx"
)

// AdminMiddleware 这里代表只有admin才能访问的中间件
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if libx.GetRole(c) != "admin" {
			c.Abort()
			libx.Err(c, 401, "需要管理员权限", nil)
			return
		}
		c.Next()
	}
}
