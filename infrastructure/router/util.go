package router

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
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

// GetQueryInt attempts to parse a request's query parameter as an integer
//
// If it succeeds, returns a pointer to the integer; otherwise, returns nil
func GetQueryInt(query url.Values, key string) *int64 {
	value := query.Get(key)
	if value != "" {
		parsed, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return &parsed
		}
	}

	return nil
}

// GetQueryIntOr attempts to parse a request's query parameter as an integer
//
// If it succeeds, returns a pointer to the integer; otherwise, returns the given default value
func GetQueryIntOr(query url.Values, key string, defaultValue int64) int64 {
	value := query.Get(key)
	if value != "" {
		parsed, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return parsed
		}
	}

	return defaultValue
}

// GetQueryBool parses a request's query parameter as a bool
//
// Returns:
//   - true, if query is exactly "true"
//   - false, if query is exactly "false"
//   - nil otherwise
func GetQueryBool(query url.Values, key string) *bool {
	value := query.Get(key)
	if value == "true" {
		result := true
		return &result
	} else if value == "false" {
		result := false
		return &result
	}
	return nil
}

// GetQueryBoolOr parses a request's query parameter as a bool
//
// Returns:
//   - true, if query is exactly "true"
//   - false, if query is exactly "false"
//   - defaultValue otherwise
func GetQueryBoolOr(query url.Values, key string, defaultValue bool) bool {
	value := query.Get(key)
	if value == "true" {
		return true
	} else if value == "false" {
		return false
	}
	return defaultValue
}
