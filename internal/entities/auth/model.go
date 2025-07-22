package auth

import (
	"mail-phone-auth/internal/api/data"
)

type Auth struct {
	data.Base
	Email string `json:"email,omitempty" gorm:"column:email;comment:Электронная почта пользователя"`
	Phone string `json:"phone,omitempty" gorm:"column:phone;comment:Телефон пользователя"`
	Code  string `json:"code" gorm:"column:code;comment:Код подтверждения авторизации;not null"`
}
