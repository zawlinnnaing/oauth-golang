package validators

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

var PasswordValidator validator.Func = func(fl validator.FieldLevel) bool {
	password := fl.Field().String()
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
