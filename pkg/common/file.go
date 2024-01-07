package common

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ReadFile(fileName string) ([]byte, error) {
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

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return data, nil
}

func ParseJSON(data []byte) (map[string]interface{}, error) {
	var parsedData map[string]interface{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return parsedData, nil
}
