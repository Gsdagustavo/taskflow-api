package usecases

import (
	"context"
	"errors"
	"strings"
	"taskflow/domain/entities"
	"taskflow/domain/rules"
	"taskflow/domain/status_codes"
	"taskflow/domain/util"
	"taskflow/infrastructure/datastore"

	"github.com/google/uuid"
)

type AuthUseCases struct {
	repository        datastore.AuthRepository
	pasetoSecurityKey string
}

func NewAuthUseCases(repository datastore.AuthRepository, pasetoSecurityKey string) AuthUseCases {
	return AuthUseCases{
		repository:        repository,
		pasetoSecurityKey: pasetoSecurityKey,
	}
}

func (a AuthUseCases) AttemptLogin(ctx context.Context, credentials entities.UserCredentials) (string, status_codes.LoginStatusCode, error) {
	// Check if the user exists
	user, err := a.repository.GetUserByEmail(ctx, credentials.Email)
	if err != nil {
		return "", status_codes.LoginFailure, errors.Join(errors.New("failed to get user by email"), err)
	}

	if user == nil {
		return "", status_codes.LoginUserNotFound, nil
	}

	validPassword, err := util.CheckValidPassword(credentials.Password, user.Password)
	if err != nil {
		return "", status_codes.LoginFailure, errors.Join(errors.New("failed to check if password is valid"), err)
	}

	if !validPassword {
		return "", status_codes.LoginInvalidCredentials, nil
	}

	// Generate token
	token, err := util.GetNewAuthToken(user.ID, user.UUID, a.pasetoSecurityKey)
	if err != nil {
		return "", status_codes.LoginFailure, errors.Join(errors.New("failed to generate auth token"), err)
	}

	return token, status_codes.LoginSuccess, nil
}

func (a AuthUseCases) RegisterUser(ctx context.Context, credentials entities.UserCredentials) (status_codes.RegisterStatusCode, error) {
	// Check if the user exists
	user, err := a.repository.GetUserByEmail(ctx, credentials.Email)
	if err != nil {
		return status_codes.RegisterFailure, errors.Join(errors.New("failed to get user by email"), err)
	}

	if user != nil {
		return status_codes.RegisterUserAlreadyExist, nil
	}

	// Validate credentials
	credentials.Email = strings.TrimSpace(credentials.Email)
	credentials.Password = strings.TrimSpace(credentials.Password)

	if !rules.ValidateEmail(credentials.Email) {
		return status_codes.RegisterInvalidEmail, nil
	}

	if !rules.ValidatePassword(credentials.Password) {
		return status_codes.RegisterInvalidPassword, nil
	}

	// Hash user password before saving
	credentials.Password, err = util.Hash(credentials.Password)
	if err != nil {
		return status_codes.RegisterFailure, errors.Join(errors.New("failed to hash password"), err)
	}

	userUUID, err := uuid.NewRandom()
	if err != nil {
		return status_codes.RegisterFailure, errors.Join(errors.New("failed to generate user UUID"), err)
	}

	user = &entities.User{
		UUID:     userUUID.String(),
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	// Save user
	err = a.repository.AddUser(ctx, user)
	if err != nil {
		return status_codes.RegisterFailure, errors.Join(errors.New("error saving user"), err)
	}

	return status_codes.RegisterSuccess, nil
}

func (a AuthUseCases) CheckCredentials(
	ctx context.Context,
	credentials entities.UserCredentials,
) (bool, error) {
	return a.repository.CheckUserCredentials(ctx, credentials)
}
