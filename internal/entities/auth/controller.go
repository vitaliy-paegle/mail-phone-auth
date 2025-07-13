package auth

import (
	"log"
	"mail-phone-auth/internal/api/request"
	"mail-phone-auth/internal/api/response"
	"net/http"
)

type Controller struct {
	router     *http.ServeMux
	repository *Repository
}

func NewController(router *http.ServeMux, repository *Repository) *Controller {
	controller := Controller{
		router:     router,
		repository: repository,
	}

	controller.router.HandleFunc("POST /api/auth/email", controller.CreateEmailAuth)
	controller.router.HandleFunc("POST /api/auth/email/confirm", controller.ConfirmEmailAuth)

	return &controller
}

// CreateEmailAuth создает запрос на аутентификацию по email
// @Summary Создать запрос на аутентификацию по email
// @Description Отправляет код подтверждения на указанный email адрес
// @Tags auth
// @Accept json
// @Produce json
// @Param request body AuthEmailRequest true "Email для аутентификации"
// @Success 201 {object} Auth "Успешно создан запрос на аутентификацию"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /auth/email [post]
func (c *Controller) CreateEmailAuth(w http.ResponseWriter, r *http.Request) {
	body, err := request.DecodeBody[AuthEmailRequest](r.Body)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	auth := Auth{
		Email: body.Email,
		Code:  "1234",
	}

	err = c.repository.CreateEmailAuth(&auth)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.JSON(w, &auth, http.StatusCreated)

}

// ConfirmEmailAuth подтверждает email аутентификацию
// @Summary Подтвердить аутентификацию по email
// @Description Проверяет код подтверждения из Email и выдает JWT-токены
// @Tags auth
// @Accept json
// @Produce json
// @Param request body AuthEmailConfirmRequest true "Email и код подтверждения"
// @Success 200 {object} AuthResponse "Успешная аутентификация"
// @Failure 400 {object} map[string]string "Неверный код или email"
// @Failure 401 {object} map[string]string "Неавторизован"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /auth/email/confirm [post]
func (controller *Controller) ConfirmEmailAuth(w http.ResponseWriter, r *http.Request) {
	// Реализация подтверждения
	log.Println("Email are confrimed, Welcome!")
}
