package main

import (
	"github.com/gofiber/fiber/v2"
	"minka/support/pkg/routes"
)

func main() {
	app := fiber.New()

	app.Get("/", routes.CreateBundle)

	app.Listen(":3000")
}
