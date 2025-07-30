package auth

type AuthEmailCodeRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type AuthEmailConfirmRequest struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

type AuthJwtTokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthRefreshRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
