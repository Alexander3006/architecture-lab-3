package interfaces

import "github.com/Alexander3006/architecture-lab-3/server/domain/entities"

type IMachineRepository interface {
	Create(entities.Machine) error
	Read(entities.Machine) (*entities.Machine, error)
	Delete(entities.Machine) error
	GetAll() ([]*entities.Machine, error)
}

type IComponentRepository interface {
	Create(entities.Component) error
	Delete(entities.Component) error
}

type IDiskRepository interface {
	Create(entities.Disk) error
	Delete(entities.Disk) error
	GetAll() ([]*entities.Disk, error)
}