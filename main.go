package main

import (
	"github/HsiaoCz/leaf/etc"
	"github/HsiaoCz/leaf/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := etc.Init()
	if err != nil {
		log.Fatal(err)
	}
	r := fiber.New()
	router.Router(r)
	r.Listen(etc.Conf.App.AppPort)
}
