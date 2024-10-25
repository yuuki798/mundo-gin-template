package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

type GroupBuilder struct {
	name        string
	middlewares []gin.HandlerFunc
	path        string
	group       *gin.RouterGroup
	routes      func(*gin.RouterGroup)
}

func NewGroupBuilder() *GroupBuilder {
	return &GroupBuilder{}
}

func (b *GroupBuilder) SetFatherGroup(group *gin.RouterGroup) *GroupBuilder {
	b.group = group
	b.path = group.BasePath()
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
	b.group.Use(middleware...)
	return b
}

func (b *GroupBuilder) AddRoute(path string) *GroupBuilder {
	b.path = b.path + path
	return b
}

func (b *GroupBuilder) Build() *gin.RouterGroup {
	if b.group == nil {
		log.Panic("fatherGroup is nil")
	}
	if b.routes == nil {
		log.Panic("routes function is nil")
	}
	group := b.group.Group(b.path)
	for _, middleware := range b.middlewares {
		group.Use(middleware)
	}
	b.routes(group)
	return group
}
