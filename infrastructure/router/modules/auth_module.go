package modules

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
	"taskflow/domain/entities"
	"taskflow/domain/status_codes"
	"taskflow/domain/usecases"
	"taskflow/infrastructure/router"

	"github.com/gorilla/mux"
)

type AuthResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

type authModule struct {
	authUseCases usecases.AuthUseCases
	name         string
	path         string
}

func NewAuthModule(authUseCases usecases.AuthUseCases) router.Module {
	return authModule{
		authUseCases: authUseCases,
		name:         "Auth",
		path:         "/auth",
	}
}

func (a authModule) Name() string {
	return a.name
}

func (a authModule) Path() string {
	return a.path
}

func (a authModule) Setup(r *mux.Router) ([]router.RouteDefinition, *mux.Router) {
	defs := []router.RouteDefinition{
		{
			Path:        "/login",
			Description: "User login",
			Handler:     a.login,
			HttpMethods: []string{http.MethodPost},
		},
		{
			Path:        "/register",
			Description: "User registration",
			Handler:     a.register,
			HttpMethods: []string{http.MethodPost},
		},
	}

	for _, d := range defs {
		r.HandleFunc(a.path+d.Path, d.Handler).Methods(d.HttpMethods...)
	}

	r.Use(a.sessionMiddleware)

	return defs, r
}

func (a authModule) sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		authHeader := r.Header.Get("Authorization")

		var user *entities.User

		// Test basic auth
		email, password, ok := r.BasicAuth()
		if ok {
			_ = entities.UserCredentials{
				Email:    email,
				Password: password,
			}

			//valid, err := CheckCredentials(ctx, credentials)
			//if err != nil {
			//	slog.ErrorContext(ctx, "failed to check credentials", "cause", err)
			//	router.WriteUnauthorized(w)
			//	return
			//}
			//
			//if !valid {
			//	slog.ErrorContext(ctx, "invalid credentials")
			//	router.WriteUnauthorized(w)
			//	return
			//}
			//
			//user, err := a.authUseCases.GetUserByEmail(ctx, credentials.Email)
			//if err != nil {
			//	slog.ErrorContext(ctx, "failed to get user by document", "cause", err)
			//	router.WriteUnauthorized(w)
			//	return
			//}

			if user == nil {
				slog.ErrorContext(ctx, "user not found")
				router.WriteUnauthorized(w)
				return
			}
		} else {
			var token string
			//var err error

			if authHeader != "" {
				token = strings.ReplaceAll(authHeader, "Bearer ", "")
			}

			if token == "" {
				slog.ErrorContext(ctx, "no token found in the request")
				router.WriteUnauthorized(w)
				return
			}

			//user, err = a.authUseCases.GetUserFromToken(ctx, token)
			//if err != nil {
			//	slog.ErrorContext(ctx, "failed to get user from token", "cause", err)
			//	router.WriteUnauthorized(w)
			//	return
			//}

			if user == nil {
				slog.ErrorContext(ctx, "user not found")
				router.WriteUnauthorized(w)
				return
			}
		}

		next.ServeHTTP(w, r.WithContext(router.WithUser(ctx, user)))
	})
}

func (a authModule) login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var credentials entities.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		slog.ErrorContext(ctx, "failed to decode request body", "cause", err)
		router.WriteBadRequest(w)
		return
	}

	token, statusCode, err := a.authUseCases.AttemptLogin(r.Context(), credentials)
	if err != nil {
		slog.ErrorContext(ctx, "failed to login user", "cause", err)
		router.WriteInternalError(w)
		return
	}

	response := AuthResponse{
		Status:  statusCode.Int(),
		Message: statusCode.String(),
		Token:   token,
	}

	router.Write(w, response)
}

func (a authModule) register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var credentials entities.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		slog.ErrorContext(ctx, "failed to decode request body", "cause", err)
		router.WriteBadRequest(w)
		return
	}

	type Response struct {
		Status  status_codes.RegisterStatusCode `json:"status"`
		Message string                          `json:"message"`
	}

	statusCode, err := a.authUseCases.RegisterUser(r.Context(), credentials)
	if err != nil {
		slog.ErrorContext(ctx, "failed to register user", "cause", err)
		router.WriteInternalError(w)
		return
	}

	response := Response{
		Status:  statusCode,
		Message: statusCode.String(),
	}

	router.Write(w, response)
}
