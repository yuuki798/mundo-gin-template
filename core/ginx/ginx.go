package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/trancecho/mundo-be-template/core/middleware/cors"
	"github.com/trancecho/mundo-be-template/core/router"
)

func GinInit() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Middleware())
	router.GenerateRouters(r)
	return r
}
