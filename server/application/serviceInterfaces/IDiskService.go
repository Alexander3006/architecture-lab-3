package serviceInterfaces

import "github.com/Alexander3006/architecture-lab-3/server/domain/entities"

type IDiskService interface {
	CreateDisk(int) error
	DeleteDisk(int) error
	GetAllDisks() ([]*entities.Disk, error)
}