package user

import (
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/app_error"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/common"
)

type SignUpBody struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"Password"`
}

func (body *SignUpBody) Validate() error {
	validationError := app_error.NewValidationError()
	if body.Email == "" {
		validationError.Set("email", "is-required")
	} else if !common.IsEmail(body.Email) {
		validationError.Set("email", "is-invalid")
	}
	if body.FirstName == "" {
		validationError.Set("first_name", "is-required")
	}
	if body.LastName == "" {
		validationError.Set("last_name", "is-required")
	}
	if body.Password == "" {
		validationError.Set("Password", "is-required")
	} else if !common.IsValidPassword(body.Password) {
		validationError.Set("Password", "is-invalid")
	}
	if len(validationError.Errors) > 0 {
		return validationError
	}
	return nil
}
