package auth

import "mail-phone-auth/pkg/jwt"

// AuthEmailRequest запрос на аутентификацию по email
type AuthEmailRequest struct {
	Email string `json:"email" validate:"required,email" example:"user@example.com"`
}

// AuthEmailConfirmRequest запрос на подтверждение email
type AuthEmailConfirmRequest struct {
	Email string `json:"email" validate:"required,email" example:"user@example.com"`
	Code  string `json:"code" validate:"required,len=4" example:"1234"`
}

// AuthResponse ответ с токенами
type AuthResponse struct {
	jwt.TokensSet
	User UserInfo `json:"user"`
}

// UserInfo информация о пользователе в ответе
type UserInfo struct {
	ID    uint   `json:"id" example:"1"`
	Name  string `json:"name" example:"Иван Иванов"`
	Email string `json:"email" example:"user@example.com"`
	Phone string `json:"phone" example:"+79123456789"`
}
