package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name111 string `json:"name111"`
	Phone string `json:"phone" gorm:"uniqueIndex"`
	Email string `json:"email" gorm:"uniqueIndex"`
}

func New() *User {
	model := User{}
	return &model
}