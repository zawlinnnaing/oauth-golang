package user

type SignInBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Token        TokenResponse `json:"token"`
	RefreshToken TokenResponse `json:"refresh_token"`
}
