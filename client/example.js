
// the SDK for channels management.

const channels = require('./machines/client');

const client = channels.Client('http://localhost:3000');

// Scenario 1: Display all machines.
{
    client.getMachines()
        .then((list) => {
            console.log('=== Scenario 1 ===');
            console.log('All machines:');
            console.log(list);
        })
        .catch((e) => {
            console.log(`Problem displaying all machines: ${e.message}`);
        });
}

// Scenario 2: Get machine by Id.
{
    const machineId = 3;
    client.getMachineById(machineId)
        .then((resp) => {
            console.log('=== Scenario 2 ===');
            console.log(`Get machine by Id = ${machineId}`);
            console.log(resp);
        })
        .catch((e) => {
            console.log(`Problem displaying a machine: ${e.message}`);
        });
}

// Scenario 3: Get all disks.
{
    client.getDisks()
        .then((list) => {
            console.log('=== Scenario 3 ===');
            console.log('Get all disks');
            console.log(list);
        })
        .catch((e) => {
            console.log(`Problem displaying all disks: ${e.message}`);
        });
}

// Scenario 4: Create machine.
{
    const machineName = 'server-1';
    const cpuCount = 4;
    client.createMachine(machineName, cpuCount)
        .then(() => {
            console.log('=== Scenario 4 ===');
            console.log(`Create machine:${machineName}`);
        })
        .catch((e) => {
            console.log(`Problem creating a new machine: ${e.message}`);
        });
}

// Scenario 5: Create disk.
{
    const space = 1000;
    client.createDisk(space)
        .then(() => {
            console.log('=== Scenario 5 ===');
            console.log(`Create disk with space:${space}`);
        })
        .catch((e) => {
            console.log(`Problem creating a new disk: ${e.message}`);
        });
}

// Scenario 6: Add disk to machine.
{
    const machineId = 2;
    const diskId = 2;
    client.addDiskToMachine(machineId, diskId)
        .then(() => {
            console.log('=== Scenario 6 ===');
            console.log(`Add disk id:${diskId} to machine id:${machineId}`);
        })
        .catch((e) => {
            console.log(`Problem adding disk to machine: ${e.message}`);
        });
}

// Scenario 7: Remove disk from machine.
{
    const machineId = 6;
    const diskId = 4;
    client.removeDiskFromMachine(machineId, diskId)
        .then(() => {
            console.log('=== Scenario 7 ===');
            console.log(`Remove disk id:${diskId} from machine id:${machineId}`);
        })
        .catch((e) => {
            console.log(`Problem removing disk to machine: ${e.message}`);
        });
}

// Scenario 8: Delete machine.
{
    const machineId = 5;
    client.deleteMachine(machineId)
        .then(() => {
            console.log('=== Scenario 8 ===');
            console.log(`Delete machine id:${machineId}`);
        })
        .catch((e) => {
            console.log(`Problem deleting a machine: ${e.message}`);
        });
}

// Scenario 9: Delete disk.
{
    const diskId = 8
    client.deleteDisk(diskId)
        .then(() => {
            console.log('=== Scenario 9 ===');
            console.log(`Delete disk id:${diskId}`);
        })
        .catch((e) => {
            console.log(`Problem deleting a disk: ${e.message}`);
        });
}
