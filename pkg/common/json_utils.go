package common

import (
	"bytes"
	"encoding/json"
	"log"
)

func PrettyPrintJSON(rawBody []byte) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, rawBody, "", " "); err != nil {
		log.Println("Error pretty printing JSON")
	}

	log.Println(prettyJSON.String())
}
