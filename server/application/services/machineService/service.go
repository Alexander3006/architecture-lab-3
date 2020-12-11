package machineService

import (
	"fmt"
	"github.com/Alexander3006/architecture-lab-3/server/domain/entities"
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
)

type MachineService struct {
	machineRepository interfaces.IMachineRepository
	diskRepository interfaces.IDiskRepository
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

func (ms MachineService) AddDisk(machineId, diskId int) error {
	disk, err := ms.diskRepository.Get(entities.Disk{Id: diskId})
	if err != nil {
		return err
	}
	if disk.MachineId != nil {
		return fmt.Errorf("the disk %d is occupied by another machine", disk.Id)
	}
	disk.MachineId = &machineId
	if err := ms.diskRepository.Update(*disk); err != nil {
		return err
	}
	return nil
}

func (ms MachineService) RemoveDisk(machineId, diskId int) error {
	disk, err := ms.diskRepository.Get(entities.Disk{Id: diskId})
	if err != nil {
	  return err
	}
	if disk.MachineId != nil && *disk.MachineId != machineId {
	  return fmt.Errorf("the machine %d has no disk %d", machineId, diskId)
	}
	disk.MachineId = nil
	if err := ms.diskRepository.Update(*disk); err != nil {
	  return err
	}
	return nil
 }

func (ms MachineService) GetAllMachines() ([]*entities.Machine, error) {
	return ms.machineRepository.GetAll()
}