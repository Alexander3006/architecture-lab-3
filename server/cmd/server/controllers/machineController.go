package controllers

import (
	"github.com/Alexander3006/architecture-lab-3/server/domain/entities"
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
	"log"
	"strconv"
)

//post
func (c Controllers) createMachine(connection interfaces.IConnection) {
	var m entities.Machine
	if err := connection.GetBody(&m); err != nil {
		log.Printf("Error decoding machine: %s", err)
		return
	}
	err := c.machineService.CreateMachine(m.Name, m.CpuCount)
	if err != nil {
		log.Printf("Error creating machine: %s", err)
		connection.SendError(409, "machine cannot be created")
	}
}

//delete
func (c Controllers) deleteMachine(connection interfaces.IConnection) {
	var m entities.Machine
	if err := connection.GetBody(&m); err != nil {
		log.Printf("Error decoding machine: %s", err)
		return
	}
	if err := c.machineService.DeleteMachine(m.Id); err != nil {
		log.Printf("Error delete machine: %s", err)
		connection.SendError(409, "machine cannot be deleted")
		return
	}
}

//Get
func (c Controllers) getMachineById(connection interfaces.IConnection) {
	paramId := connection.GetRequest().URL.Query().Get("id");
	id, err := strconv.Atoi(paramId)
	if err != nil {
		connection.SendError(400, "bad request")
		return
	}
	if machine, _ := c.machineService.GetMachine(id); machine != nil {
		connection.Ok(machine)
		return
	}
	connection.SendError(404, "not found")
}

//Get
func (c Controllers) getAllMachines(connection interfaces.IConnection) {
	machines, err := c.machineService.GetAllMachines()
	if err != nil {
		connection.SendError(500, "server error")
	}
	connection.Ok(machines)
}

//Post
func (c Controllers) addDisk(connection interfaces.IConnection) {
	var component entities.Component
	if err := connection.GetBody(&component); err != nil {
		log.Printf("Error decoding component %s: ", err)
		return
	}
	err := c.machineService.AddComponent(component.MachineId, component.DiskId)
	if err != nil {
		log.Printf("Error adding component to machine: %s", err)
		connection.SendError(409, "component cannot be added")
		return
	}
}

//Delete
func (c Controllers) removeDisk(connection interfaces.IConnection) {
	var component entities.Component
	if err := connection.GetBody(&component); err != nil {
		log.Printf("Error decoding component: %s", err)
		return
	}
	err := c.machineService.RemoveComponent(component.MachineId, component.DiskId)
	if err != nil {
		log.Printf("Error removing component to machine: %s", err)
		connection.SendError(409, "component cannot be removed")
		return
	}

}
