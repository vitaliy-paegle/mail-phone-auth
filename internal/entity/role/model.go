package role

import "mail-phone-auth/internal/entity"

type RoleData struct {
	Name string `json:"name" validate:"required" gorm:"column:name;comment:Имя роли;unique;not null;"`
}

type Role struct {
	entity.GeneralData
	RoleData
}
