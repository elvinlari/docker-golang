# Golang CRUD Application Scaffold

## Usage

### API

 - GET http://localhost:8080/rest/todos/
 - GET http://localhost:8080/rest/todos/1
 - POST http://localhost:8080/rest/todos/ { "name": "TodoName" }
 - PUT http://localhost:8080/rest/todos/1 { "name": "TodoName" }
 - DELETE http://localhost:8080/rest/todos/1
 - DELETE http://localhost:8080/rest/todos/

### docker-compose

If you want to setup docker environemnt just use `./scripts/docker-compose.yml` with [docker-compose](https://docs.docker.com/compose/).

Go to `./scripts` directory and execute

```
# start docker environment
$ docker-compose up -d (--build)

# list running services
$ docker-compose ps

# stop all containers
$ docker-compose stop

# remove all containers
$ docker-compose rm
```
 
### single docker
 - Build app: &docker build -t golang:go-app .  && $docker run --name go-crud --network=host -it -d -p 8080:8080 golang:go-app
 - You can run localDB with: $docker run -d -p 3306:3306 --name mysql-db -e MYSQL_ROOT_PASSWORD=admin -d mysql:5.7
 
### Run Locally:
 - start mysql on 3306 port (and execute db script - manually currently)
 - start app from IDE or after install with flag -env=dev and giving -configFilePath, ./go-structure -env=dev
 
