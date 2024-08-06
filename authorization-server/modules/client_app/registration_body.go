package client_app

import "github.com/zawlinnnaing/oauth-golang/authorization-server/modules/app_error"

type RegistrationBody struct {
	Name        string  `json:"name" validate:"required"`
	RedirectURI *string `json:"redirect_uri"`
}

func (body RegistrationBody) Validate() error {
	err := app_error.NewValidationError()
	if body.Name == "" {
		err.Set("name", "is-required")
	}
	if len(err.Errors) > 0 {
		return err
	}
	return nil
}
