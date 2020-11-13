package serviceInterfaces

import "github.com/Alexander3006/architecture-lab-3/server/domain/entities"

type IMachineService interface {
	CreateMachine(string, int) error
	DeleteMachine(int) error
	GetMachine(int) (*entities.Machine, error)
	GetAllMachines() ([]*entities.Machine, error)
	AddComponent(int, int) error
	RemoveComponent(int, int) error
}