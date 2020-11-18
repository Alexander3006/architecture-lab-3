package machineService

import (
	"github.com/Alexander3006/architecture-lab-3/server/application/serviceInterfaces"
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
	"github.com/google/wire"
)

func NewMachineService(mr interfaces.IMachineRepository, dr interfaces.IDiskRepository) *MachineService {
	return &MachineService{
		machineRepository: mr,
		diskRepository: dr,
	}
}

var Provider = wire.NewSet(
	wire.Bind(new(serviceInterfaces.IMachineService), new(*MachineService)),
	NewMachineService,
	)