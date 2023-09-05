package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// ResponseWithJson sends a JSON response to the client with the specified status code and payload.
func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	// Marshal the payload into JSON format.
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err) // Log any marshaling errors.
		return
	}
	
	// Set the response header to indicate JSON content.
	w.Header().Add("Content-Type", "application/json")
	
	// Write the JSON data to the response writer.
	w.Write(data)
}

// ResponseWithError sends a JSON error response to the client with the specified status code and error message.
func ResponseWithError(w http.ResponseWriter, code int, err error) {
	if err != nil {
		// If an error is provided, respond with the error message in JSON format.
		ResponseWithJson(w, code, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
	} else {
		// If no error is provided, respond with a generic internal server error message in JSON format.
		ResponseWithJson(w, code, struct {
			Error string `json:"error"`
		}{
			Error: "Internal server error",
		})
	}
}
