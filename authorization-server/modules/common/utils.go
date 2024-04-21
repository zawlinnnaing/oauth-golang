package common

import (
	"net/mail"
	"unicode"
)

func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidPassword(password string) bool {
	var (
		hasMinLen    = false
		hasUppercase = false
		hasLowercase = false
		hasNumber    = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	return hasMinLen && hasUppercase && hasLowercase && hasNumber
}
