package transport

import (
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
	"github.com/google/wire"
)

func NewHttpServer(router interfaces.IRouter, port interfaces.HttpPortNumber) *HttpServer {
	httpServer := &HttpServer{
		router: router,
		port:   port,
	}
	return httpServer
}

var Providers = wire.NewSet(NewHttpServer)
