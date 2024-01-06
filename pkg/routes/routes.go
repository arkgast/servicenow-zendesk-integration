package routes

import (
	"log"
	"minka/support/pkg/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateBundle(c *fiber.Ctx) error {
	log.Println("Creating zendesk bundle...")

	content, err := common.ReadAndParseJSON("bundle.json")
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(content)
}
