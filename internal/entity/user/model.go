package user

import (
	"mail-phone-auth/internal/entity"
	"mail-phone-auth/internal/entity/role"
)

type UserData struct {
	Name   *string    `json:"name" gorm:"column:name;comment:Имя пользователя;default:null;"`
	Phone  *string    `json:"phone" validate:"omitempty,e164" gorm:"column:phone;comment:Номер телефона;unique;default:null;"`
	Email  string     `json:"email" validate:"required,email" gorm:"column:email;comment:Электронная почта;unique;not null;"`
	RoleID *uint      `json:"roleId" gorm:"column:role_id;comment:Идентификатор роли;default:null;"`
	Role   *role.Role `json:"role,omitempty" gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type User struct {
	entity.GeneralData
	UserData
}
