package auth

import (
	"time"

	"gorm.io/gorm"
)

// Auth представляет данные для аутентификации
type Auth struct {
	ID        uint           `json:"id" gorm:"primarykey" example:"1"`
	CreatedAt time.Time      `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	Email     string         `json:"email" example:"user@example.com"`
	Phone     string         `json:"phone" example:"+79123456789"`
	Code      string         `json:"code" example:"1234"`
}
