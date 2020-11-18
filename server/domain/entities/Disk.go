package entities

type Disk struct {
	Id int		`json:"diskId"`
	Space int 	`json:"space"`
	MachineId *int `json:"machineId"`
}
