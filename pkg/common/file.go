package common

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ReadAndParseJSON(fileName string) (map[string]interface{}, error) {
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
