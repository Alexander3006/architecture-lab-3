//+build wireinject

package main

import (
	"github.com/Alexander3006/architecture-lab-3/server/application/services/diskService"
	"github.com/Alexander3006/architecture-lab-3/server/application/services/machineService"
	"github.com/Alexander3006/architecture-lab-3/server/cmd/server/controllers"
	"github.com/Alexander3006/architecture-lab-3/server/infrastructure/dal"
	"github.com/Alexander3006/architecture-lab-3/server/infrastructure/router"
	"github.com/Alexander3006/architecture-lab-3/server/infrastructure/transport"
	"github.com/google/wire"
)

func ComposeServices() (*controllers.Controllers, error) {
	panic(wire.Build(
		NewDbConnection,
		router.Provider,
		dal.Providers,
		diskService.Provider,
		machineService.Provider,
		controllers.Providers,
		GetConfig,
	))
}

func ComposeHttpServer() (*transport.HttpServer, error) {
	panic(wire.Build(
		router.Provider,
		transport.Providers,
		GetConfig,
		GetPort,
	))
}
