package util

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, error interface{}) {
	w.Header().Set("Content-Type", "application/json")

	var statusCode int
	var response map[string]interface{}

	switch e := error.(type) {
	case struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}:
		statusCode = e.Status
		response = map[string]interface{}{
			"status":  e.Status,
			"message": e.Message,
		}
	default:
		statusCode = http.StatusInternalServerError
		response = map[string]interface{}{
			"status":  statusCode,
			"message": error,
		}
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
