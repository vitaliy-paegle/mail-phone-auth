package user

import (
	"time"

	"gorm.io/gorm"
)

// User - представление пользователя в системе
type User struct {
	ID        uint           `json:"id" gorm:"primarykey" example:"1"`
	CreatedAt time.Time      `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name" example:"Иван Иванов"`
	Phone     string         `json:"phone" gorm:"uniqueIndex" example:"+79123456789"`
	Email     string         `json:"email" gorm:"uniqueIndex" example:"user@example.com"`
}

func NewUser(data UserCreateRequest) *User {
	user := User{
		Name:  data.Name,
		Phone: data.Phone,
		Email: data.Email,
	}
	return &user
}
