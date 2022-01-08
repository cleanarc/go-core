package rest

import "github.com/gin-gonic/gin"

type Router struct {
	*gin.Engine
	groups map[string]*gin.RouterGroup
}

func NewRouter() *Router {
	engine := gin.Default()

	return &Router{
		Engine: engine,
		groups: map[string]*gin.RouterGroup{},
	}
}

func (r *Router) AddGroup(path string) {
	r.groups[path] = r.Engine.Group(path)
}

func (r *Router) Group(path string) *gin.RouterGroup {
	return r.groups[path]
}
