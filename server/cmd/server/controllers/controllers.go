package controllers

import (
	"github.com/Alexander3006/architecture-lab-3/server/application/serviceInterfaces"
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
)

type Controllers struct {
	router interfaces.IRouter
	machineService serviceInterfaces.IMachineService
	diskServices serviceInterfaces.IDiskService
}

func (c Controllers) Register()  {
	r := c.router
	//machineHandlers
	r.RegisterRoute("POST", "/api/machine/create", r.CreateController(c.createMachine))

	//diskHandlers

}