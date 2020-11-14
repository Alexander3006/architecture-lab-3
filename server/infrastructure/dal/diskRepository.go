package dal

import (
	"database/sql"
	"github.com/Alexander3006/architecture-lab-3/server/domain/entities"
)

type DiskRepository struct {
	Db *sql.DB
}

func (dr *DiskRepository) Create(d entities.Disk) error {
	_, err := dr.Db.Exec(
		`INSERT INTO "Disks" (space) VALUES ($1)`,
		d.Space,
		)
	return err
}

func (dr *DiskRepository) Delete(d entities.Disk) error {
	_, err := dr.Db.Exec(
		`DELETE FROM "Disks" WHERE id = $1`,
		d.Id,
		)
	return err
}

func (dr *DiskRepository) GetAll() ([]*entities.Disk, error) {
	rows, err := dr.Db.Query(`SELECT * FROM "Disks"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*entities.Disk
	for rows.Next() {
		var disk entities.Disk
		if err := rows.Scan(&disk.Id, &disk.Space); err != nil {
			return nil, err
		}
		res = append(res, &disk)
	}
	if res == nil {
		res = make([]*entities.Disk, 0)
	}
	return res, nil
}