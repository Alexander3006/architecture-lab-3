package router

import (
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
)

type Controller struct {
	Handler func(interfaces.IConnection)
}

func (c Controller) Execute(connection interfaces.IConnection) {
	c.Handler(connection)
}