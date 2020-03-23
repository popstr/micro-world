package main

import (
	"github.com/labstack/echo"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
	emailproto "github.com/popstr/micro-world/email-service/proto"
	nameproto "github.com/popstr/micro-world/name-service/proto"
	"github.com/postr/micro-world/api-service/internal/email"
	"github.com/postr/micro-world/api-service/internal/name"
	"github.com/postr/micro-world/api-service/internal/user"
	"net/http"
)

func main() {
	println("api-service starting")

	e := echo.New()

	names := nameproto.NewNamesService("name-service", client.DefaultClient)
	namesClient := name.NewNamesClient(names)

	emailService := emailproto.NewEmailService("email-service", client.DefaultClient)
	emailClient := email.NewEmailClient(emailService)

	api := e.Group("/api")
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	api.GET("/users/:id", user.HandleGetUser(namesClient, emailClient))
	e.Static("/", "html")

	service := web.NewService(
		web.Handler(e),
		web.Address("0.0.0.0:8080"),
	)
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
