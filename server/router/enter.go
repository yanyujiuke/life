package router

import (
	"life/router/example"
	"life/router/overhead"
	"life/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Overhead overhead.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
