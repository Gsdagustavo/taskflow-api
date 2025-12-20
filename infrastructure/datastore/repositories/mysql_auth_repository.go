package repositories

import (
	"context"
	"database/sql"
	"errors"
	"taskflow/domain/entities"
	"taskflow/domain/repositories"

	"github.com/google/uuid"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) repositories.AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r authRepository) AddUser(ctx context.Context, user *entities.User) error {
	const query = `
		INSERT INTO users (uuid, name, email, password) VALUES (?, ?, ?, ?)
	`

	_, err := r.db.ExecContext(ctx, query, user.UUID, user.Name, user.Email, user.Password)
	if err != nil {
		return errors.Join(entities.ErrExecuteQuery, err)
	}

	return nil
}

func (r authRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	const query = `
		SELECT id, uuid, name, email, password FROM users WHERE email = ?
	`

	var user entities.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Join(entities.ErrExecuteQuery, err)
	}

	return &user, nil
}

func (r authRepository) GetUserByID(ctx context.Context, id int) (*entities.User, error) {
	const query = `
		SELECT id, uuid, name, email, password FROM users WHERE id = ?
	`

	var user entities.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Join(entities.ErrQueryRow, err)
	}

	return &user, nil
}

func (r authRepository) GetUserByUUID(ctx context.Context, uuid uuid.UUID) (*entities.User, error) {
	const query = `
		SELECT id, uuid, name, email, password FROM users WHERE id = ?
	`

	var user entities.User
	err := r.db.QueryRowContext(ctx, query, uuid).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Join(entities.ErrQueryRow, err)
	}

	return &user, nil
}

func (r authRepository) DeleteUser(ctx context.Context, id int) error {
	const query = `
		DELETE FROM users WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Join(entities.ErrExecuteQuery, err)
	}

	return nil
}
