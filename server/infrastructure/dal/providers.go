package dal

import (
	"database/sql"
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
	"github.com/google/wire"
)

func NewMachineRepository(db *sql.DB) *MachineRepository {
	return &MachineRepository{Db: db}
}

func NewDiskRepository(db *sql.DB) *DiskRepository {
	return &DiskRepository{Db: db}
}


var Providers = wire.NewSet(
	wire.Bind(new(interfaces.IMachineRepository), new(*MachineRepository)),
	wire.Bind(new(interfaces.IDiskRepository), new(*DiskRepository)),
	NewMachineRepository,
	NewDiskRepository,
	)