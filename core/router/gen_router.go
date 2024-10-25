package router

import (
	"github.com/gin-gonic/gin"
	"github.com/trancecho/mundo-be-template/core/middleware/response"
	"github.com/trancecho/mundo-be-template/core/middleware/web"
	"github.com/trancecho/mundo-be-template/core/router/protected"
	"log"
)

func GenerateRouters(r *gin.Engine) *gin.Engine {

	baseGroup := r.Group("")

	baseGroup = NewGroupBuilder().
		SetName("base").
		SetFatherGroup(baseGroup).
		AddRoute("").
		AddMiddleware(response.ResponseMiddleware()).
		SetRoutes(Entity{}.Router).
		Build()
	{
		// 继承/base
		protectedGroup := NewGroupBuilder().
			SetName("protected").
			SetFatherGroup(baseGroup).
			AddRoute("/api").
			AddMiddleware(web.JWTAuthMiddleware()). // 使用 JWTAuthMiddleware 中间件
			SetRoutes(protected.Entity{}.Router).
			Build()
		log.Println(protectedGroup)

	}
	return r
}
