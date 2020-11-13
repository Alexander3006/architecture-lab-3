package diskService

import (
	"fmt"
	"github.com/Alexander3006/architecture-lab-3/server/domain/entities"
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
)

type DiskService struct {
	diskRepository interfaces.IDiskRepository
}

func (ds DiskService) CreateDisk(space int) error {
	if space < 0 {
		return fmt.Errorf("disk space cannot be negative")
	}
	disk := entities.Disk{Space: space}
	err := ds.diskRepository.Create(disk)
	return err
}

func (ds DiskService) DeleteDisk(id int) error {
	disk := entities.Disk{Id: id}
	err := ds.diskRepository.Delete(disk)
	return err
}

func (ds DiskService) GetAllDisks() ([]*entities.Disk, error) {
	return ds.diskRepository.GetAll()
}