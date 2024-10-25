package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

// MyGroup 方便扩展和打印group
type MyGroup struct {
	name        string
	fatherGroup *gin.RouterGroup
	middlewares []gin.HandlerFunc

	g *gin.RouterGroup
}

type GroupBuilder struct {
	name        string
	middlewares []gin.HandlerFunc
	path        string
	group       *MyGroup
	routes      func(*gin.RouterGroup)
}

func NewGroupBuilder() *GroupBuilder {
	return &GroupBuilder{}
}

func (b *GroupBuilder) SetFatherGroup(entity *MyGroup) *GroupBuilder {
	b.group = entity
	b.path = entity.g.BasePath()
	return b
}

func (b *GroupBuilder) SetName(name string) *GroupBuilder {
	b.name = name
	return b
}

func (b *GroupBuilder) SetRoutes(routerFunc func(*gin.RouterGroup)) *GroupBuilder {
	b.routes = routerFunc
	return b
}

type Routes interface {
	Router(*gin.RouterGroup)
}

func (b *GroupBuilder) AddMiddleware(middleware ...gin.HandlerFunc) *GroupBuilder {
	if b.group == nil || b.group.g == nil {
	}
	b.group.g.Use(middleware...)
	return b
}

func (b *GroupBuilder) AddRoute(path string) *GroupBuilder {
	b.path = b.path + path
	return b
}

func (b *GroupBuilder) Build() *MyGroup {
	if b.group == nil {
		log.Panic("fatherGroup is nil")
	}
	if b.routes == nil {
		log.Panic("routes function is nil")
	}

	group := b.group.g.Group(b.path)
	for _, middleware := range b.middlewares {
		group.Use(middleware)
	}
	b.routes(group)
	return &MyGroup{
		name:        b.name,
		fatherGroup: b.group.g,
		middlewares: b.middlewares,
		g:           group,
	}
}
