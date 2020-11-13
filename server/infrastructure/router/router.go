package router

import (
	"fmt"
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
	"strings"
)

type RouteStorage map[string]map[string]interfaces.IController

type Router struct {
	routes RouteStorage
}

func (r *Router) RegisterRoute(method string, route string, controller interfaces.IController) error {
	storage := &r.routes
	methodUp := strings.ToUpper(method);
	if _, ok := (*storage)[methodUp]; !ok {
		(*storage)[methodUp] = make(map[string]interfaces.IController)
	}
	if _, ok := (*storage)[methodUp][route]; ok {
		return fmt.Errorf("controller already exists at the specified path")
	}
	(*storage)[methodUp][route] = controller
	return nil
}

func (r Router) IndicateRoute(connection interfaces.IConnection) {
	storage := r.routes
	request := connection.GetRequest()
	method := strings.ToUpper(request.Method)
	path := request.URL.Path
	if methodStorage, ok := storage[method]; ok {
		if controller, ok := methodStorage[path]; ok {
			controller.Execute(connection)
			return
		}
	}
	connection.NotFound("Not Found")
	return
}

func (r Router) CreateController(handler func(interfaces.IConnection)) interfaces.IController {
	return &Controller{
		Handler: handler,
	}
}