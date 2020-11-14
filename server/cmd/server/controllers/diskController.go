package controllers

import (
	"github.com/Alexander3006/architecture-lab-3/server/domain/entities"
	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
	"log"
)

//post
func (c Controllers) createDisk(connection interfaces.IConnection) {
	var disk entities.Disk
	if err := connection.GetBody(&disk); err!= nil {
		log.Printf("Error decoding disk: %s", err)
		return
	}
	if err := c.diskServices.CreateDisk(disk.Space); err != nil {
		log.Printf("Error creating disk: %s", err)
		connection.SendError(409, "disk cannot be created")
		return
	}
}

//delete
func (c Controllers) deleteDisk(connection interfaces.IConnection) {
	var disk entities.Disk
	if err := connection.GetBody(&disk); err != nil {
		log.Printf("Error decoding disk: %s", err)
		return
	}
	if err := c.diskServices.DeleteDisk(disk.Id); err != nil {
		log.Printf("Error delete disk: %s", err)
		connection.SendError(409, "disk cannot be deleted")
		return
	}
}

//get
func (c Controllers) getAllDisks(connection interfaces.IConnection) {
	disks, err := c.diskServices.GetAllDisks()
	if err != nil {
		connection.SendError(500, "server error")
	}
	connection.Ok(disks)
}