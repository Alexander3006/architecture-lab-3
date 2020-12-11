const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        getMachines: () => client.get('/api/machine/all'),
        getMachineById: id => client.get(`/api/machine?id=${id}`),
        getDisks: () => client.get('/api/disk/all'),
        createMachine: (name, cpuCount) => client.post('/api/machine/create', { name, cpuCount }),
        createDisk: space => client.post('/api/disk/create', { space }),
        addDiskToMachine: (machineId, diskId) => client.post('/api/machine/addDisk', { machineId, diskId }),
        deleteMachine: id => client.delete('/api/machine/delete', { id }),
        removeDiskFromMachine: (machineId, diskId) => client.delete('/api/machine/removeDisk', { machineId, diskId }),
        deleteDisk: id => client.delete('/api/disk/delete', { id })
    }

};

module.exports = { Client };