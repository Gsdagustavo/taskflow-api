package config

import (
	"fmt"
	"log"
	"net/http"
	"taskflow/infrastructure/modules"

	"github.com/gorilla/mux"
)

type Server struct {
	Port    int    `toml:"port"`
	Host    string `toml:"host"`
	BaseURL string `toml:"base_url"`
	Router  *mux.Router
}

func NewServer() *Server {
	return &Server{
		Router: mux.NewRouter(),
	}
}

func (s Server) RegisterModules(modules ...modules.Module) {
	for _, module := range modules {
		module.RegisterRoutes(s.Router)
	}
}

func (s Server) Run(cfg Config) error {
	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Starting HTTP server on %s", address)

	s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	return http.ListenAndServe(address, s.Router)
}
