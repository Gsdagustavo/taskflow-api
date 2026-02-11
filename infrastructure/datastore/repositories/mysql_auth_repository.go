package repositories

import (
	"context"
	"database/sql"
	"errors"
	"taskflow/domain/entities"
	"taskflow/domain/util"
	"taskflow/infrastructure/datastore"

	"github.com/google/uuid"
)

type authRepository struct {
	conn func() *sql.DB
}

func NewAuthRepository(settings datastore.RepositorySettings) datastore.AuthRepository {
	return authRepository{
		conn: settings.Connection,
	}
}

func (r authRepository) AddUser(ctx context.Context, user *entities.User) error {
	const query = `
		INSERT INTO users (uuid, email, password) VALUES (?, ?, ?)
	`

	_, err := r.conn().ExecContext(ctx, query, user.UUID, user.Email, user.Password)
	if err != nil {
		return errors.Join(entities.ErrExecuteQuery, err)
	}

	return nil
}

func (r authRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	const query = `
		SELECT id, uuid, email, password FROM users WHERE email = ?
	`

	var user entities.User
	err := r.conn().QueryRowContext(ctx, query, email).Scan(&user.ID, &user.UUID, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound
		}

		return nil, errors.Join(entities.ErrExecuteQuery, err)
	}

	return &user, nil
}

func (r authRepository) GetUserByID(ctx context.Context, id int) (*entities.User, error) {
	const query = `
		SELECT id, uuid, email, password FROM users WHERE id = ?
	`

	var user entities.User
	err := r.conn().QueryRowContext(ctx, query, id).Scan(&user.ID, &user.UUID, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound
		}

		return nil, errors.Join(entities.ErrQueryRow, err)
	}

	return &user, nil
}

func (r authRepository) GetUserByUUID(ctx context.Context, uuid uuid.UUID) (*entities.User, error) {
	const query = `
		SELECT id, uuid, email, password FROM users WHERE id = ?
	`

	var user entities.User
	err := r.conn().QueryRowContext(ctx, query, uuid).Scan(&user.ID, &user.UUID, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrNotFound
		}

		return nil, errors.Join(entities.ErrQueryRow, err)
	}

	return &user, nil
}

func (r authRepository) DeleteUser(ctx context.Context, id int) error {
	const query = `
		DELETE FROM users WHERE id = ?
	`

	_, err := r.conn().ExecContext(ctx, query, id)
	if err != nil {
		return errors.Join(entities.ErrExecuteQuery, err)
	}

	return nil
}

func (r authRepository) CheckUserCredentials(
	ctx context.Context,
	credentials entities.UserCredentials,
) (bool, error) {
	query := `
	SELECT password 
	FROM users
	WHERE email = ?
	  AND status_code = 0
	`

	var password, salt string
	err := r.conn().QueryRowContext(ctx, query, credentials.Email).Scan(&password, &salt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, entities.ErrNotFound
		}
		return false, errors.Join(entities.ErrQueryRow, err)
	}

	// Pre-register users have an empty password; make sure they can't log in
	if password == "" || salt == "" {
		return false, nil
	}

	valid, err := util.CheckValidPassword(credentials.Password, password)
	if err != nil {
		return false, errors.Join(errors.New("failed to check if password is valid"), err)
	}

	return valid, nil
}
