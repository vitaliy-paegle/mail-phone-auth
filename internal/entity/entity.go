package entity

import (
	"errors"
	"log"
	"mail-phone-auth/internal/api/request"
	"mail-phone-auth/internal/api/response"
	"mail-phone-auth/pkg/postgres"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type GeneralData struct {
	ID        uint       `json:"id" gorm:"column:id;comment:Идентификатор;primaryKey;autoIncrement:true;"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at;comment:Дата и время создания;not null;"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updated_at;comment:Дата и время обновления;not null;"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"column:deleted_at;comment:Дата и время удаления;default:null;"`
}

type EntityList[M any] struct {
	List  []M `json:"list"`
	Count int `json:"count"`
}

type Error struct {
	Message string `json:"message"`
}

type Entity[M any, B any] struct {
	postgres *postgres.Postgres
}

func New[M any, B any](postgres *postgres.Postgres) *Entity[M, B] {
	return &Entity[M, B]{postgres: postgres}
}

func (e *Entity[M, B]) Create(w http.ResponseWriter, r *http.Request) {

	body, err := request.DecodeBody[B](r.Body)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model, err := CreateModelInstance[M]()

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = e.UpdateData(model, &body)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := e.postgres.DB.Create(model)

	if result.Error != nil {
		log.Println(err)
		response.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	response.JSON(w, model, http.StatusCreated)

}

func (e *Entity[M, B]) Read(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model, err := CreateModelInstance[M]()

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := e.postgres.DB.First(model, id)

	if result.Error != nil {
		log.Println(result.Error)
		response.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	response.JSON(w, model, http.StatusOK)

}

func (e *Entity[M, B]) ReadAll(w http.ResponseWriter, r *http.Request) {

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

	if err != nil {
		limit = -1
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))

	if err != nil {
		offset = 0
	}

	entityList := EntityList[M]{}

	result := e.postgres.DB.Table("users").
		Where("deleted_at is NULL").
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Scan(&entityList.List)

	if result.Error != nil {
		log.Println(result.Error)
		response.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	entityList.Count = len(entityList.List)

	response.JSON(w, &entityList, http.StatusOK)

}

func (e *Entity[M, B]) Update(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := request.DecodeBody[B](r.Body)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model, err := CreateModelInstance[M]()

	result := e.postgres.DB.First(model, id)

	if result.Error != nil {
		log.Println(err)
		response.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	e.UpdateData(model, &body)

	result = e.postgres.DB.Updates(model)

	if result.Error != nil {
		log.Println(err)
		response.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	response.JSON(w, model, http.StatusOK)

}

func (e *Entity[M, B]) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		log.Println(err)
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model, err := CreateModelInstance[M]()

	result := e.postgres.DB.First(model, id)

	if result.Error != nil {
		log.Println(err)
		response.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	modelDeletedAt := reflect.ValueOf(model).Elem().FieldByName("DeletedAt")

	if modelDeletedAt.IsValid() && modelDeletedAt.CanSet() && modelDeletedAt.IsZero() {
		time := time.Now()
		modelDeletedAt.Set(reflect.ValueOf(&time))
	} else {
		err := errors.New("Error update field: 'DeletedAt'")
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result = e.postgres.DB.Updates(model)

	if result.Error != nil {
		log.Println(err)
		response.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	response.JSON(w, model, http.StatusOK)
}

func CreateModelInstance[M any]() (*M, error) {
	var zero M
	typ := reflect.TypeOf(zero)

	if typ.Kind() != reflect.Struct {
		return nil, errors.New("Error create model")
	}
	value := reflect.New(typ)

	return value.Interface().(*M), nil
}

func (e *Entity[M, B]) UpdateData(model *M, body *B) error {

	values := reflect.ValueOf(*body)
	types := reflect.TypeOf(*body)

	for index := range values.NumField() {

		fieldName := types.Field(index).Name

		entityField := reflect.ValueOf(model).Elem().FieldByName(fieldName)

		if entityField.IsValid() && entityField.CanSet() {
			entityField.Set(values.Field(index))
		} else {
			err := errors.New("Error update field: " + fieldName)
			return err
		}
	}

	entityCreatedAt := reflect.ValueOf(model).Elem().FieldByName("CreatedAt")

	if entityCreatedAt.IsValid() && entityCreatedAt.CanSet() && entityCreatedAt.IsZero() {
		entityCreatedAt.Set(reflect.ValueOf(time.Now()))
	} else {
		err := errors.New("Error update field: 'CreatedAt'")
		return err
	}

	entityUpdatedAt := reflect.ValueOf(model).Elem().FieldByName("UpdatedAt")

	if entityUpdatedAt.IsValid() && entityUpdatedAt.CanSet() {
		entityUpdatedAt.Set(reflect.ValueOf(time.Now()))
	} else {
		err := errors.New("Error update field: 'UpdatedAt'")
		return err
	}

	return nil

}
