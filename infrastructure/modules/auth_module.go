package modules

import (
	"encoding/json"
	"net/http"
	"taskflow/domain/entities"
	"taskflow/domain/status_codes"
	"taskflow/domain/usecases"

	"github.com/gorilla/mux"
)

type AuthResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

type AuthModule struct {
	authUseCases usecases.AuthUseCases
	name         string
	path         string
}

func NewAuthModule(authUseCases usecases.AuthUseCases) *AuthModule {
	return &AuthModule{
		authUseCases: authUseCases,
		name:         "auth",
		path:         "/auth",
	}
}

func (a AuthModule) Name() string {
	return a.name
}

func (a AuthModule) Path() string {
	return a.path
}

func (a AuthModule) RegisterRoutes(router *mux.Router) {
	routes := []ModuleRoute{
		{
			Name:    "Login",
			Path:    a.path + "/login",
			Handler: a.login,
			Methods: []string{http.MethodPost},
		},
		{
			Name:    "Register",
			Path:    a.path + "/register",
			Handler: a.register,
			Methods: []string{http.MethodPost},
		},
	}

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Methods...)
	}
}

func (a AuthModule) login(w http.ResponseWriter, r *http.Request) {
	var credentials entities.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, statusCode, err := a.authUseCases.AttemptLogin(r.Context(), credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := AuthResponse{
		Status:  statusCode.Int(),
		Message: statusCode.String(),
		Token:   token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		// At this point headers are already sent, log the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a AuthModule) register(w http.ResponseWriter, r *http.Request) {
	var credentials entities.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type Response struct {
		Status  status_codes.RegisterStatusCode `json:"status"`
		Message string                          `json:"message"`
	}

	statusCode, err := a.authUseCases.RegisterUser(r.Context(), credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Status:  statusCode,
		Message: statusCode.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		// At this point headers are already sent, log the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
