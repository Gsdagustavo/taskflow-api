package modules

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type HealthModule struct {
	name string
	path string
}

func NewHealthModule() *HealthModule {
	return &HealthModule{
		name: "health",
		path: "/health",
	}
}

func (h HealthModule) Name() string {
	return h.name
}

func (h HealthModule) Path() string {
	return h.path
}

func (h HealthModule) RegisterRoutes(router *mux.Router) {
	routes := []ModuleRoute{
		{
			Name:    "health",
			Path:    h.path,
			Handler: h.health,
			Methods: []string{http.MethodGet},
		},
	}

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Methods...)
	}
}

func (h HealthModule) health(w http.ResponseWriter, r *http.Request) {
	type HealthResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	response := HealthResponse{
		Status:  "OK",
		Message: "Server is healthy",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
