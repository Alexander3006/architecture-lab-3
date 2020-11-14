package machineService

import (
	"fmt"
	"github.com/Alexander3006/architecture-lab-3/server/domain/entities"
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
)

type MachineService struct {
	machineRepository interfaces.IMachineRepository
	componentRepository interfaces.IComponentRepository
}

func (ms MachineService) CreateMachine(name string, cpuCount int) error {
	if len(name) < 0 {
		return fmt.Errorf("machine name is not provided")
	}
	if cpuCount < 0 {
		return fmt.Errorf("count cpu cannot be negative")
	}
	machine := entities.Machine{
		Name: name,
		CpuCount: cpuCount,
	}
	err := ms.machineRepository.Create(machine)
	return err
}

func (ms MachineService) DeleteMachine(id int) error {
	machine := entities.Machine{
		Id: id,
	}
	err := ms.machineRepository.Delete(machine);
	return err
}

func (ms MachineService) GetMachine(id int) (*entities.Machine, error) {
	machine := entities.Machine{
		Id: id,
	}
	return ms.machineRepository.Read(machine)
}

func (ms MachineService) AddComponent(machineId, diskId int) error {
	component := entities.Component{
		MachineId: machineId,
		DiskId: diskId,
	}
	err := ms.componentRepository.Create(component)
	return err
}

func (ms MachineService) RemoveComponent(machineId, diskId int) error {
	component := entities.Component{
		MachineId: machineId,
		DiskId: diskId,
	}
	err := ms.componentRepository.Delete(component)
	return err
}

func (ms MachineService) GetAllMachines() ([]*entities.Machine, error) {
	return ms.machineRepository.GetAll()
}