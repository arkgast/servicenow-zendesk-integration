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

	erros := installJobSpec()
	if len(erros) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": erros,
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

func installJobSpec() []error {
	url := URL + "/api/services/zis/registry/job_specs/install"
	username := os.Getenv("ZENDESK_USERNAME")
	password := os.Getenv("ZENDESK_PASSWORD")

	agent := fiber.Post(url)
	agent.Set("Content-Type", "application/json")
	agent.BasicAuth(username, password)
	agent.QueryString("job_spec_name=zis:service-now-dev:job_spec:UpdateTicketSpec")

	statusCode, _, errs := agent.Bytes()

	log.Printf("Install job spec, response status code: %d", statusCode)

	return errs
}

func UpdateTask(c *fiber.Ctx) error {
	log.Println("Updating task...")

	log.Println("ID: ", c.Params("id"))
	log.Println("Authorization: ", c.Get("Authorization"))
	common.PrettyPrintJSON(c.Body())

	return nil
}

func UpdateRequirement(c *fiber.Ctx) error {
	log.Println("Updating requirement...")

	log.Println("ID: ", c.Params("id"))
	log.Println("Authorization: ", c.Get("Authorization"))
	common.PrettyPrintJSON(c.Body())

	return nil
}

func UpdateProblem(c *fiber.Ctx) error {
	log.Println("Updating problem...")

	log.Println("ID: ", c.Params("id"))
	log.Println("Authorization: ", c.Get("Authorization"))
	common.PrettyPrintJSON(c.Body())

	return nil
}

func AttachFile(c *fiber.Ctx) error {
	log.Println("Attaching file...")

	log.Println("Content-Type: ", c.Get("Content-Type"))

	log.Println("Query Params:")

	// Accessing query parameters
	tableName := c.Query("table_name")
	tableSysID := c.Query("table_sys_id")
	filename := c.Query("file_name")
	log.Println("table_name:", tableName)
	log.Println("table_sys_id:", tableSysID)
	log.Println("file_name:", filename)

	// Parse the form data, including files
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error parsing form file")
	} else {
		log.Println("Received file: ", file.Filename)
		log.Println("File size: ", file.Size)
	}

	return nil
}
