package protected

import (
	"github.com/gin-gonic/gin"
	"github.com/trancecho/mundo-be-template/core/middleware/response"
	"github.com/trancecho/mundo-be-template/core/middleware/web"
)

func Router(r *gin.Engine) {
	g := r.Group("/api")
	// 使用 ResponseMiddleware 中间件
	g.Use(response.ResponseMiddleware())
	g.Use(web.JWTAuthMiddleware())
	{
		g.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

}
