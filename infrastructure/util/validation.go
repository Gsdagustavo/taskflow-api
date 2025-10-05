package util

import (
	"strings"
)

func TrimSpace(input string) string {
	return strings.TrimSpace(input)
}

func IsEmpty(input string) bool {
	return TrimSpace(input) == ""
}

func IsNotEmpty(input string) bool {
	return !IsEmpty(input)
}
