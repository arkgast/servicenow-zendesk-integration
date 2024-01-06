package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func CreateBundle(c *fiber.Ctx) error {
	log.Println("Creating zendesk bundle...")

	fileContent, err := readAndParseJSON("bundle.json")
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	fmt.Printf("JSON: %v\n", fileContent)

	return c.SendStatus(http.StatusOK)
}

func readAndParseJSON(fileName string) (map[string]interface{}, error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filePath := filepath.Join(cwd, "data", fileName)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	jsonData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var parsedData map[string]interface{}
	if err := json.Unmarshal(jsonData, &parsedData); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return parsedData, err
}
