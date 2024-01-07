package main

import (
	"log"
	"minka/support/pkg/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	api := app.Group("/api")

	zendesk := api.Group(("/zendesk"))
	zendesk.Post("/create-zis", routes.CreateBundle)

	serviceNow := api.Group("/service-now")
	serviceNow.Patch("/sc_task/:id", routes.UpdateTask)
	serviceNow.Patch("/sc_req_item/:id", routes.UpdateRequirement)
	serviceNow.Patch("/problem/:id", routes.UpdateProblem)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Listening on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting the server on port %s: %v", port, err)
	}
}
