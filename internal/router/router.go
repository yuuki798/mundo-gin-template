package router

import (
	"github.com/gin-gonic/gin"
	"github.com/trancecho/mundo-be-template/internal/app/ping"
)

type Entity struct {
}

func (r Entity) Router(g *gin.RouterGroup) {
	g.GET("/ping", ping.Handler)
	g.GET("/mysql", ping.TestMysql)
	g.GET("/redis", ping.TestRedis)
}
