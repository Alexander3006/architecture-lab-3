package transport

import (
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
	"github.com/google/wire"
)

func NewHttpServer(router interfaces.IRouter) *HttpServer {
	httpServer := &HttpServer{
		router: router,
	}
	return httpServer
}

var Providers = wire.NewSet(NewHttpServer)