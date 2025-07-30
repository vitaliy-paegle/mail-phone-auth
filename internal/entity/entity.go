package entity

import (
	"log"
	"reflect"
	"time"
)

type Data struct {
	ID        uint       `json:"id" gorm:"column:id;comment:Идентификатор;primaryKey;autoIncrement:true;"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at;comment:Дата и время создания;not null"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updated_at;comment:Дата и время обновления;not null"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"column:deleted_at;comment:Дата и время удаления;"`
}

type Error struct {
	Message string `json:"message"`
}

func Update[E any, B any] (entity *E, body *B) {

	values := reflect.ValueOf(*body)
	types := reflect.TypeOf(*body)	

	for index := range values.NumField() {	

		fieldName := types.Field(index).Name

		entityField := reflect.ValueOf(entity).Elem().FieldByName(fieldName)

		if entityField.IsValid() && entityField.CanSet() {
			entityField.Set(values.Field(index))
		} else {
			log.Println("Error update field: ", fieldName)
		}
	}

	entityCreatedAt := reflect.ValueOf(entity).Elem().FieldByName("CreatedAt")

	if entityCreatedAt.IsValid() && entityCreatedAt.CanSet() && entityCreatedAt.IsZero() {
		entityCreatedAt.Set(reflect.ValueOf(time.Now()))
	} else {
		log.Println("Error update field: 'CreatedAt'")
	}

	entityUpdatedAt := reflect.ValueOf(entity).Elem().FieldByName("UpdatedAt")

	if entityUpdatedAt.IsValid() && entityUpdatedAt.CanSet() {
		entityUpdatedAt.Set(reflect.ValueOf(time.Now()))
	} else {
		log.Println("Error update field: 'UpdatedAt'")
	}

}

