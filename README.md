## Usage

```
$ docker-compose up
$ docker-compose exec app bash -c "cd /app/grpc_server_go && go get -v"
```

## Run

### server(golang)

```
$ docker-compose exec app go run /app/grpc_server_go/main.go
```

### client(golang)

```
$ docker-compose exec app go run /app/grpc_client_go/main.go
```


### generate golang pb file

```
$ docker-compose exec app bash
# in container
$ protoc --go_out=plugins=grpc:. helloworld.proto
```
