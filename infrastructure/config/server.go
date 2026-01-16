package config

import (
	"fmt"
	"log/slog"
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
	slog.Info("starting server", "address", address)

	s.Router.Use(CORSMiddleware)

	s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	return http.ListenAndServe(address, s.Router)
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization")
		w.Header().Set("Vary", "Origin")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
