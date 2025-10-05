package usecases

import (
	"context"
	"fmt"
	"log"
	"taskflow/domain/entities"
	"taskflow/domain/repositories"
	"taskflow/domain/rules"
	"taskflow/domain/status_codes"
	"taskflow/infrastructure/util"

	"github.com/google/uuid"
)

type AuthUseCases struct {
	repository repositories.AuthRepository
	crypt      util.Crypt
}

func NewAuthUseCases(repository repositories.AuthRepository, crypt util.Crypt) *AuthUseCases {
	return &AuthUseCases{
		repository: repository,
		crypt:      crypt,
	}
}

func (a AuthUseCases) AttemptLogin(ctx context.Context, credentials entities.UserCredentials) (string, status_codes.LoginStatusCode, error) {
	// Check if the user exists
	user, err := a.repository.GetUserByEmail(ctx, credentials.Email)
	if err != nil {
		return "", status_codes.LoginFailure, fmt.Errorf("[AttemptLogin] error checking user: %s", err)
	}

	if user == nil {
		return "", status_codes.LoginUserNotFound, nil
	}

	if !a.crypt.CheckPasswordHash(credentials.Password, user.Password) {
		return "", status_codes.LoginInvalidCredentials, nil
	}

	// Generate token
	token, err := a.crypt.GenerateAuthToken(credentials.Email)
	if err != nil {
		return "", status_codes.LoginFailure, fmt.Errorf("[AttemptLogin] error generating token: %s", err)
	}

	return token, status_codes.LoginSuccess, nil
}

func (a AuthUseCases) RegisterUser(ctx context.Context, credentials entities.UserCredentials) (status_codes.RegisterStatusCode, error) {
	// Check if the user exists
	user, err := a.repository.GetUserByEmail(ctx, credentials.Email)
	if err != nil {
		return status_codes.RegisterFailure, fmt.Errorf("[RegisterUser] error checking user: %s", err)
	}

	if user != nil {
		return status_codes.RegisterUserAlreadyExist, nil
	}

	// Validate credentials
	credentials.Email = util.TrimSpace(credentials.Email)
	credentials.Password = util.TrimSpace(credentials.Password)
	credentials.Name = util.TrimSpace(credentials.Name)

	if !rules.IsValidName(credentials.Name) {
		log.Printf("[RegisterUser] invalid name: %s", credentials.Name)
		return status_codes.RegisterInvalidName, nil
	}

	if !rules.IsValidEmail(credentials.Email) {
		log.Printf("[RegisterUser] invalid email: %s", credentials.Email)
		return status_codes.RegisterInvalidEmail, nil
	}

	if !rules.IsValidPassword(credentials.Password) {
		log.Printf("[RegisterUser] invalid password: %s", credentials.Password)
		return status_codes.RegisterInvalidPassword, nil
	}

	// Hash user password before saving
	credentials.Password, err = a.crypt.HashPassword(credentials.Password)
	if err != nil {
		return status_codes.RegisterFailure, fmt.Errorf("[RegisterUser] error hashing password: %s", err)
	}

	userUUID, err := uuid.NewRandom()
	if err != nil {
		return status_codes.RegisterFailure, fmt.Errorf("[RegisterUser] error generating user uuid: %s", err)
	}

	user = &entities.User{
		UUID:     userUUID,
		Name:     credentials.Name,
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	// Save user
	err = a.repository.AddUser(ctx, user)
	if err != nil {
		return status_codes.RegisterFailure, fmt.Errorf("[RegisterUser] error saving user: %s", err)
	}

	return status_codes.RegisterSuccess, nil
}
