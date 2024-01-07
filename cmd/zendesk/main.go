package main

import (
	"github.com/gofiber/fiber/v2"
	"minka/support/pkg/routes"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")

	zendesk := api.Group(("/zendesk"))
	zendesk.Post("/create-zis", routes.CreateBundle)

	serviceNow := api.Group("/service-now")
	serviceNow.Patch("/sc_task/:id", routes.UpdateTask)
	serviceNow.Patch("/sc_req_item/:id", routes.UpdateRequirement)
	serviceNow.Patch("/problem/:id", routes.UpdateProblem)

	app.Listen(":3000")
}
