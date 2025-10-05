package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"taskflow/domain/entities"

	"github.com/google/uuid"
)

type MySQLAuthRepository struct {
	db *sql.DB
}

func NewMySQLAuthRepository(db *sql.DB) *MySQLAuthRepository {
	return &MySQLAuthRepository{
		db: db,
	}
}

func (r MySQLAuthRepository) AddUser(ctx context.Context, user *entities.User) error {
	const query = `
		INSERT INTO users (uuid, name, email, password) VALUES (?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(ctx, query, user.UUID, user.Name, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("error adding user: %s", err)
	}

	return nil
}

func (r MySQLAuthRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	const query = `
		SELECT id, uuid, name, email, password FROM users WHERE email = ?
	`

	var user entities.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error getting user by email: %s", err)
	}

	return &user, nil
}

func (r MySQLAuthRepository) GetUserByID(ctx context.Context, id int) (*entities.User, error) {
	const query = `
		SELECT id, uuid, name, email, password FROM users WHERE id = ?
	`

	var user entities.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error getting user by email: %s", err)
	}

	return &user, nil
}

func (r MySQLAuthRepository) GetUserByUUID(ctx context.Context, uuid uuid.UUID) (*entities.User, error) {
	const query = `
		SELECT id, uuid, name, email, password FROM users WHERE id = ?
	`

	var user entities.User
	err := r.db.QueryRowContext(ctx, query, uuid).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error getting user by email: %s", err)
	}

	return &user, nil
}
func (r MySQLAuthRepository) DeleteUser(ctx context.Context, id int) error {
	const query = `
		DELETE FROM users WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %s", err)
	}

	return nil
}
