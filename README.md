# Sample web application using Scylla

### Start Scylla DB using Docker

```
$ docker run --name hello-scylla -d scylladb/scylla
```

Get IP address:
```
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' hello-scylla 
```

### Build and run demo app

```
$ make
$ ./bin/hello --addr <scylla_ip_addr>
```

### Unit tests

```
$ make test
```

### Test Rest API

Create vehicles:
```
$ curl -v -H "Accept: application/json" -H "Content-type: application/json" localhost:3001/vehicles -d '{"vin":"vin1","engine":"Combustion"}'
$ curl -v -H "Accept: application/json" -H "Content-type: application/json" localhost:3001/vehicles -d '{"vin":"vin2","engine":"Ev", ev_data: {"battery_capacity_in_kwh": 62, "soc_in_percent": 74}}
$ curl -v -H "Accept: application/json" -H "Content-type: application/json" localhost:3001/vehicles -d '{"vin":"vin3","engine":"Phev"}}'
```

Find vehicles by vin:
```
$ curl -v -H "Accept: application/json" -H "Content-type: application/json" localhost:3001/vehicles -G --data-urlencode 'vin=vin2'
```

### Check database

```
$ docker exec -it hello-scylla nodetool status
$ docker exec -it hello-scylla cqlsh
cqlsh> USE hello;
cqlsh:hello> SELECT * from vehicles;
```