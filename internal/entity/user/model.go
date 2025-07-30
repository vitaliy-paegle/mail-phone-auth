package user

import "mail-phone-auth/internal/entity"

// type Data struct {
// 	Name  *string `json:"name"`
// 	Phone *string `json:"phone" validate:"omitempty,e164"`
// 	Email string `json:"email" validate:"required,email"`
// }

type User struct {
	entity.Data
	Name  *string `json:"name" gorm:"column:name;comment:Имя пользователя;default:null;"`
	Phone *string `json:"phone" gorm:"column:phone;comment:Номер телефона;unique;default:null;"`
	Email string `json:"email" gorm:"column:email;comment:Электронная почта;unique;not null;"`
}
