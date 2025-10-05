package rules

import (
	"net/mail"
	"unicode"
)

// Password rules
const (
	PasswordMinLetters = 8
	PasswordMinNumbers = 1
	PasswordMinSpecial = 1
	PasswordMaxLetters = 32
)

// Name rules
const (
	NameMinLetters = 3
	NameMaxLetters = 32
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidPassword(password string) bool {
	var letters int
	var number bool
	var special bool

	for _, char := range password {
		letters++

		switch {
		case unicode.IsNumber(char):
			number = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			special = true
		}
	}

	return letters >= PasswordMinLetters && letters <= PasswordMaxLetters && number && special
}

func IsValidName(name string) bool {
	if len(name) < NameMinLetters || len(name) > NameMaxLetters {
		return false
	}

	return true
}
