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
	r.RegisterRoute("GET", "/api/machine", r.CreateController(c.getMachineById))
	r.RegisterRoute("GET", "/api/machine/all", r.CreateController(c.getAllMachines))
	r.RegisterRoute("POST", "/api/machine/create", r.CreateController(c.createMachine))
	r.RegisterRoute("POST", "/api/machine/addDisk", r.CreateController(c.addDisk))
	r.RegisterRoute("DELETE", "/api/machine/delete", r.CreateController(c.deleteMachine))
	r.RegisterRoute("DELETE", "/api/machine/removeDisk", r.CreateController(c.removeDisk))

	//diskHandlers
	r.RegisterRoute("GET", "/api/disk/all", r.CreateController(c.getAllDisks))
	r.RegisterRoute("POST", "/api/disk/create", r.CreateController(c.createDisk))
	r.RegisterRoute("DELETE", "/api/disk/delete", r.CreateController(c.deleteDisk))
}