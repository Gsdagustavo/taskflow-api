package modules

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"taskflow/domain/entities"
	"taskflow/domain/status_codes"
	"taskflow/domain/usecases"
	"taskflow/infrastructure/util"

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
	ctx := r.Context()

	var credentials entities.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		util.WriteBadRequest(w)
		return
	}

	token, statusCode, err := a.authUseCases.AttemptLogin(r.Context(), credentials)
	if err != nil {
		slog.ErrorContext(ctx, "failed to login user", "cause", err)
		util.WriteInternalError(w)
		return
	}

	response := AuthResponse{
		Status:  statusCode.Int(),
		Message: statusCode.String(),
		Token:   token,
	}

	util.Write(w, response)
}

func (a AuthModule) register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var credentials entities.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		slog.ErrorContext(ctx, "failed to decode request body", "cause", err)
		util.WriteBadRequest(w)
		return
	}

	type Response struct {
		Status  status_codes.RegisterStatusCode `json:"status"`
		Message string                          `json:"message"`
	}

	statusCode, err := a.authUseCases.RegisterUser(r.Context(), credentials)
	if err != nil {
		slog.ErrorContext(ctx, "failed to register user", "cause", err)
		util.WriteInternalError(w)
		return
	}

	response := Response{
		Status:  statusCode,
		Message: statusCode.String(),
	}

	util.Write(w, response)
}
