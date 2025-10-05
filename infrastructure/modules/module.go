package modules

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Module interface {
	Name() string
	Path() string
	RegisterRoutes(router *mux.Router)
}

type ModuleRoute struct {
	Name    string
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
	Methods []string
}
