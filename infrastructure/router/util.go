package router

import (
	"context"
	"errors"
	"net/http"
	"taskflow/domain/entities"
)

type userKey string

const contextUserKey userKey = "user"

// WithUser adds the given user to the request's context'
func WithUser(ctx context.Context, user *entities.User) context.Context {
	return context.WithValue(ctx, contextUserKey, user)
}

// GetAppUser attempts to retrieve the app user in the request's context
func GetAppUser(r *http.Request) (*entities.User, error) {
	contextUser := r.Context().Value(contextUserKey)
	if contextUser == nil {
		return nil, errors.New("user not found in request")
	}

	return contextUser.(*entities.User), nil
}
