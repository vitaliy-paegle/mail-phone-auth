package auth

type AuthEmailCodeRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type AuthEmailConfirmRequest struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

type AuthEmailConfirmResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
