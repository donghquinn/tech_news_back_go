package utilities

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseResponse(response *http.Response, instance interface{}) error {
	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(instance)

	if err != nil {
		log.Printf("Parsing and Decode JSON Error: %v", err)
		return err
	}

	defer response.Body.Close()

	return nil
}

func ParseBody(req *http.Request, instance interface{}) error {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(instance)

	if err != nil {
		log.Printf("Parsing and Decode JSON Error: %v", err)
		return err
	}

	defer req.Body.Close()

	return nil
}