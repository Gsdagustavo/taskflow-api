package util

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ServerResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Write(w http.ResponseWriter, data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		slog.Error("failed to marshal response", "error", err)
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		slog.Error("failed to write response", "error", err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
}

func WriteInternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("Internal server error"))
}

func WriteBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte("Bad request"))
}

func WriteForbidden(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	_, _ = w.Write([]byte("Forbidden"))
}
