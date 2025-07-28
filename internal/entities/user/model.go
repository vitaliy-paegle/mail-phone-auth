package user

import (
	"mail-phone-auth/internal/api/data"
)

type User struct {
	data.Base
	Name  string `json:"name" gorm:"column:name;comment:Имя пользователя;default:null"`
	Phone string `json:"phone" gorm:"column:phone;comment:Номер телефона;unique;default:null"`
	Email string `json:"email" gorm:"column:email;comment:Электронная почта;unique;default:null"`
}
