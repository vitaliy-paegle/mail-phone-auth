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

func (controller *Controller) ConfirmEmailAuth(w http.ResponseWriter, r *http.Request) {

}
