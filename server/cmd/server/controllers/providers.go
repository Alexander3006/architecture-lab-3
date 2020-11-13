package controllers

import (
	"github.com/Alexander3006/architecture-lab-3/server/application/serviceInterfaces"
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
	"github.com/google/wire"
)

func NewControllers(
	r interfaces.IRouter,
	ms serviceInterfaces.IMachineService,
	ds serviceInterfaces.IDiskService,
	) *Controllers {
	controllers := &Controllers{
		router: r,
		machineService: ms,
		diskServices: ds,
	}
	return controllers
}

var Providers = wire.NewSet(NewControllers)