package entities

type Machine struct {
	Id int 				`json:"id"`
	Name string			`json:"name"`
	CpuCount int		`json:"cpuCount"`
	TotalDiskSpace int	`json:"totalDiskSpace"`
}