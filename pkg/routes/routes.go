package routes

import (
	"fmt"
	"log"
	"minka/support/pkg/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateBundle(c *fiber.Ctx) error {
	log.Println("Creating zendesk bundle...")

	fileContent, err := common.ReadAndParseJSON("bundle.json")
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	fmt.Printf("JSON: %v\n", fileContent)

	return c.SendStatus(http.StatusOK)
}
