package router

import (
	"encoding/json"
	"errors"
	"net/http"
)

func Write(w http.ResponseWriter, v any) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return errors.New("failed to marshal response body")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(bytes)
	if err != nil {
		return errors.New("failed to write response body")
	}

	return nil
}

func WriteInternalError(w http.ResponseWriter) {
	http.Error(w, "Internal server error", http.StatusInternalServerError)
	w.WriteHeader(http.StatusInternalServerError)
}

func WriteBadRequest(w http.ResponseWriter) {
	http.Error(w, "Bad request", http.StatusBadRequest)
	w.WriteHeader(http.StatusBadRequest)
}

func WriteUnauthorized(w http.ResponseWriter) {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	w.WriteHeader(http.StatusUnauthorized)
}

func WriteForbidden(w http.ResponseWriter) {
	http.Error(w, "Forbidden", http.StatusForbidden)
	w.WriteHeader(http.StatusForbidden)
}
