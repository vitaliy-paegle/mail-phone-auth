package auth

import "mail-phone-auth/internal/entity"

type Auth struct {
	entity.Data
	Email string `json:"email" gorm:"column:email;comment:Электронная почта пользователя;default:null"`
	Phone string `json:"phone" gorm:"column:phone;comment:Телефон пользователя;default:null"`
	Code  string `json:"code" gorm:"column:code;comment:Код подтверждения авторизации;not null"`
}
