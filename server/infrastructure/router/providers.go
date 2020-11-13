package router

import (
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
	"github.com/google/wire"
)

var routerInstance *Router

func NewRouter() *Router {
	if routerInstance == nil {
		routerInstance = &Router{
			routes:map[string]map[string]interfaces.IController{},
		}
	}
	return routerInstance
}

var Provider = wire.NewSet(
	NewRouter,
	wire.Bind(new(interfaces.IRouter), new(*Router)),
	)