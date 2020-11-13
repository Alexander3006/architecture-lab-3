package diskService

import (
	"github.com/Alexander3006/architecture-lab-3/server/application/serviceInterfaces"
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
	"github.com/google/wire"
)

func NewDiskService(dr interfaces.IDiskRepository) *DiskService {
	return &DiskService{
		diskRepository: dr,
	}
}

var Provider = wire.NewSet(
	wire.Bind(new(serviceInterfaces.IDiskService), new(*DiskService)),
	NewDiskService,
	)
