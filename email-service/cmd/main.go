package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/popstr/micro-world/email-service/internal/email"
	proto "github.com/popstr/micro-world/email-service/proto"
	"os"
)

// INACTIVATED go:generate protoc --proto_path=../../proto --micro_out=../../common/email --go_out=../../common/email ../../proto/email.proto
//go:generate protoc --go_out=paths=source_relative:../proto --micro_out=paths=source_relative:../proto --proto_path=../../proto ../../proto/email.proto

func main() {
	println("email-service starting")

	service := micro.NewService(
		micro.Name("email-service"),
		micro.Version("latest"),
	)
	service.Init()

	emailService := email.New()
	if err := emailService.ReadUsersFile("./data/users.csv"); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	emailHandler := email.NewHandler(emailService)

	if err := proto.RegisterEmailHandler(service.Server(), emailHandler); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
