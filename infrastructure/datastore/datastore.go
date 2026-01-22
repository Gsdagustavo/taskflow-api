package datastore

import (
	"context"
	"database/sql"
	"taskflow/domain/entities"
	"time"

	"github.com/google/uuid"
)

// RepositorySettings creates and manages the access for the database
type RepositorySettings interface {
	// Connection returns a database connection
	Connection() *sql.DB

	// Dismount closes all connections with the database
	Dismount() error

	// ServerTime returns the current time on the server
	ServerTime(ctx context.Context) (*time.Time, error)
}

type AuthRepository interface {
	AddUser(ctx context.Context, user *entities.User) error
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserByID(ctx context.Context, id int) (*entities.User, error)
	GetUserByUUID(ctx context.Context, uuid uuid.UUID) (*entities.User, error)
	DeleteUser(ctx context.Context, id int) error
}
