package router

import "github.com/gin-gonic/gin"

type Entity struct {
}

func (r Entity) Router(g *gin.RouterGroup) {
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
