package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

var ErrExpiredToken = errors.New("TOKEN_HAS_EXPIRED")

type Maker interface {
	CreateToken(email string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}

func NewPayload(email string, duration time.Duration) (*Payload, error) {
	tokenUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("error generating token uuid: %s", err)
	}

	payload := &Payload{
		ID:        tokenUUID,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return payload, nil
}

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey string
}

func NewPasetoMaker(symmetricKey string) Maker {
	return &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: symmetricKey,
	}
}

func (m *PasetoMaker) CreateToken(email string, duration time.Duration) (string, error) {
	payload, err := NewPayload(email, duration)
	if err != nil {
		return "", err
	}

	encrypted, err := m.paseto.Encrypt([]byte(m.symmetricKey), payload, nil)
	if err != nil {
		return "", fmt.Errorf("error encrypting token: %s", err)
	}

	return encrypted, nil
}

func (m *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	var payload Payload
	err := m.paseto.Decrypt(token, []byte(m.symmetricKey), &payload, nil)
	if err != nil {
		return nil, fmt.Errorf("error decrypting token: %s", err)
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
