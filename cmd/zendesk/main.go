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

	now := api.Group("/now")
	now.Patch("/table/sc_task/:id", routes.UpdateTask)
	now.Patch("/table/sc_req_item/:id", routes.UpdateRequirement)
	now.Patch("/table/problem/:id", routes.UpdateProblem)

	serviceNow := api.Group("/service-now")
	serviceNow.Post("/attachment/file", routes.AttachFile)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Listening on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting the server on port %s: %v", port, err)
	}
}
