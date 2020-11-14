## Start server: 
```
go run .\cmd\server
```

## Server APIs:

#### ♦ Response types:  

```
Machine {
	id: int
	name: string
	cpuCount: int
	totalDiskSpace: int
}

Disk {
	id: int
	space: int
}
```

#### ♦ Get machine by id:

```
request: {
	url: "/api/machine"
	query: "?id="
	method: "GET"
}
response : {
	200: {
    	body: Machine
    	}
    	400: "bad request"
    	404: "not found"
}
```

#### ♦ Get all machines:

```
request: {
	url: "/api/machine/all"
    	method: "GET"
}
response {
	200: {
    	body: []Machine
    	}
    	500: "server error"
}
```

#### ♦ Create machine:

```
request: {
	url: "/api/machine/create"
   	method: "POST"
    	body: {
    		name: int
        	cpuCount: int
    }
}
response: {
	200
    	409: "machine cannot be created"
}
```

#### ♦ Add disk to machine: 

```
request: {
	url: "/api/machine/addDisk"
    	method: "POST",
    	body: {
    		machineId: int
        	diskId: int
    }
}
response: {
	200
    	409: "disk cannot be added"
}
```

#### ♦ Delete machine: 

```
request: {
	url: "/api/machine/delete"
    	method: "DELETE"
    	body: {
    		id: int
    }
}
response: {
	200
    	409: "machine cannot be deleted"
}
```

#### ♦ Remove disk from machine: 

```
request: {
	url: "/api/machine/removeDisk"
    	method: "DELETE"
    	body: {
    		machineId: int
        	diskId: int
    }
}
response: {
	200
    	409: "component cannot be removed"
}
```

#### ♦ Get all disks: 

```
request: {
	url: "/api/disk/all"
    	method: "GET"
}
response: {
	200: []disk
    	500: "server error"
}
```

#### ♦ Create disk:

```
request: {
	url: "/api/disk/create"
    	method: "POST"
    	body: {
    		space: int
    }
}
response: {
	200
    	409: "disk cannot be created"
}
```

#### ♦ Delete disk:

```
request: {
	url: "/api/disk/delete"
    	method: "DELETE"
    	body: {
    		id: int
    	}
}
response: {
	200
    	409: "disk cannot be deleted"
}
```
