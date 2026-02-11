package entities

import "time"

type User struct {
	ID         int       `json:"id"`
	UUID       string    `json:"uuid"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
