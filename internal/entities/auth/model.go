package auth

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Email string `json:"email"`
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
