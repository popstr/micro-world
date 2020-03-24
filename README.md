# Micro World

This is an example micro service environment build with Go using the [go-micro](https://github.com/micro/go-micro) framework.

The point of this repo is to show how simple it is to implement efficient and lean micro services using Go.

## Services

The environment consists of 3 different services. One service faces the outside world while the other have their own responsibilities and does not communicate outside the (imagined) cluster.

### API Service

This is a classic REST API built using [Echo](https://github.com/labstack/echo). It faces a (imagined) frontend application what wants to display a page with all information about a user given a user id. To gather this information, the API Service communicates with the other two services in this repo - Name Service and Email Service. This inter-service communication uses gRPC.

The one and only API endpoint this service has is:

`GET http://localhost:8080/api/users/{userID}`

Example reply:
```json
{
  "users": {
    "1": {
      "firstName": "Mikael",
      "lastName": "Karlsson",
      "email": "micke@mymail.com"
    }
  }
}
```

The system currently has 3 users with user id's 1, 2 and 3.

### Name Service

This service has one gRPC endpoint that returns the firstname and lastname of a user given one or more user id's.

### Email Service

As you might have guessed by now, this service has one gRPC endpoint that returns the email address of a user given one or more user id's.

## Prerequisites

Install Protobuf: https://micro.mu/docs/runtime.html#protobuf

## Application Data

User databases for the Name and Email services are emulated using two `.csv` files, stored in the respective project's `data` folder.

## Communication

The gRPC communication protocol is defined in the two `.proto` files, located in the [proto folder](https://github.com/popstr/micro-world/tree/master/proto). These files are used to generate the source code for implementing both the server and the client.

Each service can generate the files it requires. This is done using the [Go Generate](https://blog.golang.org/generate) functionality.

I have made it simple to generate code for all services with one command using `make`. To generate all required source files, run this command in the project root folder:

```bash
make gen
```

To generate code for one service only, run this from the service's `cmd` folder:

```bash
go generate
```

## Future Improvements
 
The intention is to keep this project simple. Future additions must therefore not obfuscate the original message too much.

Things I might add:

* Dockerfiles
* 2 different databases for the name and email services (one SQL and one NoSQL)
* Protobuf compilation using docker image, https://github.com/TheThingsIndustries/docker-protobuf
* https://github.com/go-swagger/go-swagger
* Make it easy to use Etcd as the service discovery system

