package dal

import (
	"database/sql"
	"github.com/Alexander3006/architecture-lab-3/server/domain/entities"
)

type MachineRepository struct {
	Db *sql.DB
}

func (mr *MachineRepository) Create(m entities.Machine) error {

	_, err := mr.Db.Exec(
		"INSERT INTO \"Machines\" (name, \"cpuCount\") VALUES ($1, $2)",
		m.Name,
		m.CpuCount,
		)
	return err
}

func (mr *MachineRepository) Delete(m entities.Machine) error {
	_,err := mr.Db.Exec(`DELETE FROM "Machines" WHERE id = $1)`, m.Id);
	return err
}

func (mr *MachineRepository) Read(m entities.Machine) (*entities.Machine, error) {
	rows, err := mr.Db.Query(`SELECT * FROM "MachinesInfo"() WHERE id = $1)`, m.Id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	res := &entities.Machine{};
	if err :=rows.Scan(m.Id, m.Name, m.CpuCount, m.TotalDiskSpace); err != nil {
		return nil, err
	}
	return res, nil
}

func (mr *MachineRepository) GetAll() ([]*entities.Machine, error) {
	rows, err := mr.Db.Query(`SELECT * FROM "MachinesInfo"()`)
	if err != nil {
		return nil, err
	}
	var res []*entities.Machine
	defer rows.Close()
	for rows.Next() {
		var machine *entities.Machine
		if err := rows.Scan(machine.Id, machine.Name, machine.CpuCount, machine.TotalDiskSpace); err != nil {
			return nil, err
		}
		res = append(res, machine)
	}
	if res != nil {
		res = make([]*entities.Machine, 0)
	}
	return res, nil
}