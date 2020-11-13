package dal

import (
	"database/sql"
	"github.com/Alexander3006/architecture-lab-3/server/domain/entities"
)

type ComponentRepository struct {
	Db *sql.DB
}

func (cp *ComponentRepository) Create(c entities.Component) error {
	_, err := cp.Db.Exec(
		`INSERT INTO "MachineComponents"("machineId","diskId") VALUES ($1, $2)`,
		c.MachineId,
		c.DiskId,
	)
	return err
}

func (cp *ComponentRepository) Delete(c entities.Component) error {
	_, err := cp.Db.Exec(
		`DELETE  from "MachineComponents" WHERE "machineId" = $1 AND "diskId = $2"`,
		c.MachineId,
		c.DiskId,
	)
	return err
}

