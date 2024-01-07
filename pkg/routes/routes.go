package routes

import (
	"log"
	"net/http"

	"minka/support/pkg/common"

	"github.com/gofiber/fiber/v2"
)

func CreateBundle(c *fiber.Ctx) error {
	log.Println("Creating zendesk bundle...")

	content, err := common.ReadAndParseJSON("bundle.json")
	if err != nil {
		log.Println("Error reading bundle.json file")
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(content)
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
