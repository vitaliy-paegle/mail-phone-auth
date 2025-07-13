package user

import (
	"log"
	"mail-phone-auth/internal/api/request"
	"mail-phone-auth/internal/api/response"
	"net/http"
	"strconv"
)

type IAPI interface {
	TestAPI()
}

type Handler struct {
	Router     *http.ServeMux
	Repository *Repository
	API        IAPI
}

func NewHandler(router *http.ServeMux, repository *Repository, api IAPI) *Handler {
	handler := Handler{
		Router:     router,
		Repository: repository,
		API:        api,
	}

	handler.API.TestAPI()

	handler.Router.HandleFunc("GET /api/user/{id}", handler.Read)
	handler.Router.HandleFunc("GET /api/user/all", handler.ReadAll)
	handler.Router.HandleFunc("POST /api/user", handler.Create)
	handler.Router.HandleFunc("PATCH /api/user/{id}", handler.Update)
	handler.Router.HandleFunc("DELETE /api/user/{id}", handler.Delete)

	return &handler
}

// Create создает нового пользователя
// @Summary Создать пользователя
// @Description Создает нового пользователя в системе
// @Tags users
// @Accept json
// @Produce json
// @Param request body UserCreateRequest true "Данные пользователя"
// @Success 201 {object} User "Пользователь успешно создан"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 409 {object} map[string]string "Пользователь уже существует"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /user [post]
func (handler *Handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := request.DecodeBody[UserCreateRequest](r.Body)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := NewUser(body)

	err = handler.Repository.Create(user)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		response.JSON(w, user, http.StatusCreated)
	}

}

// Read получает пользователя по ID
// @Summary Получить пользователя по ID
// @Description Возвращает данные пользователя по его ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} User "Данные пользователя"
// @Failure 400 {object} map[string]string "Неверный ID"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /user/{id} [get]
func (handler *Handler) Read(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := handler.Repository.Read(id)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		response.JSON(w, user, http.StatusOK)
	}

}

// ReadAll получает список всех пользователей
// @Summary Получить список пользователей
// @Description Возвращает список всех пользователей с пагинацией
// @Tags users
// @Accept json
// @Produce json
// @Param limit query int false "Лимит записей" default(10)
// @Param offset query int false "Смещение" default(0)
// @Success 200 {object} UserAllResponse "Список пользователей"
// @Failure 400 {object} map[string]string "Неверные параметры"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /user/all [get]
func (handler *Handler) ReadAll(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

	if err != nil {
		limit = -1
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))

	if err != nil {
		offset = 0
	}

	users := handler.Repository.ReadAll(limit, offset)
	response.JSON(w, &users, http.StatusOK)

}

// Update обновляет данные пользователя
// @Summary Обновить данные пользователя
// @Description Обновляет данные пользователя по его ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param request body UserUpdateRequest true "Обновленные данные"
// @Success 200 {object} User "Обновленные данные пользователя"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /user/{id} [patch]
func (handler *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := request.DecodeBody[UserUpdateRequest](r.Body)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// user := User{
	// 	Model: gorm.Model{ID: uint(id)},
	// 	Name: body.Name,
	// 	Phone: body.Phone,
	// }

	user, err := handler.Repository.Read(id)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Name = body.Name
	user.Phone = body.Phone

	err = handler.Repository.Update(user)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		response.JSON(w, user, http.StatusOK)
	}
}

// Delete удаляет пользователя
// @Summary Удалить пользователя
// @Description Удаляет пользователя по его ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} User "Данные удаленного пользователя"
// @Failure 400 {object} map[string]string "Неверный ID"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера, попробуйте позже"
// @Security BearerAuth
// @Router /user/{id} [delete]
func (handler *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := handler.Repository.Read(id)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.Repository.Delete(id)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		response.JSON(w, user, http.StatusOK)
	}
}
