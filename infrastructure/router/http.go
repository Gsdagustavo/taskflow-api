package router

import (
	"net/http"
)

func Write(w http.ResponseWriter, v []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(v)
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
