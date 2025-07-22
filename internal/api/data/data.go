package data

import "time"

type Base struct {
	ID        uint       `json:"id" gorm:"column:id;comment:Идентификатор;primaryKey;autoIncrement:true;"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at;comment:Дата и время создания;not null"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updated_at;comment:Дата и время обновления;not null"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"column:deleted_at;comment:Дата и время удаления"`
}

type Error struct {
	Message string `json:"message"`
}
