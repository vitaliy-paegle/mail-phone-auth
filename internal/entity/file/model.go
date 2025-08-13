package file

import (
	"mail-phone-auth/internal/entity"
	"mail-phone-auth/internal/entity/user"
)

type FileData struct {
	Name      string     `json:"name" gorm:"column:name;comment:Имя файла;not null;"`
	Extension *string    `json:"extension" gorm:"column:extension;comment:Расширение файла;default:null;"`
	Hash      string     `json:"hash" gorm:"column:hash;comment:Хэш сумма файла;unique;not null;"`
	Link      string     `json:"link" gorm:"column:link;comment:Ссылка на файл;unique;not null;"`
	UserID    uint       `json:"userId" gorm:"column:user_id;comment:Идентификатор пользователя;not null;"`
	User      *user.User `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type File struct {
	entity.GeneralData
	FileData
}
