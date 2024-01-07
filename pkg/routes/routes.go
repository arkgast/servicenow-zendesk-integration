package routes

import (
	"log"
	"os"

	"minka/support/pkg/common"

	"github.com/gofiber/fiber/v2"
)

var URL = "https://minkasupport.zendesk.com"

func CreateBundle(c *fiber.Ctx) error {
	log.Println("Creating zendesk bundle...")

	content, err := common.ReadFile("bundle.json")
	if err != nil {
		log.Println("Error reading bundle.json file")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	errs := createBundle(content)

	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": errs,
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func createBundle(content []byte) []error {
	url := URL + "/api/services/zis/registry/service-now-dev/bundles"
	username := os.Getenv("ZENDESK_USERNAME")
	password := os.Getenv("ZENDESK_PASSWORD")

	agent := fiber.Post(url)
	agent.Set("Content-Type", "application/json")
	agent.Body(content)
	agent.BasicAuth(username, password)

	statusCode, _, errs := agent.Bytes()

	log.Printf("Create bundle, response status code: %d", statusCode)

	return errs
}

func UpdateTask(c *fiber.Ctx) error {
	log.Println("Updating task...")

	common.PrettyPrintJSON(c.Body())

	return nil
}

func UpdateRequirement(c *fiber.Ctx) error {
	log.Println("Updating requirement...")

	common.PrettyPrintJSON(c.Body())

	return nil
}

func UpdateProblem(c *fiber.Ctx) error {
	log.Println("Updating problem...")

	common.PrettyPrintJSON(c.Body())

	return nil
}
