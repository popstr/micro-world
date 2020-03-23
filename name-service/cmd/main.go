package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/popstr/micro-world/name-service/internal/name"
	proto "github.com/popstr/micro-world/name-service/proto"
	"os"
)

//go:generate protoc --proto_path=../../ --micro_out=../../common/names --go_out=../../common/names ../../proto/names.proto

func main() {
	println("name-service starting")

	service := micro.NewService(
		micro.Name("name-service"),
		micro.Version("latest"),
	)
	service.Init()

	nameService := name.New()
	if err := nameService.ReadUsersFile("./data/users.csv"); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	nameHandler := name.NewHandler(nameService)

	if err := proto.RegisterNamesHandler(service.Server(), nameHandler); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
