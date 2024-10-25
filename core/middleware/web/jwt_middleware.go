package web

import (
	"github.com/gin-gonic/gin"
	"github.com/trancecho/mundo-be-template/core/auth"
	"github.com/trancecho/mundo-be-template/core/libx"
	"log"
	"net/http"
)

// JWTAuthMiddleware 是一个Gin中间件，用于验证JWT token
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		log.Println("token:", token)
		if token == "" {
			libx.Err(c, http.StatusUnauthorized, "需要提供鉴权token", nil)
			c.Abort()
			return
		}
		if token[:7] != "Bearer " {
			libx.Err(c, http.StatusUnauthorized, "无效的token格式", nil)
			c.Abort()
			return
		}
		token = token[7:]

		claims, err := auth.ParseToken(token)
		if err != nil {
			libx.Err(c, http.StatusUnauthorized, "无效的或过期的token", err)
			c.Abort()
			return
		}

		// 将用户信息存储在上下文中，以便后续处理使用
		c.Set("uid", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
