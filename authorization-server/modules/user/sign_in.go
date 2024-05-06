package user

import "github.com/zawlinnnaing/oauth-golang/authorization-server/modules/app_error"

type SignInBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (body SignInBody) Validate() error {
	err := app_error.NewValidationError()
	if body.Email == "" {
		err.Set("email", "is-required")
	}
	if body.Password == "" {
		err.Set("password", "is-required")
	}
	if len(err.Errors) > 0 {
		return err
	}
	return nil
}

type SignInResponse struct {
	Token        TokenResponse `json:"token"`
	RefreshToken TokenResponse `json:"refresh_token"`
}
