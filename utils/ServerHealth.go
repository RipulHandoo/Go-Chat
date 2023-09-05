package utils

import (
	"fmt"
	"net/http"

	"github.com/RipulHandoo/goChat/db"
)

// resp represents a simple response structure with a "Status" field.
type resp struct {
	Status string
}

// ServerHealth is a handler function for a health check endpoint.
func ServerHealth(w http.ResponseWriter, r *http.Request) {
	// Get the database instance.
	db := db.DBInstance()

	// Check if the database instance is nil (indicating a database error).
	if db == nil {
		// If there's a database error, respond with an error message.
		ResponseWithError(w, http.StatusBadRequest, fmt.Errorf("database error"))
		return
	}

	// If the database is reachable, respond with a JSON message indicating "OK" status.
	ResponseWithJson(w, 200, resp{
		Status: "OK",
	})
}
