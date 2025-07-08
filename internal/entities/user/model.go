package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Phone string `json:"phone" gorm:"uniqueIndex"`
	Email string `json:"email" gorm:"uniqueIndex"`
}

func NewUser(data UserCreateRequest) *User {
	user := User{
		Name: data.Name,
		Phone: data.Phone,
		Email: data.Email,
	}
	return &user
}