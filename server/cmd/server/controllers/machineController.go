package controllers

import (
	"encoding/json"
	"github.com/Alexander3006/architecture-lab-3/server/domain/entities"
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
	"log"
)

//post
func (c Controllers) createMachine(connection interfaces.IConnection) {
	var m entities.Machine
	if err := json.NewDecoder(connection.GetRequest().Body).Decode(&m); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		connection.SendError(400, "bad Json payload")
	}
	err := c.machineService.CreateMachine(m.Name, m.CpuCount)
	if err != nil {
		log.Printf("Error creating machine: %s", err)
		connection.SendError(500, "server error")
	}
}

//delete
func (c Controllers) deleteMachine(connection interfaces.IConnection) {
}

func (c Controllers) getMachineById(connection interfaces.IConnection) {

}

func (c Controllers) getAllMachines(connection interfaces.IConnection) {

}

