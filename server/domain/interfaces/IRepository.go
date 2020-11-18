package interfaces

import "github.com/Alexander3006/architecture-lab-3/server/domain/entities"

type IMachineRepository interface {
	Create(entities.Machine) error
	Read(entities.Machine) (*entities.Machine, error)
	Delete(entities.Machine) error
	GetAll() ([]*entities.Machine, error)
}

type IDiskRepository interface {
	Update(entities.Disk) error
	Create(entities.Disk) error
	Delete(entities.Disk) error
	GetAll() ([]*entities.Disk, error)
	Get(entities.Disk) (*entities.Disk, error)
}