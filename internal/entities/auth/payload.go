package auth

import "mail-phone-auth/pkg/jwt"

type AuthEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type AuthEmailConfirmRequest struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

type AuthResponse struct {
	jwt.TokensSet
}
